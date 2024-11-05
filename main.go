package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strings" // Add this line
    "github.com/fiatjaf/khatru"
    "github.com/nbd-wtf/go-nostr"
)

func main() {
    // Create a new relay instance
    relay := khatru.NewRelay()

    // Set up properties for your relay (optional)
    relay.Info.Name = "My Custom Relay"
    relay.Info.Description = "A custom Nostr relay with content filtering"

    // Define a list of banned words for filtering
    bannedWords := []string{"badword1", "badword2"} // Replace with your chosen words

    // Add custom content filtering to RejectEvent
    relay.RejectEvent = append(relay.RejectEvent,
        func(ctx context.Context, event *nostr.Event) (reject bool, msg string) {
            for _, word := range bannedWords {
                if strings.Contains(event.Content, word) {
                    log.Printf("Event rejected: %s, Reason: Contains prohibited content", event.ID)
                    return true, "Content rejected due to prohibited words"
                }
            }
            return false, ""
        },
    )

    // Add a custom HTTP handler for event submission
    http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
            return
        }

        var event nostr.Event
        err := json.NewDecoder(r.Body).Decode(&event)
        if err != nil {
            http.Error(w, "Invalid event format", http.StatusBadRequest)
            return
        }

        ctx := context.Background()
        rejected, msg := false, ""

        // Run the event through the rejection filters
        for _, rejectFunc := range relay.RejectEvent {
            if r, m := rejectFunc(ctx, &event); r {
                rejected, msg = r, m
                break
            }
        }

        if rejected {
            http.Error(w, msg, http.StatusForbidden)
            log.Printf("Rejected event: %s, Reason: %s", event.ID, msg)
            return
        }

        // If not rejected, acknowledge receipt
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Event received"))
    })

    // Start the relay server on a different port for handling relay-specific connections
    started := make(chan bool)
    go func() {
        if err := relay.Start("localhost", 3334, started); err != nil {
            log.Fatalf("Failed to start relay: %v", err)
        }
    }()

    // Start the HTTP server for custom handler
    go func() {
        log.Println("HTTP server running on :3335")
        http.ListenAndServe(":3335", nil)
    }()

    // Wait for the relay to start
    <-started
    fmt.Println("Relay is running on http://localhost:3334")

    // Keep the server running
    select {}
}


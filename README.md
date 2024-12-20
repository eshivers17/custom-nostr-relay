Custom Nostr Relay

A custom Nostr relay implementation using the khatru framework, designed to filter events and reject submissions with prohibited content. The project now includes SGX sample code (in the sgx directory) to demonstrate the groundwork for future Trusted Execution Environment (TEE) integration.
Features

    Event content filtering with customizable banned words
    HTTP endpoint /submit for event submissions
    Logging for rejected events
    Integration of SGX sample code under sgx/ directory, providing a base for future TEE-based security enhancements

Directory Structure

custom-nostr-relay/
├─ main.go          // Relay server entry point
├─ sgx/             // SGX-related code and samples
│  ├─ SampleCode/
│  │  ├─ LocalAttestation/
│  │  ├─ EnclaveResponder/
│  │  ├─ EnclaveInitiator/
│  │  └─ App/
│  └─ ...other SGX files
└─ README.md

Installation

    Clone the repository:

git clone https://github.com/eshivers17/custom-nostr-relay.git
cd custom-nostr-relay

Install Go dependencies: Ensure Go is installed and properly set up:

    go mod tidy

Usage

    Run the relay:

go run main.go

The relay should start listening on http://localhost:3334 for relay connections and http://localhost:3335 for /submit events.

Submit an event using curl:

    curl -X POST http://localhost:3335/submit \
         -H "Content-Type: application/json" \
         -d '{"content": "Test event content"}'

    If the content is prohibited, the event will be rejected with an appropriate message logged by the relay.

SGX Integration

    The sgx directory contains the SGX sample code (LocalAttestation, EnclaveResponder, EnclaveInitiator, and App) from the Intel SGX SDK examples.
    This code provides a foundation for future TEE integration, where the relay’s filtering logic can be run inside a secure enclave.
    Before using SGX samples, ensure you have the Intel SGX SDK installed and properly set up on your system. Refer to the Intel SGX documentation for environment setup, attestation procedures, and building the SGX code.

Next Steps

    Advanced Filtering Logic: Add regex-based filtering or integrate AI/ML models for content moderation.
    User Authentication: Introduce user keys and authorization to control who can submit events.
    Full TEE Integration: Run filtering logic inside an SGX enclave for cryptographic proof of integrity and secure content handling.
    Improved User-Facing Interface: Develop a frontend or CLI client that interacts with the relay more intuitively.

Acknowledgments

    khatru framework for making custom Nostr relay development simpler.
    Intel SGX SDK samples for providing a reference for TEE integration

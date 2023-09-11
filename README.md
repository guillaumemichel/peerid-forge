# Peer ID Forge

> Generate human-readable short libp2p peer IDs for testing convenience.

Author: [Guillaume Michel](https://github.com/guillaumemichel)

Simple command line utility generating short base 58 peer IDs, ending with the provided suffix. The output is a valid libp2p peer ID, but note that there are no cryptographic keypair associated with the peer ID. Peer IDs generated with this tool cannot be used to communicate over libp2p, but are convenient to use in tests. It is convevient to use these peer ids in tests that don't require network communication, because the identifiers are shorter, more human readable, and don't take CPU cycles to generate a cryptographic keypair.

This command-line tool generates short Bitcoin base58-encoded peer IDs with a specified suffix. While the output is a valid libp2p peer ID, it's important to understand that there's no cryptographic keypair linked with these IDs.

## Key Features:

- **Human-friendly**: The shorter peer IDs are easier to read and recognize.
- **Test-ready**: Ideal for unit tests and scenarios that don't involve actual libp2p network communication.
- **Efficiency**: The tool quickly produces peer IDs without the computational overhead of generating a cryptographic keypair.

⚠️ **Caution**: Peer IDs from this tool are not suitable for actual libp2p communications. They're designed primarily for testing purposes.

## Usage

Simply run the program, and provide the desired suffix.

```sh
$ go run . ooooPEER
> 1EooooPEER
```

Note that the suffix should only use Bitcoin Base 58 characters (`123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`).

# Go-Alchemy-SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/teyz/go-alchemy-sdk.svg)](https://pkg.go.dev/github.com/teyz/go-alchemy-sdk) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE) [![Release](https://img.shields.io/github/v/release/Teyz/go-alchemy-sdk)](https://github.com/Teyz/go-alchemy-sdk/releases) [![Tests Passing](https://github.com/teyz/go-alchemy-sdk/actions/workflows/tests.yml/badge.svg)](https://github.com/teyz/go-alchemy-sdk/actions)

`go-alchemy-sdk` is a Go client library for interacting with the [Alchemy API](https://www.alchemy.com/) to access blockchain data and services. This library aims to provide a simple and efficient interface for developers building on top of Solana and other blockchains supported by Alchemy.

## Features

- **Solana Support**: Access Solana blockchain data with ease.
- **Customizable**: Configure client options, including API keys and rate limits.
- **Well-Documented**: Comprehensive documentation and examples.
- **Efficient**: Optimized for performance and minimal overhead.

## Installation

```bash
go get -u github.com/teyz/go-alchemy-sdk
```

## Quick Start

Here’s an example to get started with the `go-alchemy-sdk`:

```go
package main

import (
	"fmt"
	"log"

	alchemy "github.com/teyz/go-alchemy-sdk"
)

func main() {
	// Create a new client with your Alchemy API key
	client, err := alchemy.NewClient("your-alchemy-api-key")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example: Get the balance
	balance, err := client.GetBalance("your-solana-address")
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	fmt.Printf("Balance: %v\n", balance)
}
```

## Documentation

Full API documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/teyz/go-alchemy-sdk).

### Available Methods

- `GetBlock`: Retrieves block.
- `GetTransaction`: Fetches transaction details by hash.
- `GetBalance`: Gets the balance of an Solana address.
- **And more!**

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for any improvements or additional features.

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to [Alchemy](https://www.alchemy.com/) for providing an excellent API for blockchain developers.
- Inspired by libraries like [solana-go](https://github.com/gagliardetto/solana-go).

---

Made with :heart: by [Bastien "Teyz" Rigaud](https://github.com/teyz).

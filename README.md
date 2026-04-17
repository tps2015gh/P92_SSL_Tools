# P92 SSL Tools

A simple Go-based CLI tool to check SSL certificate details for multiple domains.

## Features
- Check expiration date
- Check start date (Not Before)
- Check issuer (3rd party CA)
- Check company (Organization)
- Supports multiple domains as arguments
- Supports `https://` prefix automatically

## Requirements
- Go 1.16 or higher

## Technology Stack
- **Language**: [Go (Golang)](https://golang.org/) - Chosen for its strong standard library support for TLS/SSL and efficient concurrency.
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra) - A powerful library for creating modern CLI applications, providing easy command and argument parsing.
- **Networking/Security**: 
  - `crypto/tls`: Standard Go library used to establish secure connections and inspect peer certificates.
  - `net`: Used for TCP connection handling and timeout management.
- **Build System**: Custom batch script for Windows environments.

## Installation
Clone the repository and run:
```bash
go mod tidy
```

## Usage
Run the tool by passing one or more domains as arguments:
```bash
go run main.go google.com github.com https://microsoft.com
```

### Batch Processing
You can also provide a list of domains in a text file (one URL per line). Lines are automatically trimmed and empty lines are skipped.
```bash
go run main.go --file domains.txt
```

### Build the binary
To build a standalone executable:
```bash
go build -o ssl-check.exe main.go
```
Then run:
```bash
./ssl-check.exe google.com
./ssl-check.exe --file list.txt
```

## Benefits
- **Zero Dependencies**: Compiled Go binary runs without needing any runtime or libraries.
- **Fast Execution**: Uses Go's efficient networking to audit dozens of domains in seconds.
- **Safe Inspection**: Connects via TLS to read public certificate metadata without requiring administrative access.
- **Human-Readable**: Provides a clear summary with status flags (VALID, EXPIRING SOON, EXPIRED).
- **Automation Friendly**: The `--file` flag allows for easy integration into cron jobs or shell scripts.

## Use Cases
- **DevOps/SRE Audits**: Quickly check the status of all microservice endpoints before a release.
- **Compliance Monitoring**: Ensure all public-facing company domains meet security standards (correct CA, valid dates).
- **Pre-migration Checks**: Verify SSL validity after changing DNS records or moving to a new Load Balancer.
- **3rd Party Vendor Tracking**: Keep track of when certificates for external partner APIs are set to expire.

## Team Dev
This project is a collaborative effort between:
- **Human Lead**: [tps2015gh](https://github.com/tps2015gh) (Architecture, Direction, and Review)
- **AI Developer**: **Gemini CLI** (Implementation, Logic Design, and Documentation)

## License
Distributed under the MIT License. See `LICENSE` for more information.

## Example Output
```text
--------------------------------------------------
 Domain:      google.com:443
 Common Name: *.google.com
 Company:     
 Issuer (CA): Google Trust Services
 Start Date:  2024-03-30 08:35:08
 Expiry Date: 2024-06-22 08:35:07
 Days Left:   66 days
 Status:      VALID
--------------------------------------------------
```

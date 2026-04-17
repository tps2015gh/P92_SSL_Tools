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

### Build the binary
To build a standalone executable:
```bash
go build -o ssl-check.exe main.go
```
Then run:
```bash
./ssl-check.exe google.com
```

## Team
This project is a collaborative effort between:
- **Human Developer**: P92 Team / User (Design, Review, and Direction)
- **AI Assistant**: Gemini CLI (Architecture, Implementation, and Documentation)

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

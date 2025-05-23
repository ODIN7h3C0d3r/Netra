# Netra

Netra is a powerful, extensible, and user-friendly CLI tool for IP geolocation, network intelligence, and OSINT research. It enables security professionals, researchers, and network engineers to query IP addresses, analyze metadata, and automate lookups for a variety of use cases.

---

## Table of Contents

- [Features](#features)
- [Demo](#demo)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Usage](#usage)
- [Output Formats](#output-formats)
- [Configuration](#configuration)
- [Advanced Usage](#advanced-usage)
- [Testing](#testing)
- [Development](#development)
- [Troubleshooting](#troubleshooting)
- [FAQ](#faq)
- [License](#license)
- [Author](#author)

---

## Features

- Query geolocation and network metadata for any IP address
- Supports output formats: **text**, **JSON**, **CSV**, **YAML**
- Batch lookup from file
- Save results to file
- Interactive REPL mode for quick lookups
- Customizable output fields
- Caching and retry logic for API requests
- Proxy and custom DNS support
- Modular, extensible Go codebase
- Colorful CLI output and banners
- Open source (MIT License)

---

## Demo

```sh
$ ./netra 8.8.8.8
Ip: 8.8.8.8
Country: United States
City: Mountain View
Region: California
ASN: AS15169
...etc
```

---

## Installation

### Prerequisites

- Go 1.21 or later
- Linux, macOS, or Windows

### Build from Source

```sh
git clone https://github.com/ODIN7h3C0d3r/Netra.git
cd Netra
go build -o netra ./cmd/netra/
```

---

## Quick Start

- **Single IP Lookup:**

  ```sh
  ./netra 8.8.8.8
  ```

- **Batch Lookup:**

  ```sh
  ./netra -file ips.txt
  ```

- **Save Output:**

  ```sh
  ./netra -output results.txt 8.8.8.8 1.1.1.1
  ```

- **Interactive Mode:**

  ```sh
  ./netra --interactive
  ```

---

## Usage

### Command-Line Options

| Option         | Description                                      |
| -------------- | ------------------------------------------------ |
| `-file`        | Path to file containing IPs (one per line)       |
| `-output`      | Save output to file                             |
| `-format`      | Output format: text/json/csv/yaml (default text)|
| `-fields`      | Comma-separated fields to display               |
| `-quiet`       | Suppress progress output                        |
| `-interactive` | Enter interactive mode                          |
| `-help`        | Show help message                               |
| `-version`     | Show version info                               |

### Examples

- Lookup multiple IPs:

  ```sh
  ./netra 8.8.8.8 1.1.1.1
  ```

- Custom output fields:

  ```sh
  ./netra -fields ip,country,asn 8.8.8.8
  ```

- Change output format:

  ```sh
  ./netra -format json 8.8.8.8
  ./netra -format csv -file ips.txt
  ```

---

## Output Formats

- **Text** (default): Human-readable, multi-line per IP
- **JSON**: Machine-readable, suitable for scripting
- **CSV**: For spreadsheets and data analysis
- **YAML**: For config and integration

---

## Configuration

- Edit `config/config.json` to customize:
  - API base URL
  - API token (if required)
  - Retry and timeout settings
  - Output format and fields
  - Proxy and DNS servers
  - UI preferences (color, quiet mode)

---

## Advanced Usage

- **Proxy Support:** Set proxy in `config/config.json` or via environment variable.
- **Custom DNS:** Use custom DNS servers for lookups.
- **Field Filtering:** Display only the fields you care about.
- **Quiet Mode:** Suppress banners and info output with `-quiet`.
- **Integration:** Use JSON/CSV output for automation and pipelines.

---

## Testing

- Run all tests:

  ```sh
  go test ./test/
  ```

- Example test: `TestNetraHelp`, `TestNetraOutputFile`, etc.

---

## Development

- Modular Go codebase under `internal/`
- CLI logic in `internal/cli/`
- Core logic in `internal/core/`
- Network and HTTP in `internal/network/`
- Formatting in `internal/formatter/`
- Utility functions in `internal/util/`
- Tests in `test/`
- Contributions welcome! Fork and submit a PR.

---

## Troubleshooting

- **No output file created?**
  - Place flags before positional arguments: `./netra -output file.txt 8.8.8.8`
- **API errors or rate limits?**
  - Check your API base URL and token in `config/config.json`.
- **Permission errors?**
  - Ensure you have write access to the output directory.
- **Other issues?**
  - Run with `-help` or check logs for more info.

---

## FAQ

**Q: Can I use my own IP geolocation API?**
A: Yes! Edit `config/config.json` to set your own API base URL and token.

**Q: Does Netra support IPv6?**
A: Yes, both IPv4 and IPv6 are supported.

**Q: How do I add new output formats or fields?**
A: Extend the code in `internal/formatter/` and add your logic.

**Q: Is Netra open source?**
A: Yes, MIT License. See [LICENSE](LICENSE).

---

## License

MIT License. See [LICENSE](LICENSE).

## Author

ODIN7h3C0d3r ([github.com/ODIN7h3C0d3r](https://github.com/ODIN7h3C0d3r))

---

*Netra: Network intelligence at your fingertips.*

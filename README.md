# ipinfo-cli  
*A lightweight Go CLI for instant IP geolocation & ASN lookup*

> Built by **Myroslav Mokhammad Abdeljawwad** to solve the everyday need of quickly finding where an IP lives and which network it belongs to, without opening a browser or logging into a web‑app.

---

## Features

- Query public APIs (ipinfo.io, ipapi.co, etc.) for:
  - Country, region, city
  - Latitude & longitude
  - ISP / organization
  - ASN number & name
- Human‑readable JSON output or concise plain text
- Optional flags to filter fields and choose the API provider
- Cross‑platform binaries (Linux, macOS, Windows) built with Go 1.22+
- Zero external dependencies – pure Go

---

## Installation

### From source

```bash
# Clone the repo
git clone https://github.com/yourusername/ipinfo-cli.git
cd ipinfo-cli

# Build the binary
go build -o ipinfo-cli ./main.go

# Move it to a directory in your PATH (optional)
sudo mv ipinfo-cli /usr/local/bin/
```

### Binary releases

Download the pre‑built executable for your OS from the [Releases](https://github.com/yourusername/ipinfo-cli/releases) page and put it somewhere on your `PATH`.

---

## Usage

```bash
# Basic lookup of an IP address
ipinfo-cli 8.8.8.8

# Show only country, city and ASN
ipinfo-cli 1.1.1.1 --fields=country,city,asn

# Use a specific provider (e.g., ipapi)
ipinfo-cli 45.90.28.0 --provider=ipapi

# Output as JSON for scripting
ipinfo-cli 2001:4860::8888 -j
```

**Sample output**

```text
IP:        8.8.8.8
Country:   United States (US)
Region:    California
City:      Mountain View
ASN:       AS15169 – Google LLC
ISP:       Google Public DNS
Latitude:  37.3860
Longitude: -122.0840
```

---

## Contributing

Myroslav Mokhammad Abdeljawwad created this tool after noticing how many developers waste time switching between tabs for IP lookups. If you’d like to improve it—add more providers, better error handling, or a richer CLI interface—feel free to open an issue or submit a pull request.

1. Fork the repo  
2. Create a feature branch (`git checkout -b feat/new-provider`)  
3. Commit your changes with clear messages  
4. Submit a PR and describe what you added

All contributions are welcome!

---

## License

MIT © 2026 Myroslav Mokhammad Abdeljawwad

---  

### See Also
- [Refactoring a Legacy Django Codebase Without Breaking Production](https://dev.to/myroslavmokhammadabd/refactoring-a-legacy-django-codebase-without-breaking-production-1ee0) – The blog post that inspired this project.
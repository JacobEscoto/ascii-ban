# Changelog

## [0.1.0] - 2026-07-23

### Added
- Initial release of ASCII banner generator CLI
- `ascii-ban` command for text-to-ASCII conversion
- `ascii-ban clock` command for live ASCII terminal clock
- `--font` flag to select output style
- Help documentation (`--help`)
- Basic error handling

### Features
- **ASCII Art Banner Generator**: Convert text to large ASCII art banners
- **Live ASCII Clock**: Display time in ASCII format in the terminal
- **Font Selection**: Choose between `standard` and `slant` fonts

### Supported Fonts
- `standard` (default)
- `slant`

### Known Limitations
- Only 2 fonts available in this release
- Clock updates at fixed intervals

### Quick Start
```bash
ascii-ban --help
ascii-ban "banner go" --font slant
ascii-ban clock
```

### Requirements
- Linux/macOS/Windows support

### Feedback
Report issues at: [Github Issues](https://github.com/JacobEscoto/ascii-ban/issues)

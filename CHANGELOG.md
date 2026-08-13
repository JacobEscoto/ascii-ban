# Changelog

## [0.3.1] - 2026-08-13

## Fixed
- **ANSI Export Fix**: Prevented ANSI color codes from being written to output text files when using the `--color` flag alongside export options.

---

## [0.3.0] - 2026-08-12

### Added
- **Text Coloring**: Added `--color` flag to colorize ASCII art outputs using HEX codes, or ANSI values powered by `lipgloss`.
- **Flexible Hex Input**: Automatic detection and prefixing of `#` for HEX color codes (e.g., accepts both `FFF023` and `#FFF023`).

### Refactored
- **Internal Packages**: Consolidated `align` and color utilities into a single `style` package for cleaner layout and maintenance.

---

## [0.2.0] - 2026-07-28

### Added
- **New Font**: Added support for `ansi-shadow` font style.
- **Text Alignment**: Added `--align` flag supporting `left`, `center`, and `right` positioning powered by `lipgloss`.
- **Automatic Formatting**: Implemented `UppercaseOnly` font capability for automatic text conversion on supported fonts like `ansi-shadow`.

---

## [0.1.0] - 2026-07-23

### Added
- Initial release of ASCII banner generator CLI
- `ascii-ban` command for text-to-ASCII conversion
- `ascii-ban clock` command for live ASCII terminal clock
- `--font` flag to select output style
- Help documentation (`--help`)
- Basic error handling

[0.3.0]: https://github.com/JacobEscoto/ascii-ban/releases/tag/v0.3.0
[0.2.0]: https://github.com/JacobEscoto/ascii-ban/releases/tag/v0.2.0
[0.1.0]: https://github.com/JacobEscoto/ascii-ban/releases/tag/v0.1.0

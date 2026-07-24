# ASCII Banner Generator

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

**ascii-ban** (ASCII Banner Generator) is a command-line tool that, as its name suggests, generates ASCII art in the terminal with a provided text.

## Features

- Support for special characters (with accents and diacritics), icons such as the copyright symbol or the eszett character.
- Option to save the generated ASCII banner to a text file using the `--output` or `-o` flags.
- Display of a live clock that updates on the terminal.

## Installation

### Prerequisites

- **Go**: Version 1.21.x or higher
- **Git**

> [!TIP]
> If you use Windows, it is recommended to use [WSL](https://learn.microsoft.com/en-us/windows/wsl/) for installation.

### Build from Source

1. Clone the repository and access the directory

```bash
git clone https://github.com/JacobEscoto/ascii-ban.git
cd ascii-ban
```

2. Download dependencies

```bash
go mod download
```

3. Compile from source code

```bash
go build -o ascii-ban .
```

4. Run the compiled binary to verify that everything works

**Linux / macOS**:
```bash
./ascii-ban "hello world"
```

**Windows**:
```powershell
.\ascii-ban.exe "hello world"
```

5. Install globally

Move the binary to a folder included in your `PATH`:

**Linux / macOS**:

```bash
sudo mv ascii-ban /usr/local/bin
```

**Windows (PowerShell)**:

Open Windows Powershell as Administrator

```powershell
Move-Item .\ascii-ban.exe C:\CLI-Tools\ # (Ensure that C:\CLI-Tools\ is in your PATH)
```

**Alternative with `go install`**:

If you want to use it without manually moving it, you can use:

```bash
go install .
```

## Usage

### Basic Usage
```bash
ascii-ban "Hello World!"
```

### Save to a file
```bash
ascii-ban "Hello World!" --output gen_banner.txt
```

### Clock display
```bash
ascii-ban clock
```

### Choosing a font
```bash
# Uses the 'standard' font by default if not specified
ascii-ban "Banner" --font slant

# Live clock using a custom font
ascii-ban clock --font slant
```

## License

[MIT](https://choosealicense.com/licenses/mit/)



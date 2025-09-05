# CLI Calculator with TUI (Bubble Tea)

A command-line calculator built in Go that starts simple but scales up.  
At its core, it provides a predictable and script-friendly **CLI tool** (`add`, `sub`, `mul`, `div`).  
On top of that, it offers a rich **Text User Interface (TUI)** built with [Bubble Tea](https://github.com/charmbracelet/bubbletea), designed for interactivity, history, and usability — all inside your terminal.

---

## ✨ Features

### Core CLI

- Arithmetic subcommands: `add`, `sub`, `mul`, `div`
- Explicit input/output rules
- Support for integers and floating-point numbers
- Consistent error handling with exit codes
- Cross-platform: Linux, macOS, Windows

### TUI (Bubble Tea)

- Keyboard-driven interactive calculator
- Input box for expressions, evaluated on <kbd>Enter</kbd>
- Expression validation with inline feedback
- Scrollable history of results (with timestamps)
- Copy results to clipboard with a shortcut
- Configurable decimal precision
- Light/dark themes with high-contrast mode
- Pluggable hotkeys via config file

---

## 📸 Screenshots

_Coming soon — examples of the TUI in action._

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.22+](https://go.dev/dl/) installed on your system

### Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/yourusername/cli-calculator-tui.git
cd cli-calculator-tui
go build -o calc .
```

Run the calculator:

```
./calc add 1 2 3
# Output: 6
```

Start the TUI:

```
./calc tui
```

---

## **🛠 Usage**

### **CLI Examples**

```
./calc add 1 2 3       # Output: 6
./calc sub 10 3        # Output: 7
./calc mul 2 3 4       # Output: 24
./calc div 10 4        # Output: 2.5
```

### **TUI Examples**

- Launch with ./calc tui
- Type 3 + 4 \* 2 and press Enter → Result 11
- Navigate history with ↑ / ↓
- Copy a previous result with Ctrl+C

---

## **📂 Project Structure**

```
.
├── cmd/           # CLI entrypoints (subcommands)
├── internal/      # Core compute module (parsing, arithmetic)
├── tui/           # Bubble Tea TUI implementation
├── docs/          # Documentation and screenshots
└── main.go        # Main entrypoint
```

---

## **✅ Roadmap**

- CLI MVP (add, sub, mul, div)
- Error handling and exit codes
- Expression parser (--expr) and stdin mode
- Precision control (--precision)
- Bubble Tea TUI with history, theming, and clipboard integration
- Configurable hotkeys

---

## **🤝 Contributing**

Contributions are welcome!

If you’d like to help, feel free to fork the repo, open a pull request, or suggest new features.

---

## **📜 License**

MIT License © 2025 [Cary O'Neill]

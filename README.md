# Make/Workfront Formatter

A web-based formatter for Make.com and Adobe Workfront Fusion expressions that provides proper indentation and structure to improve readability and maintainability of your automation scripts.

## 🌐 Live Demo

**[View in Action](https://jlentink.github.io/make-workfront-formatter/)**

## ✨ Features

- **Real-time Formatting**: Automatically formats your Make/Fusion expressions as you type
- **Local Storage**: Automatically saves your input to prevent data loss
- **Clean Interface**: Simple, intuitive two-panel layout with input and output
- **Accessibility**: Full screen reader support with proper ARIA labels
- **Mobile Friendly**: Responsive design that works on all devices
- **Fast Performance**: Built with TinyGo WebAssembly for optimal speed and small file size

## 🚀 What It Does

This tool takes complex, nested Make.com and Adobe Workfront Fusion expressions and formats them with proper indentation and structure. For example:

**Input:**
```
{{switch(1; if(sum(5; 3); "yes"; "no"); 3)}}
```

**Output:**
```
{{
  switch(
    1;
    if(
      sum(5; 3);
      "yes";
      "no"
    );
    3
  )
}}
```

## 🛠 Technology Stack

- **Frontend**: HTML5, CSS3, JavaScript
- **Backend Logic**: Go compiled to WebAssembly using TinyGo
- **Build System**: Just (command runner)
- **Deployment**: GitHub Pages

## 🏗 Development

### Prerequisites

- [TinyGo](https://tinygo.org/getting-started/install/) - For compiling Go to WebAssembly
- [Just](https://github.com/casey/just) - Command runner for build tasks
- A local HTTP server (e.g., `http-server` via npm)

### Building

```bash
# Build the project
just build

# Run tests
just test

# Clean build artifacts
just clean

# Build and serve locally
just preview
```

### Project Structure

```
├── src/
│   └── wasm/
│       ├── main.go              # WebAssembly entry point
│       ├── indent_logic.go      # Core formatting logic
│       └── indent_logic_test.go # Unit tests
├── index.html                   # Main web interface
├── justFile                     # Build configuration
├── go.mod                       # Go module definition
└── README.md                    # This file
```

## 🎯 Use Cases

- **Make.com Scenarios**: Format complex automation expressions for better readability
- **Workfront Fusion**: Clean up calculated field expressions and conditional logic
- **Code Review**: Make expressions easier to review and debug
- **Documentation**: Create clean, readable examples for team documentation
- **Learning**: Understand the structure of complex nested expressions

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`just test`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🐛 Issues & Support

If you encounter any issues or have suggestions for improvements, please [open an issue](https://github.com/jlentink/make-workfront-formatter/issues) on GitHub.

---

**Made with ❤️ for the Make.com and Adobe Workfront Fusion community**

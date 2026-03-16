# Contributing to CoinPilot

Thank you for your interest in contributing to CoinPilot! We welcome contributions from everyone.

## 🚀 Getting Started

### Prerequisites

- Go 1.25.6 or later
- Git
- Basic understanding of Go and CLI applications

### Development Setup

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/yourusername/coinpilot.git
   cd coinpilot
   ```
3. **Install dependencies**:
   ```bash
   go mod tidy
   ```
4. **Build the project**:
   ```bash
   make build
   ```
5. **Run tests**:
   ```bash
   make test
   ```

## 📝 How to Contribute

### Reporting Bugs

Before creating bug reports, please check the existing issues to avoid duplicates.

**When filing a bug report, please include:**
- A clear, descriptive title
- Steps to reproduce the issue
- Expected vs actual behavior
- Your operating system and Go version
- Any relevant error messages or logs

### Suggesting Features

We welcome feature suggestions! Please:
- Check existing issues and discussions first
- Provide a clear description of the feature
- Explain the use case and benefits
- Consider the scope and complexity

### Code Contributions

1. **Create a branch** for your feature or fix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following our coding standards

3. **Add tests** for new functionality

4. **Run the test suite**:
   ```bash
   make test
   make test-coverage
   ```

5. **Format your code**:
   ```bash
   make fmt
   ```

6. **Commit your changes** with a clear message:
   ```bash
   git commit -m "Add feature: brief description"
   ```

7. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

8. **Create a Pull Request** on GitHub

## 🎯 Development Guidelines

### Code Style

- Follow standard Go conventions and formatting
- Use `gofmt` to format your code
- Write clear, self-documenting code
- Add comments for complex logic

### Testing

- Write unit tests for new functions
- Maintain or improve test coverage
- Test edge cases and error conditions
- Use table-driven tests where appropriate

### Commit Messages

Use clear, descriptive commit messages:
- Start with a verb in present tense
- Keep the first line under 50 characters
- Add detailed description if needed

Examples:
```
Add CSV import functionality
Fix position calculation for SELL trades
Update README with installation instructions
```

### Pull Request Guidelines

- Keep PRs focused on a single feature or fix
- Include tests for new functionality
- Update documentation as needed
- Ensure all tests pass
- Respond to review feedback promptly

## 🏗️ Project Structure

```
coinpilot/
├── cmd/coinpilot/          # Main application entry
├── internal/               # Private application code
│   ├── cli/               # CLI commands and interface
│   ├── models/            # Data structures
│   ├── services/          # Business logic
│   └── filehandler/       # File I/O operations
├── pkg/utils/             # Public utilities
├── scripts/               # Build and deployment scripts
└── docs/                  # Documentation
```

## 🔄 Development Workflow

### Phase-Based Development

CoinPilot follows a three-phase development approach:

- **Phase 1** (Complete): Basic MVP functionality
- **Phase 2** (In Progress): Extended features and multi-exchange support
- **Phase 3** (Planned): Advanced features and optimizations

### Current Priorities

Check our [GitHub Projects](https://github.com/Crows-Storm/coinpilot/projects) for current development priorities and Phase 2 tasks.

## 📋 Code of Conduct

### Our Standards

- Be respectful and inclusive
- Focus on constructive feedback
- Help others learn and grow
- Maintain a professional tone

### Unacceptable Behavior

- Harassment or discrimination
- Trolling or inflammatory comments
- Personal attacks
- Publishing private information

## 🆘 Getting Help

- **Documentation**: Check the README and code comments
- **Issues**: Search existing GitHub issues
- **Discussions**: Use GitHub Discussions for questions
- **Code Review**: Ask for feedback in your PR

## 📄 License

By contributing to CoinPilot, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to CoinPilot! 🚀
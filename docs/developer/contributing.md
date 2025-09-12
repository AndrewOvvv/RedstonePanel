# Contributing Guide

Thank you for your interest in contributing to RedstonePanel! This guide will help you understand our development process and how to submit contributions.

## Getting Started

### Setting Up Development Environment

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/yourusername/RedstonePanel.git
   cd RedstonePanel
   ```
3. Add upstream remote:
   ```bash
   git remote add upstream https://github.com/original/RedstonePanel.git
   ```

## Development Workflow

### Branching Strategy
We use GitHub Flow:
- `main` branch contains stable code
- `dev` branch for ongoing, feature development and quick fixes

### Making Changes

1. Create a new branch from `dev`:
   ```bash
   git checkout dev
   git pull upstream dev
   git checkout -b feature/your-feature-name
   ```

2. Make your changes
3. Write tests for your changes
4. Run tests and ensure they pass
5. Update documentation if necessary

### Code Standards
- Follow Go best practices and conventions
- Use `gofmt` to format your code
- Write clear, descriptive commit messages

### Testing
- Write unit tests for new functionality
- Ensure all existing tests pass
- Add integration tests where appropriate

### PR Guidelines
1. Create a descriptive title
2. Write clear PR's description
2. Link related issues if exists
3. Request review from maintainers

Contributors will be recognized in:
- CONTRIBUTORS.md file
- Release notes
- Project documentation

# LazyAlias

> 🚀 A smart command-line tool for managing and executing your frequently used commands across different projects with an interactive menu.

## 📋 Table of Contents
- [Why LazyAlias?](#-why-lazyalias)
- [Quick Start](#-quick-start)
- [Installation](#installation)
- [Configuration](#-configuration)
- [Usage](#usage)
- [Features](#-features)
- [Shell Integration](#-shell-integration)
- [Keyboard Shortcuts](#-keyboard-shortcuts)
- [Troubleshooting](#-troubleshooting)
- [Contributing](#-contributing)

## 🤔 Why LazyAlias?

Tired of remembering and typing long commands for your different projects? LazyAlias solves this by providing:

- 🎯 Quick access to your most-used commands
- 🎨 Interactive menu for easy command selection
- 📂 Project-specific command organization
- ⚡️ Support for command arguments and options
- 🔄 Automatic directory switching
- 📋 One-click command execution

Whether you're working with multiple projects or just want to streamline your workflow, LazyAlias makes command execution faster and more intuitive.

## 🚀 Quick Start

```bash
# Install with Homebrew
brew install sergiorivas/lazyalias/lazyalias

# Or install with Go
go install github.com/sergiorivas/lazyalias/cmd/lazyalias@latest

# Create your config file
mkdir -p ~/.config/lazyalias
touch ~/.config/lazyalias/config.yaml

# Run LazyAlias
lazyalias
```

## Installation

#### With brew

```bash
brew install sergiorivas/lazyalias/lazyalias
```

#### With go

```bash
go install github.com/sergiorivas/lazyalias/cmd/lazyalias@latest
```

## 🔧 Configuration

Create a config file at `~/.config/lazyalias/config.yaml`. Here's a basic example:

```yaml
frontend:
  folder: "/projects/frontend"
  commands:
    - name: "Start Dev Server"
      command: "npm run dev"
    - name: "Build"
      command: "npm run build"
    - name: "Test"
      command: "npm run test"

api:
  folder: "/projects/api"
  commands:
    - name: "Run Server"
      command: "go run main.go"
    - name: "Build"
      command: "go build -o api"
    - name: "Test"
      command: "go test ./..."
```

## Usage
##### If you're in a project directory

```bash
lazyalias
```
This will show a menu with the commands configured for that project.

```
[/projects/frontend]% lazyalias
Welcome to LAZYALIAS 🎉🎉🎉
Use the arrow keys to navigate: ↓ ↑ → ←
Select Command
    Start Dev Server
  👉 Build
    Test
    ⬅️ Back to Projects
--------- Command ----------
Name:           Build
Command:        npm run build
```

And then

```
Welcome to LAZYALIAS 🎉🎉🎉
• Selected Command: Build
📋 Command has been copied to clipboard!
💻 Command to execute:
npm run build
```

##### If you're outside project directories
```bash
lazyalias
```
This will first show a menu to select the project, then show its commands. It will automatically change to the project directory before executing the command.

```
[/projects]% lazyalias
Welcome to LAZYALIAS 🎉🎉🎉
Use the arrow keys to navigate: ↓ ↑ → ←
Select Project
  👉 frontend
    api

--------- Project ----------
Name:            frontend
Commands:        3 available
Folder:          /projects/frontend
```

Then

```
Welcome to LAZYALIAS 🎉🎉🎉
• Selected Project: frontend
Use the arrow keys to navigate: ↓ ↑ → ←
Select Command
    Start Dev Server
  👉 Build
    Test
    ⬅️ Back to Projects
--------- Command ----------
Name:           Build
Command:        npm run build
```

And then

```
Welcome to LAZYALIAS 🎉🎉🎉
• Selected Project: frontend
• Selected Command: Build
Command has been copied to clipboard!
cd '/projects/frontend' && npm run build
```

##### After selecting a command
- The command will be copied to your clipboard
- You can paste and execute it in your terminal
- If you're outside the project directory, the command will include the necessary `cd` command

## With Arguments
Commands can include interactive arguments that will be prompted when executing the command:

- Use args to define a list of arguments for a command
- Each argument requires:
  - `name`: Description of what the argument is for
  - `options`: Available options for the argument
      - Use `*` for free text input
      - Use `|` to separate fixed options (e.g., `option1|option2|option3`)
- Reference arguments in commands using `$arg_1`, `$arg_2`, etc.

Example `config.yaml`
```yaml
frontend:
  folder: "/projects/frontend"
  commands:
    - name: "Start Dev Server"
      command: "npm run dev"
    - name: "Build"
      command: "npm run build"
    - name: "Test with Coverage"
      args:
        - name: "Coverage threshold"
          options: "80|85|90|95"
      command: "npm run test -- --coverage-threshold=$arg_1"

api:
  folder: "/projects/api"
  commands:
    - name: "Run Server"
      command: "go run main.go"
    - name: "Build with Tags"
      args:
        - name: "Build tags"
          options: "*"
      command: "go build -tags $arg_1 -o api"
    - name: "Test Package"
      args:
        - name: "Package path"
          options: "*"
        - name: "Test flags"
          options: "-v|-race|-cover"
      command: "go test $arg_2 $arg_1"

docker:
  name: "🐳 Docker"
  commands:
    - name: "Run Container"
      args:
        - name: "Container name"
          options: "*"
        - name: "Port"
          options: "8080|3000|5432"
      command: "docker run -p $arg_2:$arg_2 --name $arg_1 $arg_1"
```

## 🎯 Features

- 🎨 Interactive command selection menu
- 📂 Project-specific command sets
- ⚡️ Command arguments with:
  - Free text input
  - Predefined options
  - Multiple arguments support
- 🔄 Automatic directory switching
- 📋 Command copied to clipboard for easy execution
- 📝 Simple YAML configuration
- 🛠️ Support for any shell command
- 🏷️ Project aliases/names for better organization

## 🔌 Shell Integration

### Ghostty
Add this to your Ghostty config (`~/.config/ghostty/config`):

```
keybind=super+shift+l=text:lazyalias && eval "$(pbpaste)"\n
```

### Fish Shell
Add this to your Fish config (`~/.config/fish/config.fish`):

```fish
bind \cs 'lazyalias && eval (pbpaste)'
```

### Bash/Zsh
Add this to your `.bashrc` or `.zshrc`:

```bash
bind '"\C-s": "lazyalias && eval $(pbpaste)\n"'
```

## 🤝 Contributing

Contributions are welcome! Feel free to:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⭐️ Show Your Support

Give a ⭐️ if this project helped you!

## ⌨️ Keyboard Shortcuts

- `↑/↓`: Navigate through options
- `Enter`: Select option
- `Esc`: Go back/exit
- `Ctrl+C`: Exit at any time

## 🔧 Troubleshooting

### Common Issues

1. **Command not found**
   - Ensure LazyAlias is properly installed
   - Check if the binary is in your PATH
   - Try reinstalling with `brew reinstall lazyalias` or `go install github.com/sergiorivas/lazyalias/cmd/lazyalias@latest`

2. **Config file not found**
   - Verify the config file exists at `~/.config/lazyalias/config.yaml`
   - Check file permissions
   - Ensure YAML syntax is correct

3. **Shell integration not working**
   - Verify the shell config changes are saved
   - Restart your terminal
   - Check if the keybinding conflicts with other applications

### Getting Help

- Check the [GitHub Issues](https://github.com/sergiorivas/lazyalias/issues)
- Open a new issue with:
  - Your OS and version
  - Installation method
  - Error message
  - Steps to reproduce

# LazyAlias

LazyAlias is a command-line tool that helps you manage and execute frequently used commands across different projects. It provides an interactive menu to select commands defined in a YAML configuration file and copies them to your clipboard for easy execution.

## Installation

#### With brew

```bash
brew install sergiorivas/tap/lazyalias
```

#### With go

```bash
go install github.com/sergiorivas/lazyalias/cmd/lazyalias@latest
```

## Configuration
Create a config.yaml file in your home directory `~/.config/lazyalias/config.yaml`:

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

--------- Command ----------
Name:           Build
Command:        npm run build
```

And then

```
Welcome to LAZYALIAS 🎉🎉🎉
• Selected Command: Build

💻 Command to execute:
------------------------
npm run build

📋 Command has been copied to clipboard!
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

--------- Command ----------
Name:           Build
Command:        npm run build
```

And then

```
Welcome to LAZYALIAS 🎉🎉🎉
• Selected Project: frontend
• Selected Command: Build

💻 Command to execute:
------------------------
cd '/projects/frontend' && npm run build

📋 Command has been copied to clipboard!
```

##### After selecting a command
- The command will be copied to your clipboard
- You can paste and execute it in your terminal
- If you're outside the project directory, the command will include the necessary `cd` command

## Features
- Interactive command selection menu
- Project-specific command sets
- Automatic directory switching
- Command copied to clipboard for easy execution
- Simple YAML configuration
- Support for any shell command

## Contributing
- Contributions are welcome! Please feel free to submit a Pull Request.

## License
- MIT License

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

##### Command Arguments
Commands can include interactive arguments that will be prompted when executing the command:

- Use args to define a list of arguments for a command
- Each argument requires:
  - name: Description of what the argument is for
  - options: Available options for the argument
      - Use "*" for free text input
      - Use "|" to separate fixed options (e.g., "option1|option2|option3")
- Reference arguments in commands using $arg_1, $arg_2, etc.

## Features
- Interactive command selection menu
- Project-specific command sets
- Command arguments with:
  - Free text input
  - Predefined options
  - Multiple arguments support
- Automatic directory switching
- Command copied to clipboard for easy execution
- Simple YAML configuration
- Support for any shell command
- Project aliases/names for better organization

## Contributing
- Contributions are welcome! Please feel free to submit a Pull Request.

## License
- MIT License
Example with arguments:

yaml

Copy
sample-with-args:
  name: "Sample Command"
  commands:
    - name: "Run with Args"
      args:
        - name: "Environment"
          options: "dev|staging|prod"
        - name: "Port"
          options: "3000|8080|9000"
        - name: "Custom flag"
          options: "*"
      command: "./run.sh --env $arg_1 --port $arg_2 --flag $arg_3"
Usage
If you're in a project directory
bash

Copy
lazyalias
This will show a menu with the commands configured for that project.

bash

Copy
[/projects/frontend]% lazyalias
Welcome to LAZYALIAS 🎉🎉🎉
Use the arrow keys to navigate: ↓ ↑ → ←
Select Command
    Start Dev Server
  👉 Test with Coverage
    Build

--------- Command ----------
Name:           Test with Coverage
Command:        npm run test -- --coverage-threshold=$arg_1

Select Coverage threshold
  👉 80
    85
    90
    95

And then

Welcome to LAZYALIAS 🎉🎉🎉
• Selected Command: Test with Coverage
• Coverage threshold: 80

💻 Command to execute:
------------------------
npm run test -- --coverage-threshold=80

📋 Command has been copied to clipboard!
If you're outside project directories
bash

Copy
lazyalias
This will first show a menu to select the project, then show its commands. It will automatically change to the project directory before executing the command.

Features
Interactive command selection menu
Project-specific command sets
Command arguments with:
Free text input
Predefined options
Multiple arguments support
Automatic directory switching
Command copied to clipboard for easy execution
Simple YAML configuration
Support for any shell command
Project aliases/names for better organization
Example Configurations
Basic Commands
yaml

Copy
ghostty:
  name: "💻 Ghostty"
  commands:
    - name: "Edit config"
      command: "code ~/.config/ghostty/config"
Commands with Arguments
yaml

Copy
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
Development Tools
yaml

Copy
lazyalias:
  name: "🥴 Lazyalias"
  folder: "/Users/username/projects/lazyalias"
  commands:
    - name: "Open"
      command: "code /Users/username/projects/lazyalias"
    - name: "Edit config"
      command: "code ~/.config/lazyalias/config.yaml"
Contributing
Contributions are welcome! Please feel free to submit a Pull Request.


Copy

Este README mejorado incluye:
- Documentación completa sobre argumentos
- Ejemplos de uso con argumentos
- Más ejemplos de configuración
- Mejor formato y estructura
- Emojis para mejor visualización
- Ejemplos más realistas y prácticos
Share




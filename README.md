# promote-cli

## Usage

### Install

```bash
go install github.com/ahanafy/promote-cli@latest
```

### Config

```yaml
# promote-cli.yaml
gitpath: /path/to/git/repo
environments:
  - development
  - staging
  - production
```

### promote-cli [command]

```bash
A cli tool to check if it is safe to promote to an environment.
        This tool will check if the environment you are promoting to is ahead of the environment you are promoting from.
        To use this tool you need to have a git repository with tags for each environment.

Usage:
  promote-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  promote     The 'promote' subcommand checks if it is safe to promote to an environment.
  view        The 'view' subcommand checks if it is safe to view to an environment.

Flags:
      --config string   config file (default is ./promote-cli.yaml)
  -h, --help            help for promote-cli

Use "promote-cli [command] --help" for more information about a command.
```

### promote

```bash
The 'promote' subcommand checks if it is safe to promote to an environment.

'<cmd> promote' will check if the environment you are promoting to is ahead of the environment you are promoting from.

Usage:
  promote-cli promote [flags]

Flags:
  -c, --check string   Environment to check if promote-cli (required)
  -h, --help           help for promote

Global Flags:
      --config string   config file (default is ./promote-cli.yaml)
```

### view

```bash
The 'view' subcommand checks if it is safe to view to an environment.

'<cmd> view' will check if the environment you are promoting to is ahead of the environment you are promoting from.

Usage:
  promote-cli view [flags]

Flags:
  -h, --help   help for view

Global Flags:
      --config string   config file (default is ./promote-cli.yaml)
```

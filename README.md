# adocli

done using [azuredevops](https://pkg.go.dev/github.com/microsoft/azure-devops-go-api/azuredevops#NewPatConnection) golang package

## Set up before usage

Go to your Azure DevOps console and create a PAT

Export an Environment Variable with the PAT.

On Linux

```sh
export ADOCLI_PAT=$(<~/.adoclipat)
```

On Windows

```pwsh
#
$env:ADOCLI_PAT = Get-Content -Path "$HOME/.adoclipat" -Raw
```

## Build from source

## Install

## Usage

```sh
A command line interface to interact with Azure DevOps

Usage:
  adocli [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  item         Displays azure devops items in a given project
  organization Manages organization options
  projects     handles projects commands
  teams        Handles Azure Devops teams

Flags:
  -h, --help                  help for adocli
  -o, --organization string   organization string to be used in a single run
  -p, --project string        project id string to be used in a single run

Use "adocli [command] --help" for more information about a command.

```

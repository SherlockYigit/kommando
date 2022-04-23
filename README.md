# Kommando

Simple and usable cli tool for go lang.

## Installation

- `go mod init <your project name>`
- `go get github.com/SherlockYigit/kommando`

## Example
```go
import (
    "github.com/SherlockYigit/kommando"
    "fmt"
)

func main() {
    handler := kommando.NewKommando(kommando.KommandoConfig{
		AppName:             "Kommando Test",
		Template:            "Welcome to {AppName}! That's a command list. Type 'help <command name>' to get help with any command.\n{CommandList}",
		CommandListTemplate: "{CommandName} | {CommandDescription}",
		CommandHelpTemplate: "{CommandName} | Help\n{FlagList}\n{CommandAliases}",
		FlagListTemplate:    "--{FlagName} {FlagDescription}",
	})

	handler.AddCommand(kommando.Command{
		Name:        "test",
		Description: "This is a test command.",
		Execute: func(Res *kommando.CommandResponse) {
			fmt.Println("Hello world!")
		},
	})
}
```

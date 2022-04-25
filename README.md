# Kommando

![go-stat](https://goreportcard.com/report/github.com/SherlockYigit/kommando)

Simple and usable cli tool for go lang.

## Installation

- `go mod init <your project name>`
- `go get github.com/SherlockYigit/kommando`

## Example
```go
package main

import (
	"fmt"
	"github.com/SherlockYigit/kommando"
)

func main() {
	handler := kommando.NewKommando(kommando.KommandoConfig{
		AppName:             "Kommando Test",
		Template:            "Welcome to {AppName}! That's a command list. Type 'help <command name>' to get help with any command.\n{CommandList}",
		CommandListTemplate: "{CommandName} | {CommandDescription}",
		CommandHelpTemplate: "{CommandName} | Info\n{CommandDescription}\n{FlagList}\n{CommandAliases}",
		FlagListTemplate:    "--{FlagName} {FlagDescription}",
	})

	handler.AddCommand(kommando.Command{
		Name:        "test",
		Description: "Hello world test example.",
		Execute: func(Res kommando.CommandResponse) {
			fmt.Println("Hello world!")
		},
	})

	handler.Run()
}
```

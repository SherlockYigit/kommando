package kommando

import (
	"fmt"
	"os"
	"strings"
)

func NewKommando(KommandoConf KommandoConfig) KommandoApp {
	append(KommandoConf.Commands, Command{
		Name:        "help",
		Description: "Basic helper command where you can get information about commands.",
		Execute: func(Res *CommandResponse) {
			args := os.Args[1:]

			commandName := args[1]

			for _, val := range KommandoConf.Commands {
				if val.Name == commandName {
					var flaglist []string

					if len(val.Flags) > 0 {

						for _, flagVal := range val.Flags {
							var flag string = strings.Replace(KommandoConf.FlagListTemplate, "{FlagName}", flagVal.Name, -1)
							flag = strings.Replace(flag, "{FlagDescription}", flagVal.Description, -1)

							append(flaglist, flag)
						}
					}

					var helpMessage string = strings.Replace(KommandoConf.CommandHelpTemplate, "{CommandName}", commandName, -1)
					helpMessage = strings.Replace(helpMessage, "{FlagList}", strings.Join(flaglist, "\n"), -1)
					helpMessage = strings.Replace(helpMessage, "{CommandAliases}", strings.Join(val.Aliases, ", "), -1)

					fmt.Println(helpMessage)
				}
			}
		},
	})

	var logmsgCommands []string

	for _, val := range app.Commands {
		var command string = strings.Replace(KommandoConf.CommandListTemplate, "{CommandName}", val.Name, -1)
		command = strings.Replace(command, "{CommandDescription}", val.Description, -1)

		append(logmsgCommands, command)
	}

	var logmsg string = strings.Replace(KommandoConf.Template, "{AppName}", KommandoConf.AppName, -1)
	logmsg = strings.Replace(logmsg, "{CommandList}", strings.Join(logMsgCommands, "\n"), -1)

	fmt.Println(logmsg)

	// Command handler
	args := os.Args[1:]

	for _, cmd := range KommandoConf.Commands {
		if cmd.Name == args[0] {
			cmd.Execute(CommandResponse{
				Command: cmd,
			})
		}
	}

	return KommandoApp{
		KommandoConfig: KommandoConf,
		AddCommand: func(cmd Command) {
			append(KommandoConf.Commands, cmd)
		},
	}
}

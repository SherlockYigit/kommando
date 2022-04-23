package kommando

import (
	"fmt"
	"os"
	"strings"
)

func CreateCommandList(KommandoConf KommandoConfig) {
	var logmsgCommands []string

	for i := 0; i < len(KommandoConf.Commands); i++ {
		val := KommandoConf.Commands[i]

		var command string = strings.Replace(KommandoConf.CommandListTemplate, "{CommandName}", val.Name, -1)
		command = strings.Replace(command, "{CommandDescription}", val.Description, -1)

		logmsgCommands = append(logmsgCommands, command)
	}

	var logmsg string = strings.Replace(KommandoConf.Template, "{AppName}", KommandoConf.AppName, -1)
	logmsg = strings.Replace(logmsg, "{CommandList}", strings.Join(logmsgCommands, "\n"), -1)

	fmt.Println(logmsg)
}

func NewKommando(KommandoConf KommandoConfig) KommandoApp {
	return KommandoApp{
		KommandoConfig: KommandoConf,
		AddCommand: func(cmd Command) {
			KommandoConf.Commands = append(KommandoConf.Commands, cmd)
		},
		Run: func() {
			KommandoConf.Commands = append(KommandoConf.Commands, Command{
				Name:        "help",
				Description: "Basic helper command where you can get information about commands.",
				Execute: func(Res CommandResponse) {
					args := os.Args[2:]

					if len(args) > 0 {
						cmdname := args[0]

						for i := 0; i < len(KommandoConf.Commands); i++ {
							val := KommandoConf.Commands[i]

							if val.Name == cmdname {
								var helpMessage string = strings.Replace(KommandoConf.CommandHelpTemplate, "{CommandName}", cmdname, -1)
								var flaglist []string

								if len(val.Description) > 0 {
									helpMessage = strings.Replace(helpMessage, "{CommandDescription}", fmt.Sprintf("%s", val.Description), -1)
								} else {
									helpMessage = strings.Replace(helpMessage, "{CommandDescription}", "--Doesn't have a description--", -1)
								}

								if len(val.Flags) > 0 {
									for flagI := 0; flagI < len(val.Flags); flagI++ {
										flagVal := val.Flags[flagI]

										var flag string = strings.Replace(KommandoConf.FlagListTemplate, "{FlagName}", flagVal.Name, -1)
										flag = strings.Replace(flag, "{FlagDescription}", flagVal.Description, -1)

										flaglist = append(flaglist, flag)
									}

									helpMessage = strings.Replace(helpMessage, "{FlagList}", strings.Join(flaglist, "\n"), -1)
								} else {
									helpMessage = strings.Replace(helpMessage, "{FlagList}", "--Doesn't have a flag--", -1)
								}

								if len(val.Aliases) > 0 {
									helpMessage = strings.Replace(helpMessage, "{CommandAliases}", strings.Join(val.Aliases, ", "), -1)
								} else {
									helpMessage = strings.Replace(helpMessage, "{CommandAliases}", "--No have aliases--", -1)
								}

								fmt.Println(helpMessage)
							}
						}
					} else {
						CreateCommandList(KommandoConf)
					}
				},
			})

			args := os.Args[1:]

			if len(args) > 0 {
				for i := 0; i < len(KommandoConf.Commands); i++ {
					cmd := KommandoConf.Commands[i]

					if cmd.Name == args[0] {
						cmd.Execute(CommandResponse{
							Command: cmd,
						})
					}
				}
			} else {
				CreateCommandList(KommandoConf)
			}
		},
	}
}

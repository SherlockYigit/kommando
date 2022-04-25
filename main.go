package kommando

import (
	"fmt"
	"os"
	"strings"
)

func parseArgs(args []string) []string {
	var out []string

	for i := 0; len(args) > i; i++ {
		val := args[i]

		if strings.Contains(val, "--") && strings.Contains(val, "=") {
			out = append(out, strings.Split(val, "--")[1])
		} else if strings.Contains(val, "--") {
			if len(args)-1 == i {
				out = append(out, fmt.Sprintf("%s=nil", strings.Contains(val, "--")))
			} else {
				var str string = strings.Split(val, "--")[1]
				str = fmt.Sprintf("%s=%s", str, args[i+1])

				out = append(out, str)
			}
		} else if i == 0 || i > 0 && !strings.Contains(args[i-1], "--") {
			out = append(out, val)
		}
	}

	return out
}

func (KommandoConf *KommandoConfig) AddCommand(cmd Command) {
	KommandoConf.Commands = append(KommandoConf.Commands, cmd)
}

func (KommandoConf *KommandoConfig) Run() {
	args := os.Args[1:]
	conf := KommandoConf

	conf.Commands = append(conf.Commands, Command{
		Name:        "help",
		Description: "Basic helper command where you can get information about commands.",
		Execute: func(Res CommandResponse) {
			args := os.Args[2:]

			if len(args) > 0 {
				var sresult bool = false
				cmdname := args[0]

				for i := 0; i < len(conf.Commands); i++ {
					command := conf.Commands[i]

					if command.Name == cmdname {
						sresult = true

						var helpMessage string = strings.Replace(conf.CommandHelpTemplate, "{CommandName}", cmdname, -1)
						var flaglist []string

						if len(command.Description) > 0 {
							helpMessage = strings.Replace(helpMessage, "{CommandDescription}", fmt.Sprintf("%s", command.Description), -1)
						} else {
							helpMessage = strings.Replace(helpMessage, "{CommandDescription}", "--Doesn't have a description--", -1)
						}

						if len(command.Flags) > 0 {
							for flagI := 0; flagI < len(command.Flags); flagI++ {
								flagVal := command.Flags[flagI]

								var flag string = strings.Replace(conf.FlagListTemplate, "{FlagName}", flagVal.Name, -1)
								flag = strings.Replace(flag, "{FlagDescription}", flagVal.Description, -1)

								flaglist = append(flaglist, flag)
							}

							helpMessage = strings.Replace(helpMessage, "{FlagList}", strings.Join(flaglist, "\n"), -1)
						} else {
							helpMessage = strings.Replace(helpMessage, "{FlagList}", "--Doesn't have a flag--", -1)
						}

						if len(command.Aliases) > 0 {
							helpMessage = strings.Replace(helpMessage, "{CommandAliases}", strings.Join(command.Aliases, ", "), -1)
						} else {
							helpMessage = strings.Replace(helpMessage, "{CommandAliases}", "--No have aliases--", -1)
						}

						fmt.Println(helpMessage)
					} else if !sresult && len(conf.Commands)-1 == i {
						createCommandList(*conf)
					}
				}
			} else {
				createCommandList(*conf)
			}
		},
	})

	if len(args) > 0 {
		for i := 0; i < len(KommandoConf.Commands); i++ {
			cmd := KommandoConf.Commands[i]
			var cmdargs []string = args[1:]

			if cmd.Name == args[0] {
				if len(cmdargs) > 0 {
					cmd.Execute(CommandResponse{
						Command: cmd,
						Args:    parseArgs(cmdargs),
					})
				} else {
					cmd.Execute(CommandResponse{
						Command: cmd,
					})
				}
			}
		}
	} else {
		createCommandList(*KommandoConf)
	}
}

func NewKommando(KommandoConf KommandoConfig) KommandoConfig {
	var conf KommandoConfig = KommandoConf

	return conf
}

func createCommandList(KommandoConf KommandoConfig) {
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

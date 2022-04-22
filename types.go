package kommando

type KommandoConfig struct {
	AppName, Template, CommandListTemplate, CommandHelpTemplate, FlagListTemplate string
	Commands                                                                      []Command
}

// Command Types

type Flag struct {
	Name, Description string
	RequiredValue     bool
}
type SelectedFlag struct {
	Flag
	Value string
}
type CommandResponse struct {
	Command
	SelectedFlags []SelectedFlag
}
type Command struct {
	Name, Description string
	Aliases           []string
	Flags             []Flag
	Execute           func(Res *CommandResponse)
}

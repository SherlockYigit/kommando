package kommando

import (
	"fmt"
	"strings"
)

func NewKommando(KommandoConf KommandoConfig) {
	var logmsg string = strings.Replace(KommandoConf.Template, "{AppName}", KommandoConf.AppName, -1)

	fmt.Println(logmsg)
}

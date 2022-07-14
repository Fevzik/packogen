package generator

import (
	"log"
	"github.com/Fevzik/packogen/constants"
	"github.com/Fevzik/packogen/helper"
	"path"
)

func BuildServer() {
	err := helper.ProcessTemplate(
		"./templates/main_go.tmpl",
		path.Join(projectDir, constants.CmdDirectory, "main.go"),
		map[string]interface{}{
			"repo": InitConf.Repo,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

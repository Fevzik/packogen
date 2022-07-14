package generator

import (
	"github.com/Fevzik/packogen/constants"
	"github.com/Fevzik/packogen/helper"
	"log"
	"path"
)

func BuildConcreteHandler() {
	for _, model := range InitConf.Models {
		data := make(map[string]interface{})
		name := model.Name
		model.formatName()
		data["name"] = model.Name
		data["nameLC"] = name
		data["repo"] = InitConf.Repo
		data["fields"] = model.Fields
		err := helper.ProcessTemplate(
			"./templates/concrete_handler.tmpl",
			path.Join(internalDir, constants.HandlerDirectory, name+".go"),
			data,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

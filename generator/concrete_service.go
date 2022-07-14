package generator

import (
	"log"
	"github.com/Fevzik/packogen/constants"
	"github.com/Fevzik/packogen/helper"
	"path"
)

func BuildModelServices() {
	for _, model := range InitConf.Models {
		data := make(map[string]interface{})
		name := model.Name
		model.formatName()
		data["name"] = model.Name
		data["repo"] = InitConf.Repo
		err := helper.ProcessTemplate(
			"./templates/concrete_service.tmpl",
			path.Join(internalDir, constants.ServiceDirectory, name+".go"),
			data,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

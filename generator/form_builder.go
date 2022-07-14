package generator

import (
	"github.com/Fevzik/packogen/constants"
	"github.com/Fevzik/packogen/helper"
	"path"
)

type FormModel struct {
	Name   string
	Add    bool
	Edit   bool
	Fields map[string]Field
}

func BuildForms() {
	for _, model := range InitConf.Models {
		add := false
		for _, field := range model.Fields {
			if field.Form.Add != nil {
				add = *field.Form.Add
				break
			}
		}

		edit := false
		for _, field := range model.Fields {
			if field.Form.Edit != nil {
				edit = *field.Form.Edit
				break
			}
		}
		if add || edit {
			name := model.Name
			model.formatName()
			fModel := FormModel{
				Name:   model.Name,
				Add:    add,
				Edit:   edit,
				Fields: model.Fields,
			}

			data := map[string]interface{}{
				"data": fModel,
				"repo": InitConf.Repo,
			}

			err := helper.ProcessTemplate(
				"./templates/form_templ.tmpl",
				path.Join(internalDir, constants.FormDirectory, name+".go"),
				data,
			)
			if err != nil {
				panic(err)
				//log.Fatal(err)
			}
		}
	}
}

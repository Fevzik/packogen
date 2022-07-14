package generator

import (
	"github.com/Fevzik/packogen/helper"
	"log"
	"path"
)

type SQlConfDB struct {
	Name     string `yaml:"name"`
	Primary  bool   `yaml:"primary"`
	Index    bool   `yaml:"index"`
	Type     string `yaml:"type"`
	Length   *int   `yaml:"length"`
	Nullable bool   `yaml:"nullable"`
}

type SqlField struct {
	Type string    `yaml:"type"`
	DB   SQlConfDB `yaml:"db"`
	Json string    `yaml:"json"`
}

type SQLModel struct {
	Name   string              `yaml:"name"`
	Fields map[string]SqlField `yaml:"fields"`
}

func BuildSQL() {
	sqlModels := []SQLModel{}
	for _, model := range InitConf.Models {
		sqlModel := SQLModel{
			Name:   model.Name,
			Fields: nil,
		}
		sqlFields := make(map[string]SqlField)
		for k, v := range model.Fields {
			isPrimary := false
			if v.DB.Primary != nil {
				isPrimary = *v.DB.Primary
			}
			isIndex := false
			if v.DB.Index != nil {
				isIndex = *v.DB.Index
			}
			isNullable := false
			if v.DB.Nullable != nil {
				isNullable = *v.DB.Nullable
			}
			sqlField := SqlField{
				Type: v.Type,
				DB: SQlConfDB{
					Name:     v.DB.Name,
					Primary:  isPrimary,
					Index:    isIndex,
					Type:     v.DB.Type,
					Length:   v.DB.Length,
					Nullable: isNullable,
				},
				Json: v.Json,
			}
			sqlFields[k] = sqlField
		}
		sqlModel.Fields = sqlFields
		sqlModels = append(sqlModels, sqlModel)
	}
	data := map[string]interface{}{
		"data": sqlModels,
	}
	err := helper.ProcessSQLTemplate(
		"./templates/sql_templ.tmpl",
		path.Join(projectDir, "initial_migration.sql"),
		data,
	)
	if err != nil {
		log.Fatal(err)
	}
}

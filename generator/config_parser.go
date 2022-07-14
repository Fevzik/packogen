package generator

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
)

/*
Необходимо считать несколько параметров:
корневая директория проекта(rootDir)
наименование проекта(project)
репозиторий(repo)
код сервиса(serviceCode)
наименование сервиса(serviceLabel)
*/

type ConfDB struct {
	Name     string `yaml:"name"`
	Primary  *bool  `yaml:"primary"`
	Index    *bool  `yaml:"index"`
	Type     string `yaml:"type"`
	Length   *int   `yaml:"length"`
	Nullable *bool  `yaml:"nullable"`
}

type Form struct {
	Validation []map[string]interface{}
	Add        *bool `yaml:"add"`
	Edit       *bool `yaml:"edit"`
}

type Field struct {
	Type   string `yaml:"type"`
	DB     ConfDB `yaml:"db"`
	Json   string `yaml:"json"`
	Form   Form   `yaml:"form"`
	Search bool   `yaml:"search"`
}

type Model struct {
	Name   string           `yaml:"name"`
	DBName string           `yaml:"dbName"`
	Fields map[string]Field `yaml:"fields"`
}

type InitialConfig struct {
	RootDir      string           `yaml:"rootDir"`
	Project      string           `yaml:"project"`
	Repo         string           `yaml:"repo"`
	ServiceCode  string           `yaml:"serviceCode"`
	ServiceLabel string           `yaml:"serviceLabel"`
	Models       map[string]Model `yaml:"models"`
}

func ParseConfig(conf string) {
	// получаем абсолютный путь файла конфигурации
	filename, err := filepath.Abs(conf)
	if err != nil {
		log.Fatal(err)
	}
	// читаем файл в массив байт []byte
	confFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// превращаем YML-файл в нашу структуру конфига
	err = yaml.Unmarshal(confFile, &InitConf)
	if err != nil {
		log.Fatal(err)
	}
	projectDir = path.Join(InitConf.RootDir, InitConf.Project)

	// todo YML-file validation
}

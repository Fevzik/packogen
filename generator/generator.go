package generator

import (
	"github.com/Fevzik/packogen/constants"
	"github.com/Fevzik/packogen/helper"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	projectDir  string
	internalDir string
	InitConf    InitialConfig
)

func CleanDir() {
	_, err := os.Stat(projectDir)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll(projectDir)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildDirectories() {
	// create project dir
	err := os.MkdirAll(projectDir, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// create cmd directory
	internalDir = path.Join(projectDir, constants.CmdDirectory)
	err = os.MkdirAll(internalDir, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// create internal directory
	internalDir = path.Join(projectDir, constants.InternalDirectory)
	err = os.MkdirAll(internalDir, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// create directories in internal
	for _, v := range constants.InternalDirectories {
		itemDir := path.Join(internalDir, v)
		err = os.MkdirAll(itemDir, fs.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func EnableGoMod() {
	comm := "cd " + projectDir + " && go mod init " + InitConf.Repo
	err := exec.Command("bash", "-c", comm).Run()
	if err != nil {
		log.Fatal(err)
	}

	// add .gitignore file
	err = exec.Command("bash", "-c", "cp .gitignore "+projectDir).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func BuildModuleGo() {
	data := make(map[string]interface{})
	data["repo"] = InitConf.Repo
	packageName := InitConf.Repo
	parts := strings.Split(packageName, "/")
	if len(parts) > 0 {
		packageName = parts[len(parts)-1]
	}
	data["packageName"] = packageName
	err := helper.ProcessTemplate("./templates/module_templ.tmpl",path.Join(projectDir, "module.go"),data)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildRepositoryHelper() {
	s := []string{}
	for _, model := range InitConf.Models {
		model.formatName()
		s = append(s, model.Name)
	}
	// создадим шаблон абстрактного сервиса и сгенерируем его
	err := helper.ProcessTemplate(
		"./templates/repository_helper.tmpl",
		path.Join(internalDir, constants.RepositoryDirectory, "helper.go"),
		map[string]interface{}{
			"repo": InitConf.Repo,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildRepositoryInterfaces() {
	s := []string{}
	for _, model := range InitConf.Models {
		model.formatName()
		s = append(s, model.Name)
	}
	// создадим шаблон абстрактного сервиса и сгенерируем его
	err := helper.ProcessTemplate(
		"./templates/irepository.tmpl",
		path.Join(internalDir, constants.RepositoryDirectory, "irepository.go"),
		map[string]interface{}{
			"repo":     InitConf.Repo,
			"services": s,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildRepositoryGo() {
	s := []string{}
	for _, model := range InitConf.Models {
		model.formatName()
		s = append(s, model.Name)
	}
	// сгененрируем шаблон абстрактного репозитория
	err := helper.ProcessTemplate(
		"./templates/repository_templ.tmpl",
		path.Join(internalDir, constants.RepositoryDirectory, "repository.go"),
		map[string]interface{}{
			"repo":         InitConf.Repo,
			"repositories": s,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildServiceInterfaces() {
	s := []string{}
	for _, model := range InitConf.Models {
		model.formatName()
		s = append(s, model.Name)
	}
	// создадим шаблон абстрактного сервиса и сгенерируем его
	err := helper.ProcessTemplate(
		"./templates/iservice.tmpl",
		path.Join(internalDir, constants.ServiceDirectory, "iservice.go"),
		map[string]interface{}{
			"repo":     InitConf.Repo,
			"services": s,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildServiceGo() {
	s := []string{}
	for _, model := range InitConf.Models {
		model.formatName()
		s = append(s, model.Name)
	}
	// создадим шаблон абстрактного сервиса и сгенерируем его
	err := helper.ProcessTemplate(
		"./templates/service_templ.tmpl",
		path.Join(internalDir, constants.ServiceDirectory, "service.go"),
		map[string]interface{}{
			"repo":     InitConf.Repo,
			"services": s,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

type Handlers struct {
	NameLc string
	Name   string
}

func BuildHandlerGo() {
	hs := []Handlers{}
	for _, model := range InitConf.Models {
		nameLC := model.Name
		model.formatName()
		hs = append(hs, Handlers{Name: model.Name, NameLc: nameLC})
	}
	// шаблон абстрактного хендлера
	err := helper.ProcessTemplate(
		"./templates/handler_templ.tmpl",
		path.Join(internalDir, constants.HandlerDirectory, "handler.go"),
		map[string]interface{}{
			"repo":        InitConf.Repo,
			"handlers":    hs,
			"serviceCode": InitConf.ServiceCode,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

type ModelsModified struct {
	Name   string
	NameLC string
}

func BuildConstantsGo() {
	data := make(map[string]interface{})
	modelsModified := []ModelsModified{}
	for _, model := range InitConf.Models {
		model.formatName()
		modelsModified = append(modelsModified, ModelsModified{
			Name:   model.Name,
			NameLC: model.DBName,
		})
	}
	data["models"] = modelsModified
	data["serviceCode"] = InitConf.ServiceCode
	data["serviceLabel"] = InitConf.ServiceLabel
	err := helper.ProcessTemplate(
		"./templates/constants_templ.tmpl",
		path.Join(internalDir, constants.ConstantsDirectory, "constants.go"),
		data,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildResponsesGo() {
	err := helper.ProcessTemplate(
		"./templates/responses.tmpl",
		path.Join(internalDir, constants.ResponseDirectory, "response.go"),
		map[string]interface{}{},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildQueueGo() {
	err := helper.ProcessTemplate(
		"./templates/queue_templ.tmpl",
		path.Join(internalDir, constants.QueueDirectory, "queue.go"),
		map[string]interface{}{},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildErrorsGo() {
	// создадим шаблон для собственного пакета ошибок. Здесь мы будем хранить наши уникальные сообщения ошибок
	err := helper.ProcessTemplate(
		"./templates/errors_templ.tmpl",
		path.Join(internalDir, constants.ErrorsDirectory, "errors.go"),
		map[string]interface{}{},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func ModInitProject() {
	comm := "cd " + projectDir + " && go mod tidy "
	err := exec.Command("bash", "-c", comm).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func BuildDomains() {

	for _, model := range InitConf.Models {
		name := model.Name
		model.formatName()
		data := map[string]interface{}{
			"data": model,
		}
		// создадим шаблон для собственного пакета ошибок. Здесь мы будем хранить наши уникальные сообщения ошибок
		err := helper.ProcessTemplate(
			"./templates/structure_templ.tmpl",
			path.Join(internalDir, constants.DomainDirectory, name+".go"),
			data,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

}

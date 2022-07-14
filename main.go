package main

import (
	"flag"
	"github.com/Fevzik/packogen/generator"
)

var (
	configPath = flag.String("cfg", "./test.yml", "configuration file abs path")
)

func main() {
	flag.Parse()
	generator.ParseConfig(*configPath)
	generator.CleanDir()
	generator.BuildDirectories()
	generator.EnableGoMod()
	generator.BuildModuleGo()

	generator.BuildRepositoryHelper()
	generator.BuildRepositoryInterfaces()
	generator.BuildModelRepositories()
	generator.BuildRepositoryGo()

	generator.BuildServiceInterfaces()
	generator.BuildModelServices()
	generator.BuildServiceGo()

	generator.BuildConcreteHandler()
	generator.BuildHandlerGo()

	generator.BuildConstantsGo()
	generator.BuildResponsesGo()

	generator.BuildQueueGo()
	generator.BuildErrorsGo()

	generator.BuildDomains()
	generator.BuildSQL()
	generator.BuildForms()

	generator.ModInitProject()
	generator.BuildServer()
}

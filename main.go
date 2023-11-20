package main

import (
	"github.com/dkZzzz/codecc/chat"
	"github.com/dkZzzz/codecc/config"
	"github.com/dkZzzz/codecc/report"
)

func main() {

	name := "test"
	projectKey := "test"
	branch := "main"
	// sonar_api.CreateProject(name, projectKey)
	// token := sonar_api.GenerateToken(name, projectKey)
	// sonar_api.Scan(".", projectKey, token)
	// sonar_api.GetProjectStatus(projectKey)
	// data := sonar_api.SearchHistory(projectKey)["measures"].([]interface{})
	// for k, v := range data {
	// 	fmt.Println(k, v)
	// }
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	r := report.GenerateReport(name, projectKey, branch)
	r.Print()
	answer := chat.Optimize(config.OpenAISecretKey)
	for _, v := range answer {
		println(v)
	}
}

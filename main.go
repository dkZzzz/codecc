package main

import (
	report "github.com/dkZzzz/codecc/report"
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
	r := report.GenerateReport(name, projectKey, branch)

	r.Print()
}

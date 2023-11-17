package sonarapi

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	request "github.com/dkZzzz/codecc/request"
)

func init() {
}

// create sonar analysis project
// name: project name
// projectKey: project key
// return: map[string]string
func CreateProject(name, projectKey string) map[string]string {
	log.Println("now is creating project")
	url := "http://127.0.0.1:9000/api/projects/create"
	username := "admin"
	password := "123456"
	formData := map[string]string{
		"project":    projectKey,
		"name":       name,
		"mainBranch": "main",
	}
	response := request.FormDataReq(url, username, password, formData)
	log.Println(response["project"].(map[string]string))
	log.Println("create project success")
	return response["project"].(map[string]string)
}

// GenerateToken generate token for sonar analysis
// name: token name
// projectKey: project key
// return: token string
func GenerateToken(name, projectKey string) string {
	log.Println("now is generating token")

	url := "http://127.0.0.1:9000/api/user_tokens/generate"
	username := "admin"
	password := "123456"

	formData := map[string]string{
		"name":       name,
		"type":       "PROJECT_ANALYSIS_TOKEN",
		"projectKey": projectKey,
	}
	response := request.FormDataReq(url, username, password, formData)

	log.Println("generate token success")
	log.Println(response["token"].(string))
	return response["token"].(string)
}

// Get project status
// projectKey: project key
// return: project status
func GetProjectStatus(projectKey string) string {
	log.Println("now is getting project status")

	url := "http://127.0.0.1:9000/api/qualitygates/project_status"
	username := "admin"
	password := "123456"

	queryData := map[string]string{
		"projectKey": projectKey,
	}
	response := request.GET(url, username, password, queryData)

	response = response["projectStatus"].(map[string]interface{})
	log.Println("get project status success")
	log.Println(response["status"].(string))
	return response["status"].(string)
}

// Scan project
// path: project path
// projectKey: project key
// token: token string
func Scan(path, projectKey, token string) {
	cmd := exec.Command("sonar-scanner",
		fmt.Sprintf("-Dsonar.projectKey=%s", projectKey),
		fmt.Sprintf("-Dsonar.sources=%s", path),
		"-Dsonar.host.url=http://localhost:9000",
		fmt.Sprintf("-Dsonar.token=%s", token))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Search history by project key
// projectKey: project key
// return: project status
func SearchHistory(projectKey string) map[string]interface{} {
	log.Println("now is searching history")
	url := "http://127.0.0.1:9000/api/measures/search_history"
	username := "admin"
	password := "123456"

	queryData := map[string]string{
		"component": projectKey,
		"metrics":   "bugs,vulnerabilities,sqale_index,duplicated_lines_density,ncloc,coverage,code_smells,reliability_rating,security_rating,sqale_rating",
		"ps":        "1000",
	}
	response := request.GET(url, username, password, queryData)
	log.Println(response)
	log.Println("search history success")
	return response
}

func IssueSearch(projectKey string) map[string]interface{} {
	log.Println("now is searching issue by component")
	url := "http://127.0.0.1:9000/api/issues/search"
	username := "admin"
	password := "123456"

	queryData := map[string]string{
		"components": projectKey,
	}
	response := request.GET(url, username, password, queryData)
	log.Println(response)
	log.Println("search issue success")
	return response
}

package report

import (
	"fmt"

	sonarapi "github.com/dkZzzz/codecc/sonar_api"
)

func init() {
}

func (r *Report) PaserSearchHistory() {
	SearchHistoryMeasures := sonarapi.SearchHistory(r.ProjectKey)["measures"].([]interface{})

	for _, v := range SearchHistoryMeasures {
		vMap := v.(map[string]interface{})
		switch vMap["metric"].(string) {
		case "bugs":
			r.BugCnt = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		case "code_smells":
			r.CodeSmellCnt = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		case "coverage":
			r.Coverage = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		case "duplicated_lines_density":
			r.DuplicatedLinesDensity = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		case "ncloc":
			r.NclocCnt = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		case "reliability_rating":
			r.ReliabilityRating = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		case "security_rating":
			r.SecurityRating = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		case "sqale_index":
			r.SqaleIndex = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		case "sqale_rating":
			r.SqaleRating = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		case "vulnerabilities":
			r.Vulnerabilities = vMap["history"].([]interface{})[0].(map[string]interface{})["value"].(string)
		}
	}
}

func (r *Report) PaserProjectStatus() {
	r.Status = sonarapi.GetProjectStatus(r.ProjectKey)
}

func (r *Report) PaserIssueSearch() {
	response := sonarapi.IssueSearch(r.ProjectKey)
	r.IssueCnt = fmt.Sprintf("%v", response["total"].(float64))
	r.Issues = make([]Issue, int(response["total"].(float64)))
	// issues := response["issues"]
	for i := 0; i < int(response["total"].(float64)); i++ {
		fmt.Println()
		issue := response["issues"].([]interface{})[i].(map[string]interface{})
		fmt.Println(issue)
		r.Issues[i].Rule = issue["rule"].(string)
		r.Issues[i].Severity = issue["severity"].(string)
		r.Issues[i].Component = issue["component"].(string)
		r.Issues[i].Line = fmt.Sprintf("%v", issue["line"].(float64))
		TextRange := issue["textRange"].(map[string]interface{})
		r.Issues[i].TextRange = make(map[string]int)
		r.Issues[i].TextRange["startLine"] = int(TextRange["startLine"].(float64))
		r.Issues[i].TextRange["endLine"] = int(TextRange["startLine"].(float64))
		r.Issues[i].TextRange["startOffset"] = int(TextRange["startLine"].(float64))
		r.Issues[i].TextRange["endOffset"] = int(TextRange["startLine"].(float64))
		r.Issues[i].Status = issue["status"].(string)
		r.Issues[i].Message = issue["message"].(string)
		r.Issues[i].Tag = issue["tags"].([]interface{})[0].(string)
		r.Issues[i].Type = issue["type"].(string)
		r.Issues[i].Impact = make(map[string]string)
		r.Issues[i].Impact["softwareQuality"] = issue["impacts"].([]interface{})[0].(map[string]interface{})["softwareQuality"].(string)
		r.Issues[i].Impact["severity"] = issue["impacts"].([]interface{})[0].(map[string]interface{})["severity"].(string)
		fmt.Println()
	}
	fmt.Println(r.Issues)
}

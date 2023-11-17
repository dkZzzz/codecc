package report

import (
	"fmt"
	"log"
	"strconv"
)

type Report struct {
	ProjectKey             string  // 项目key
	ProjectName            string  // 项目名称
	Branch                 string  // 分支
	Status                 string  // 项目状态
	BugCnt                 string  // 功能性缺陷数
	CodeSmellCnt           string  // 代码异味数
	Coverage               string  // 测试覆盖率
	NclocCnt               string  // 未提交代码行数
	DuplicatedLinesDensity string  // 重复行密度
	ReliabilityRating      string  // 可靠性
	SecurityRating         string  // 安全性
	SqaleIndex             string  // 解决组件上的所有问题并遵守所有要求的总工作量（单位：天）
	SqaleRating            string  // 技术负债率评级
	Vulnerabilities        string  // 安全性漏洞数
	IssueCnt               string  // 问题数
	Issues                 []Issue // 问题列表
}

func init() {
}

func GenerateReport(projectName, projectKey, Branch string) *Report {
	r := new(Report)
	r.Issues = make([]Issue, 0)
	r.ProjectKey = projectKey
	r.ProjectName = projectName
	r.Branch = Branch
	r.PaserSearchHistory()
	r.PaserProjectStatus()
	r.PaserIssueSearch()
	return r
}

func (r *Report) Print() {
	fmt.Println("======================Report=========================")
	fmt.Println("ProjectKey: ", r.ProjectKey)
	fmt.Println("ProjectName: ", r.ProjectName)
	fmt.Println("Branch: ", r.Branch)
	fmt.Println("Status: ", r.Status)
	fmt.Println("BugCnt: ", r.BugCnt)
	fmt.Println("CodeSmellCnt: ", r.CodeSmellCnt)
	fmt.Println("Coverage: ", r.Coverage)
	fmt.Println("NclocCnt: ", r.NclocCnt)
	fmt.Println("DuplicatedLinesDensity: ", r.DuplicatedLinesDensity)
	fmt.Println("ReliabilityRating: ", r.ReliabilityRating)
	fmt.Println("SecurityRating: ", r.SecurityRating)
	fmt.Println("SqaleIndex: ", r.SqaleIndex)
	fmt.Println("SqaleRating: ", r.SqaleRating)
	fmt.Println("Vulnerabilities: ", r.Vulnerabilities)
	fmt.Println("IssueCnt: ", r.IssueCnt)

	tmp, err := strconv.Atoi(r.IssueCnt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	for i := 0; i < tmp; i++ {
		fmt.Printf("=====Issues[%d]=====\n", i)
		r.Issues[i].Print()
		fmt.Println()
	}
	fmt.Print("======================================================")
}

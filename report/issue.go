package report

import "fmt"

func init() {
}

type Issue struct {
	Rule      string            // 规则
	Severity  string            // 问题等级
	Component string            // 文件路径
	Line      string            // 行号
	TextRange map[string]int    // 问题描述
	Status    string            // 问题状态
	Message   string            // 问题描述
	Tag       string            // 问题标签
	Type      string            // 问题类型
	Impact    map[string]string // 问题影响
}

func (i *Issue) Print() {
	fmt.Println("Rule: ", i.Rule)
	fmt.Println("Severity: ", i.Severity)
	fmt.Println("Component: ", i.Component)
	fmt.Println("Line: ", i.Line)
	fmt.Println("TextRange: ", i.TextRange)
	fmt.Println("Status: ", i.Status)
	fmt.Println("Message: ", i.Message)
	fmt.Println("Tag: ", i.Tag)
	fmt.Println("Type: ", i.Type)
	fmt.Println("Impact: ", i.Impact)
}

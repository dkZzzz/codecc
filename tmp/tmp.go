package tmp

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
)

type CloneProjectRequest struct {
	ProjectURL  string `json:"project_url"`
	ProjectName string `json:"project_name"`
}

func cloneHandler(c *gin.Context) {
	var request CloneProjectRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求参数"})
		return
	}
	// 执行 `git clone` 命令
	targetDir := "/Users/danko/Downloads/" + request.ProjectName
	_, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL:      request.ProjectURL,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Git存储库克隆成功!"})
}

func main() {
	r := gin.Default()
	r.POST("/clone_project", cloneHandler)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("error")
	}
}

package logic

import (
	"fmt"
	"os"
	"path/filepath"

	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/pkg/utils"
	"github.com/issueye/code_magic/backend/service"
)

type Template struct{}

// 保存代码到文件
func (t *Template) SaveCodeToFile(tpId string, code string) error {
	srv := commonService.NewService(&service.Template{})
	info, err := srv.Get(tpId)
	if err != nil {
		return err
	}

	// 获取文件名称和文件类型
	fileName := fmt.Sprintf("%s.js", info.FileName)

	path := filepath.Join("runtime", "static", "code")
	return os.WriteFile(filepath.Join(path, fileName), []byte(code), 0644)
}

// 获取文件内容
func (t *Template) GetCode(tpId string) (string, error) {
	srv := commonService.NewService(&service.Template{})
	info, err := srv.Get(tpId)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s.js", info.FileName)
	path := filepath.Join("runtime", "static", "code", fileName)

	exist := utils.FileExist(path)
	fmt.Println("exist", exist)
	if !exist {
		return "", nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

package Parse

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bingxindan/bxd_go_lib/logger"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

type Parse struct {
}

func (p *Parse) ParseToml(ctx context.Context) {
	var (
		tag = "TomlParse"
	)

	// 解析json文件，获取toml文件名称
	activeFile, err := p.getActiveFiles("Conf/ActiveFile.json")
	if err != nil {
		logger.Ex(ctx, tag, "ParseToml, err: %+v", err)
		return
	}

	// 通过","，分隔，得到每个toml文件名称
	tomlFiles := strings.Split(activeFile, ",")
	if len(tomlFiles) == 0 {
		logger.Ex(ctx, tag, "Split, tomlFiles is null")
		return
	}

	// 解析每个toml文件
	for i := 0; i < len(tomlFiles); i++ {
		tomlFile := tomlFiles[i]
		fmt.Println(tomlFile)
	}
}

type ActiveFile struct {
	Files string `json:"files"`
}

func (p *Parse) getActiveFiles(filePath string) (string, error) {
	if filePath == "" {
		return "", errors.New("文件地址为空")
	}

	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	//读取的数据为json格式，需要进行解码
	activeFile := ActiveFile{}
	err = json.Unmarshal(data, &activeFile)
	if err != nil {
		return "", err
	}

	return activeFile.Files, nil
}

package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	//设置配置名称
	vp.SetConfigName("config")
	//添加配置文件目录
	vp.AddConfigPath("configs/")
	//指定配置文件格式
	vp.SetConfigType("yaml")
	//读取配置文件
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	//将读取的配置文件进行输出 可省略
	err = vp.SafeWriteConfigAs("configs/output.yaml")
	if err != nil {
		fmt.Println(err)
	}
	//返回setting构造
	return &Setting{vp}, nil
}

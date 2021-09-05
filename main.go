package main

import (
	"fmt"
	"go-blog-server/global"
	"go-blog-server/internal/routers"
	"go-blog-server/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
	    log.Fatalf("init setupSetting Error : %v",err)
	}
}

func main() {
	//实例化路由处理程序
	router := routers.NewRouter()
	// 自定义http Server
	s := &http.Server {
		//设置监听端口
	    Addr : ":" + global.ServerSetting.HttpPort,
	    //添加我们写好的处理程序
	    Handler: router,
	    //设置允许读取写入的最大时间
	    ReadTimeout : global.ServerSetting.ReadTimeout,
	    WriteTimeout : global.ServerSetting.WriteTimeout,
	    //设置请求头最大字节数 2^20bytes 1Mb
	    MaxHeaderBytes: 1<<20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}


//setupSetting 安装配置
func setupSetting() error{
	mySetting,err := setting.NewSetting()
	if err != nil {
	    return err
	}
	err = mySetting.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}
	err = mySetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = mySetting.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}

	//将数据设置为秒单位
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
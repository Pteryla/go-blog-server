package main

import (
	"fmt"
	"myblog-server/internal/routers"
	"net/http"
	"time"
)

func main() {
	//实例化路由处理程序
	router := routers.NewRouter()
	// 自定义http Server
	s := &http.Server {
		//设置监听端口
	    Addr : ":8088",
	    //添加我们写好的处理程序
	    Handler: router,
	    //设置允许读取写入的最大时间
	    ReadTimeout : 10 * time.Second,
	    WriteTimeout : 10*time.Second,
	    //设置请求头最大字节数 2^20bytes 1Mb
	    MaxHeaderBytes: 1<<20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

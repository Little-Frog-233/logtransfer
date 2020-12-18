package main

import (
	"fmt"
	"logtransfer/conf"
	"logtransfer/es"
	"logtransfer/kafka"
	"sync"

	"gopkg.in/ini.v1"
)

// 将日志数据从kafka中取出来发往ES

var (
	wg  sync.WaitGroup
	cfg = new(conf.LogTransfer)
)

func main() {
	wg.Add(1)
	// 0. 加载配置文件
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Println("load ini failed, err:", err)
		return
	}
	// 1. 初始化
	// 1.1 初始化es
	err = es.Init(cfg.ESCfg.Address)
	if err != nil {
		fmt.Println("init elasticsearch failed, err:", err)
		return
	}
	fmt.Println("init elasticsearch success")

	// 1.2 初始化kafka
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	// err = kafka.Init()
	if err != nil {
		fmt.Println("init kafka consmuer failed, err:", err)
		return
	}
	fmt.Println("init kafka success")
	wg.Wait()
}

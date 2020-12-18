package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

// 初始化es，准备接受kafka中发送来的数据

// LogData ...
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     = make(chan *LogData, 1000)
)

// Init 初始化elasticsearch
func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(address))
	if err != nil {
		return
	}

	fmt.Println("connect to es success")
	go sendToES()
	return
}

// SendToESChan 往通道中发送数据
func SendToESChan(msg *LogData) {
	ch <- msg
}

// SendToES 发送数据到ES
func sendToES() {
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().
				Index(msg.Topic). //数据库
				Type("log").      //表
				BodyJson(&msg).
				Do(context.Background())
			if err != nil {
				// Handle error
				fmt.Println(err)
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}

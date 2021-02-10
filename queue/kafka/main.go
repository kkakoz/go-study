package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main()  {
	// 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // ack确认机制
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 选择分区-随机分区
	config.Producer.Return.Successes = true // 确认

	// 封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "demo1"
	msg.Value = sarama.StringEncoder("第一个消息")

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"47.98.156.22:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
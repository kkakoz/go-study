package demo2

import (
	"github.com/Shopify/sarama"
	"log"
	"time"
)

func OffsetManagerConsumer(topic string) {
	config := sarama.NewConfig()
	// 配置开启自动提交 offset，这样 samara 库会定时帮我们把最新的 offset 信息提交给 kafka
	config.Consumer.Offsets.AutoCommit.Enable = true              // 开启自动 commit offset
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // 自动 commit时间间隔
	client, err := sarama.NewClient([]string{addr}, config)
	if err != nil {
		log.Fatal("NewClient err: ", err)
	}
	defer client.Close()
	// offsetManager 用于管理每个 consumerGroup的 offset
	// 根据 groupID 来区分不同的 consumer，注意: 每次提交的 offset 信息也是和 groupID 关联的
	offsetManager, _ := sarama.NewOffsetManagerFromClient("myGroupID", client) // 偏移量管理器
	defer offsetManager.Close()
	// 每个分区的 offset 也是分别管理的，demo 这里使用 0 分区，因为该 topic 只有 1 个分区
	partitionOffsetManager, _ := offsetManager.ManagePartition(topic, 0) // 对应分区的偏移量管理器
	defer partitionOffsetManager.Close()
	// defer 在程序结束后在 commit 一次，防止自动提交间隔之间的信息被丢掉
	defer offsetManager.Commit()
	consumer, _ := sarama.NewConsumerFromClient(client)
	// 根据 kafka 中记录的上次消费的 offset 开始+1的位置接着消费
	nextOffset, _ := partitionOffsetManager.NextOffset() // 取得下一消息的偏移量作为本次消费的起点
	pc, _ := consumer.ConsumePartition(topic, 0, nextOffset)
	defer pc.Close()

	for message := range pc.Messages() {
		value := string(message.Value)
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, value)
		// 每次消费后都更新一次 offset,这里更新的只是程序内存中的值，需要 commit 之后才能提交到 kafka
		partitionOffsetManager.MarkOffset(message.Offset+1, "modified metadata") // MarkOffset 更新最后消费的 offset
	}
}
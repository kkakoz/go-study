package demo2

import (
	"github.com/Shopify/sarama"
	"log"
	"sync"
)

func SinglePartition(topic string) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{addr}, config)
	if err != nil {
		log.Fatalln("new consumer err:", err)
	}
	defer consumer.Close()
	// 参数1 指定消费哪个 topic
	// 参数2 分区 这里默认消费 0 号分区 kafka 中有分区的概念，类似于ES和MongoDB中的sharding，MySQL中的分表这种
	// 参数3 offset 从哪儿开始消费起走，正常情况下每次消费完都会将这次的offset提交到kafka，然后下次可以接着消费，
	// 这里demo就从最新的开始消费，即该 consumer 启动之前产生的消息都无法被消费
	// 如果改为 sarama.OffsetOldest 则会从最旧的消息开始消费，即每次重启 consumer 都会把该 topic 下的所有消息消费一次
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("consume partition err:", err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value:%s\n", message.Partition, message.Offset, message.Value)
	}
}

func Partitions(topic string) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{addr}, config)
	if err != nil {
		log.Fatalln("new consumer err:", err)
	}
	// 先查询该 topic 有多少分区
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatal("Partitions err: ", err)
	}
	var wg sync.WaitGroup
	// 然后每个分区开一个 goroutine 来消费
	for _, partitionId := range partitions {
		wg.Add(1)
		go consumeByPartition(topic, consumer, partitionId, &wg)
	}
	wg.Wait()
}

func consumeByPartition(topic string, consumer sarama.Consumer, partitionId int32, wg *sync.WaitGroup) {
	defer wg.Done()
	partitionConsumer, err := consumer.ConsumePartition(topic, partitionId, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ConsumePartition err: ", err)
	}
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}
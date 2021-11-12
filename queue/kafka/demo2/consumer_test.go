package demo2

import (
	"testing"
)

func TestAsyncPoduct(t *testing.T)  {
	AsyncProducer("test2", 3)
}

func TestSyncProducer(t *testing.T) {
	SyncProducer("test3", 2)
}

func TestSinglePartition(t *testing.T) {
	SinglePartition("test2")
	//go demo1.Consumer("test2")
	//AsyncProducer("test2", 3)
}

func TestPartitions(t *testing.T)  {
	Partitions("test2")
}

func TestOffsetManagerConsumer(t *testing.T) {
	OffsetManagerConsumer("test3")
}

func TestConsumerGroup(t *testing.T) {
	go ConsumerGroup("test2", "group1", "consumer1")
	go ConsumerGroup("test2", "group2", "consumer2")
	select {

	}
}
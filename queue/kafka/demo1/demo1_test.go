package demo1

import "testing"

func TestConsumer(t *testing.T) {
	Consumer("test1")
}

func TestProducer(t *testing.T) {
	Produce("test1", 10)
}
package demo2

import (
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	addr = "127.0.0.1:9092"
)

// 异步生产者
func AsyncProducer(topic string, limit int) {
	config := sarama.NewConfig()
	// 异步生产者不建议把 Errors 和 Successes 都开启，一般开启 Errors 就行
	// 同步生产者就必须都开启，因为会同步返回发送成功或者失败
	config.Producer.Return.Errors = false   // 设定需要返回错误信息
	config.Producer.Return.Successes = true // 设定需要返回成功信息
	producer, err := sarama.NewAsyncProducer([]string{addr}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.AsyncClose()
	var count int64
	go func() {
		for {
			select {
			case s, ok := <-producer.Successes():
				if !ok {
					return
				}
				log.Printf("[Producer] key:%v msg:%+v \n", s.Key, s.Value)
			case e := <-producer.Errors():
				if e != nil {
					log.Printf("[Producer] err:%v msg:%+v \n", e.Msg, e.Err)
				}
			}
		}
	}()
	// 异步发送
	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(str)}
		// 异步发送只是写入内存了就返回了，并没有真正发送出去
		// sarama 库中用的是一个 channel 来接收，后台 goroutine 异步从该 channel 中取出消息并真正发送
		producer.Input() <- msg
		atomic.AddInt64(&count, 1)
		if atomic.LoadInt64(&count)%1000 == 0 {
			log.Printf("已发送消息数:%v\n", count)
		}

	}
	time.Sleep(1 * time.Second) // 等待异步发送成功
	log.Printf("发送完毕 总发送消息数:%v\n", limit)
}

func SyncProducer(topic string, limit int) {
	config := sarama.NewConfig()
	// 同步生产者必须同时开启 Return.Successes 和 Return.Errors
	// 因为同步生产者在发送之后就必须返回状态，所以需要两个都返回
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true // 这个默认值就是 true 可以不用手动 赋值
	// 同步生产者和异步生产者逻辑是一致的，Success或者Errors都是通过channel返回的，
	// 只是同步生产者封装了一层，等channel返回之后才返回给调用者
	// 具体见 sync_producer.go 文件72行 newSyncProducerFromAsyncProducer 方法
	// 内部启动了两个 goroutine 分别处理Success Channel 和 Errors Channel
	// 同步生产者内部就是封装的异步生产者
	// type syncProducer struct {
	// 	producer *asyncProducer
	// 	wg       sync.WaitGroup
	// }
	producer, err := sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()
	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(str)}
		partition, offset, err := producer.SendMessage(msg) // 发送逻辑也是封装的异步发送逻辑，可以理解为将异步封装成了同步
		if err != nil {
			log.Println("SendMessage err: ", err)
			return
		}
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, str)
	}
}
package kafka

import (
	"fmt"
	"testing"

	"github.com/IBM/sarama"
)

// kafka 消费示例

func TestKafka(t *testing.T) {
	broker := []string{"127.0.0.1:9092"}
	topic := "system_log"

	client, err := sarama.NewConsumer(broker, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	fmt.Println("connect kafka success")

	partitionList, err := client.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}

	fmt.Printf("topic:%s, partition:%v\n", topic, partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		cp, err := client.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer cp.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range cp.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(cp)
		select {}
	}
}

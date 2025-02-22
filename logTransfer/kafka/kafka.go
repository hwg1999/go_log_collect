package kafka

import (
	"fmt"
	"sync"

	"github.com/IBM/sarama"
)

// Init 初始化kafka连接，准备发送数据给es
func Init(addr []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}

	fmt.Printf("topic:%s, partition:%v\n", topic, partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		var cp sarama.PartitionConsumer
		cp, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}

		// 异步从每个分区消费信息
		var wg sync.WaitGroup
		defer wg.Done()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range cp.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 准备发送数据给es
			}
		}(cp)

		wg.Wait()
	}

	return
}

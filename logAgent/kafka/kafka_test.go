package kafka

import (
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
)

func TestKafkaProducer(t *testing.T) {
	// 配置 Kafka 生产者
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 需要所有副本确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机选择分区
	config.Producer.Return.Successes = true                   // 成功交付的消息返回到 success channel

	// 构造消息
	msg := &sarama.ProducerMessage{
		Topic: "web_log",
		Value: sarama.StringEncoder("this is a test log"), // 消息内容
	}

	// 连接 Kafka 生产者
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:29092"}, config)
	assert.NoError(t, err, "Failed to start Kafka producer") // 使用assert检查错误

	defer func() {
		err := client.Close()
		assert.NoError(t, err, "Failed to close Kafka producer") // 使用assert检查关闭是否有错误
	}()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	assert.NoError(t, err, "Failed to send message")

	// 断言消息发送成功
	t.Logf("Message sent successfully! Partition: %d, Offset: %d\n", pid, offset)
	assert.Greater(t, pid, int32(0), "Partition should be greater than 0")
	assert.Greater(t, offset, int64(0), "Offset should be greater than 0")
}

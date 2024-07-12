package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Topic         string
	SarmaProducer sarama.SyncProducer
}

func NewKafkaProducer(urls []string, topic string) (*KafkaProducer, error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(urls, config)
	if err != nil {
		return nil, err
	}

	kp := KafkaProducer{topic, producer}
	return &kp, nil
}

func (k *KafkaProducer) ProduceMessage(m string) {
	msg := &sarama.ProducerMessage{
		Topic: k.Topic,
		Value: sarama.StringEncoder(m),
	}
	_, _, _ = k.SarmaProducer.SendMessage(msg)

	fmt.Println("Sending to topic with offset ", msg.Offset)

}

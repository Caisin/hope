package provider

import (
	"github.com/nsqio/go-nsq"
	"hope/pkg/conf"
)

func NewNsqProducer(c *conf.Data) *nsq.Producer {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(c.Nsq.NsqEndpoint, config)
	if err != nil {
		panic(err)
	}
	return producer
}

//NewNsqConsumer channel作用是用来做消息复制的
func NewNsqConsumer(topic, channel, lookupEndpoint string, handler nsq.Handler) (*nsq.Consumer, error) {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, err
	}
	consumer.AddHandler(handler)
	err = consumer.ConnectToNSQLookupd(lookupEndpoint)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

package events

import (
	"github.com/NubeIO/lib-redis/libredis"
	"github.com/NubeIO/rubix-automater/automater/model"
	log "github.com/sirupsen/logrus"
)

const automater = "automater-transaction"

func (inst *Events) initRedis() libredis.Client {
	inst.Addr = ""
	client, err := libredis.New(&libredis.Config{})
	if err != nil {
		log.Errorln("eventbus-initRedis()", err)
	}
	return client
}

func (inst *Events) redisSubscriber() {
	messages := make(chan string)
	go func() {
		for {
			msg := <-messages
			data := &model.PublishTransaction{}
			err := inst.redis.Decode(msg, data)
			if err != nil {
				log.Errorln("eventbus-decode-redisSubscriber()", err)
			}
			err = inst.EventBus.Emit(automater, data)
			if err != nil {
				log.Errorln("eventbus-emit-redisSubscriber()", err)
			}
		}
	}()
	err := inst.redis.Subscribe(automater, messages)
	if err != nil {
		log.Errorln("eventbus-redis-subscribe-redisSubscriber()", err)
	}
}

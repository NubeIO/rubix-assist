package events

import (
	"github.com/NubeIO/lib-bus/eventbus"
	"github.com/NubeIO/lib-redis/libredis"
)

type Events struct {
	Addr     string
	redis    libredis.Client
	EventBus *eventbus.EventBus
}

func New(events *Events) *Events {
	events.InitEventBus()
	events.InitRedis()
	go events.redisSubscriber()
	return events

}
func (inst *Events) InitEventBus() {
	b, _ := eventbus.New()
	inst.EventBus = b
}

func (inst *Events) InitRedis() {
	inst.redis = inst.initRedis()
}

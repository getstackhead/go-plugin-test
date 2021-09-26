package pluginlib

import "github.com/asaskevich/EventBus"

var bus = EventBus.New()

func Subscribe(eventName string, fn interface{}) error {
	return bus.Subscribe(eventName, fn)
}

func Unsubscribe(eventName string, fn interface{}) error {
	return bus.Unsubscribe(eventName, fn)
}

func Publish(eventName string, args ...interface{}) {
	bus.Publish(eventName, args)
}

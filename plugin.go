package pluginlib

type Plugin interface {
	LaunchServer(callback func(int))
}

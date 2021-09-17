package pluginlib

type Plugin interface {
	Deploy(project Project)
	Setup()
}

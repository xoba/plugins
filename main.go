package plugins

type Interface interface {
	Run(string) (string, error)
}

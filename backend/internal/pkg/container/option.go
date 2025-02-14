package container

type FilterContainerOption struct {
	Name string
	ID   string
	Port string
}

type CreateContainerOption struct {
	Name    string
	Port    string
	Image   string
	Command []string
	Labels  map[string]string
}

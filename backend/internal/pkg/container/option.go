package container

type FilterContainer struct {
	Name string
	ID   string
	Port string
}

type CreateContainer struct {
	Name    string
	Port    string
	Image   string
	Command []string
	Labels  map[string]string
}

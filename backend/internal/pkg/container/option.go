package container

type FilterContainerOption struct {
	Name string
	ID   string
	Port string
}

type CreateContainerOption struct {
	Name            string
	Port            string
	Image           string
	Command         []string
	VolumeName      string
	Labels          map[string]string
	Entrypoint      []string
	Env             []string
	HostConfigBinds []string
	WorkingDir      string
}

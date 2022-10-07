package configurations

type Configuration interface {
	Register()
}

type Configurations []Configuration

func (c Configurations) Init() {
	for _, config := range c {
		config.Register()
	}
}

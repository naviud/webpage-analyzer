package configurations

type Configuration interface {
	Register()
}

type Configurations []Configuration

// Init function is responsible to call the
// Register function of all the Configuration
// objects.
func (c Configurations) Init() {
	for _, config := range c {
		config.Register()
	}
}

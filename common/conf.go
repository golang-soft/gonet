package common

type (
	MServer struct {
		Endpoints []string `yaml:"endpoints"`
	}

	Server struct {
		Ip   string `yaml:"ip"`
		Port int    `yaml:"port"`
	}

	Db struct {
		Ip           string `yaml:"ip"`
		Name         string `yaml:"name"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
	}

	Redis struct {
		OpenFlag bool   `yaml:"open"`
		Ip       string `yaml:"ip"`
		Password string `yaml:"password"`
	}

	Etcd struct {
		Endpoints []string `yaml:"endpoints"`
	}

	SnowFlake struct {
		Endpoints []string `yaml:"endpoints"`
	}

	Nats struct {
		Endpoints string `yaml:"endpoints"`
	}

	Raft struct {
		Endpoints []string `yaml:"endpoints"`
	}

	Center struct {
		Ip   string `yaml:"ip"`
		Port int    `yaml:"port"`
	}
)

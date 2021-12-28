package common

type (
	MServer struct {
		Endpoints []string `yaml:"endpoints"`
	}

	Server struct {
		Ip       string `yaml:"ip"`
		Port     int    `yaml:"port"`
		GrpcPort int64  `yaml:"grpcport"`
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
		Prefix   string `yaml:"prefix"`
		OpenFlag bool   `yaml:"open"`
		Ip       string `yaml:"ip"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Db       int    `yaml:"db"`
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
		Ip       string `yaml:"ip"`
		Port     int    `yaml:"port"`
		GrpcPort int    `yaml:"grpcport"`
	}

	PvpWeb struct {
		Endpoints []string `yaml:"endpoints"`
	}
)

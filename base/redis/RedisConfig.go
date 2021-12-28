package redis

type RedisConfig struct {
	Prefix string
	Host   string
	Port   string
	Pass   string
	Db     int
}

func (this *RedisConfig) SetConfig(prefix, host, port, pass string, db int) {
	this.Prefix = prefix
	this.Host = host
	this.Port = port
	this.Pass = pass
	this.Db = db
}

func (this *RedisConfig) Get() *RedisConfig {
	return this
}

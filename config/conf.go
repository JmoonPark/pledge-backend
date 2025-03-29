package config

var Config *Conf

type Conf struct {
	Mysql        MysqlConfig
	Redis        RedisConfig
	DefaultAdmin DefaultAdminConfig
	Email        EmailConfig
	Jwt          JwtConfig
	Env          EnvConfig
}

type MysqlConfig struct {
	Address      string `toml:"address"`
	Port         string `toml:"port"`
	DbName       string `toml:"db_name"`
	UserName     string `toml:"user_name"`
	Password     string `toml:"password"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxLifeTime  int    `toml:"max_life_time"`
}

type RedisConfig struct {
	Address     string `toml:"address"`
	Port        string `toml:"port"`
	Db          int    `toml:"db"`
	Password    string `toml:"password"`
	MaxIdle     int    `toml:"max_idle"`
	MaxActive   int    `toml:"max_active"`
	IdleTimeout int    `toml:"idle_timeout"`
}

type EnvConfig struct {
	Port               string `toml:"port"`
	Version            string `toml:"version"`
	Protocol           string `toml:"protocol"`
	DomainName         string `toml:"domain_name"`
	TaskDuration       int64  `toml:"task_duration"`
	WssTimeoutDuration int64  `toml:"wss_timeout_duration"`
	TaskExtendDuration int64  `toml:"task_extend_duration"`
}

type JwtConfig struct {
	SecretKey  string `toml:"secret_key"`
	ExpireTime int    `toml:"expire_time"` // duration, s
}

type DefaultAdminConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type EmailConfig struct {
	Username string   `toml:"username"`
	Pwd      string   `toml:"pwd"`
	Host     string   `toml:"host"`
	Port     string   `toml:"port"`
	From     string   `toml:"from"`
	Subject  string   `toml:"subject"`
	To       []string `toml:"to"`
	Cc       []string `toml:"cc"`
}

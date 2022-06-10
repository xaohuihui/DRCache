package config

type ServerConfig struct {
	Name               string               `mapstructure:"name"`
	Port               int                  `mapstructure:"port"`
	MysqlInfo          MysqlConfig          `mapstructure:"mysql"`
	PGInfo             PGConfig             `mapstructure:"pg"`
	RedisInfo          RedisConfig          `mapstructure:"redis"`
	LogsAddress        string               `mapstructure:"logsAddress"`
	PasswordLevel      int                  `mapstructure:"passwordLevel"`
	JWTKey             JWTConfig            `mapstructure:"jwt"`
	ConsistentHashInfo ConsistentHashConfig `mapstructure:"consistentHash"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type PGConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsistentHashConfig struct {
	LN        string `mapstructure:"localNode"`
	VNCount int    `mapstructure:"virtualNodeCount"`
}

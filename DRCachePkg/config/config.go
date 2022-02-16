package config

// author: songyanhui
// datetime: 2022/2/16 19:12:14
// software: GoLand

type RedisConfig struct {
	Host      string `ini:"host"`
	Port      string `ini:"port"`
	DefaultDB string `ini:"default_db"`
	Password  string `ini:"password"`
}

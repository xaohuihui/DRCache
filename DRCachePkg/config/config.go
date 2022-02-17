package config

// author: songyanhui
// datetime: 2022/2/16 19:12:14
// software: GoLand

type RedisConfig struct {
	Enable    bool   `ini:"enable"`
	Host      string `ini:"host"`
	Port      string `ini:"port"`
	DefaultDB int    `ini:"default_db"`
	Password  string `ini:"password"`
	PoolSize  int    `ini:"pool_size"`
	TimeOut   string `ini:"timeout"`
}

type LogrusConfig struct {
	Enabled      bool   `ini:"enabled"`
	Path         string `ini:"path"`
	Level        string `ini:"level"`
	Formatter    string `ini:"formatter"`
	OutputType   string `ini:"output_type"`
	ReportCaller bool   `ini:"report_caller"`
	Suffix       string `ini:"suffix_format"`
}

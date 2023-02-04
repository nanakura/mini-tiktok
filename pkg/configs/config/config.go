package config

type GlobalConfig struct {
	DbConfig struct {
		Host     string
		Port     int
		Username string
		Password string
		Database string
	} `mapstructure:"db"`
	MqConfig struct {
		Host string
		Port int
	} `mapstructure:"mq"`
	RedisConfig struct {
		Host string
		Port int
	} `mapstructure:"redis"`
	FtpConfig struct {
		Host     string
		Port     int
		Username string
		Password string
	} `mapstructure:"ftp"`
	MinioConfig struct {
		BucketName string `mapstructure:"bucketName"`
		Endpoint       string `mapstructure:"endpoint"`
		AccessKeyID    string `mapstructure:"accessKeyID"`
		SecretAccessKey string `mapstructure:"secretAccessKey"`
		UseSSL bool `mapstructure:"useSSL"`
	} `mapstructure:"minio"`
}

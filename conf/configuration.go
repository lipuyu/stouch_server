package conf

type Configuration struct {
	Redis struct{
		Node string `yaml:"node"`
		Db int `yaml:"db"`
		Password string `yaml:"password"`
	}

	Database struct{
		DriverName string `yaml:"driverName"`
		Url string `yaml:"url"`
	}

	Oss struct {
		EndPoint string `yaml:"endPoint"`
		AccessKeyId string `yaml:"accessKeyId"`
		AccessKeySecret string `yaml:"accessKeySecret"`
		BucketName string `yaml:"bucketName"`
	}
}

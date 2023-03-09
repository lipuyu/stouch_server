package core

type Configuration struct {
	Redis struct {
		Node     string `yaml:"node"`
		Db       int    `yaml:"db"`
		Password string `yaml:"password"`
	}

	Database struct {
		DriverName string `yaml:"driverName"`
		Url        string `yaml:"url"`
	}

	Oss struct {
		EndPoint        string `yaml:"endPoint"`
		AccessKeyId     string `yaml:"accessKeyId"`
		AccessKeySecret string `yaml:"accessKeySecret"`
		BucketName      string `yaml:"bucketName"`
	}

	Log struct {
		Level       string `yaml:"level"`
		LogFileName string `yaml:"logFileName"`
		LogFilePath string `yaml:"logFilePath"`
	}

	Application struct {
		Addr string `yaml:"addr"`
		Mode string `yaml:"mode"`
	}

	StaticRoot    string   `yaml:"static_root"`
	PathWhiteList []string `yaml:"path_white_list"`
}

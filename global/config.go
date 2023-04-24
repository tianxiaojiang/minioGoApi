package global

type Config struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Minio Minio `mapstructure:"minio" json:"minio" yaml:"minio"`
}

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"appname" json:"appname" yaml:"appname"`
	AppUrl  string `mapstructure:"appurl" json:"appurl" yaml:"appurl"`
	SignKey string `mapstructure:"signkey" json:"signkey" yaml:"signkey"`
}

type Minio struct {
	AccessKey string `mapstructure:"accesskey" json:"accesskey" yaml:"accesskey"`
	SecretKey string `mapstructure:"secretkey" json:"secretkey" yaml:"secretkey"`
	Endpoint  string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	UseSSL    bool   `mapstructure:"usessl" json:"usessl" yaml:"usessl"`
}

var Conf = new(Config)

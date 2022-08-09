package config

import (
	"time"
)

// Public Public配置
type Public struct {
	Server Server `yaml:"Server"`
	Log    Log    `yaml:"Log"`
	Rule   Rule   `yaml:"Rule"`
	App    App    `yaml:"App"`
	Page   Page   `yaml:"Page"`
	Worker Worker `yaml:"Worker"`
	Auto   Auto   `yaml:"Auto"`
	Limit  Limit  `yaml:"Limit"`
}

// Private Private配置
type Private struct {
	Postgresql Postgresql `yaml:"Postgresql"`
	Redis      Redis      `yaml:"Redis"`
	Email      Email      `yaml:"Email"`
	Token      Token      `yaml:"AccessToken"`
	AliyunOSS  AliyunOSS  `yaml:"AliyunOSS"`
}

type Token struct {
	Key                  string        `yaml:"key"`
	UserTokenDuration    time.Duration `yaml:"UserTokenDuration"`
	AccountTokenDuration time.Duration `yaml:"AccountTokenDuration"`
	AuthorizationKey     string        `yaml:"AuthorizationKey"`
	AuthorizationType    string        `yaml:"AuthorizationType"`
}

type Email struct {
	Password string   `yaml:"Password"`
	IsSSL    bool     `yaml:"IsSSL"`
	From     string   `yaml:"From"`
	To       []string `yaml:"To"`
	Host     string   `yaml:"Host"`
	Port     int      `yaml:"Port"`
	UserName string   `yaml:"UserName"`
}

type AliyunOSS struct {
	BucketUrl       string `yaml:"BucketUrl"`
	BasePath        string `yaml:"BasePath"`
	Endpoint        string `yaml:"Endpoint"`
	AccessKeyId     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	BucketName      string `yaml:"BucketName"`
	GroupAvatarUrl  string `yaml:"GroupAvatarUrl"`
}

type Redis struct {
	Address   string        `yaml:"Address"`
	DB        int           `yaml:"DB"`
	Password  string        `yaml:"Password"`
	PoolSize  int           `yaml:"PoolSize"`
	CacheTime time.Duration `yaml:"CacheTime"`
}

type Postgresql struct {
	DriverName string `yaml:"DriverName"`
	SourceName string `yaml:"SourceName"`
}

type Worker struct {
	TaskChanCapacity   int `yaml:"TaskChanCapacity"`
	WorkerChanCapacity int `yaml:"WorkerChanCapacity"`
	WorkerNum          int `yaml:"WorkerNum"`
}

type Rule struct {
	UsernameLenMax   int           `yaml:"UsernameLenMax"`
	UsernameLenMin   int           `yaml:"UsernameLenMin"`
	PasswordLenMax   int           `yaml:"PasswordLenMax"`
	PasswordLenMin   int           `yaml:"PasswordLenMin"`
	UserMarkDuration time.Duration `yaml:"UserMarkDuration"`
	CodeMarkDuration time.Duration `yaml:"CodeMarkDuration"`
	CodeLength       int           `yaml:"CodeLength"`
	AccountNumMax    int32         `yaml:"AccountNumMax"`
	DefaultAvatarURL string        `yaml:"DefaultAvatarURL"`
	BiggestFileSize  int64         `yaml:"BiggestFileSize"`
}

type Server struct {
	RunMode               string        `yaml:"RunMode"`
	Address               string        `yaml:"Address"`
	ReadTimeout           time.Duration `yaml:"ReadTimeout"`
	WriteTimeout          time.Duration `yaml:"WriteTimeout"`
	DefaultContextTimeout time.Duration `yaml:"DefaultContextTimeout"`
}

type Page struct {
	MaxPageSize     int32  `yaml:"MaxPageSize"`
	PageKey         string `yaml:"PageKey"`
	PageSizeKey     string `yaml:"PageSizeKey"`
	DefaultPageSize int32  `yaml:"DefaultPageSize"`
}

type Log struct {
	Level         string `yaml:"Level"`
	LogSavePath   string `yaml:"LogSavePath"`
	LowLevelFile  string `yaml:"LowLevelFile"`
	LogFileExt    string `yaml:"LogFileExt"`
	HighLevelFile string `yaml:"HighLevelFile"`
	MaxSize       int    `yaml:"MaxSize"`
	MaxAge        int    `yaml:"MaxAge"`
	MaxBackups    int    `yaml:"MaxBackups"`
	Compress      bool   `yaml:"Compress"`
}

type App struct {
	Name      string `yaml:"Name"`
	Version   string `yaml:"Version"`
	StartTime string `yaml:"StartTime"`
	MachineID int64  `yaml:"MachineID"`
}

type Limit struct {
	IPLimit  IPLimit  `yaml:"IPLimit"`
	APILimit APILimit `yaml:"APILimit"`
}

type IPLimit struct {
	Cap     int64 `yaml:"Cap"`
	GenNum  int64 `yaml:"GenNum"`
	GenTime int64 `yaml:"GenTime"`
	Cost    int64 `yaml:"Cost"`
}

type Bucket struct {
	Count    int           `yaml:"Count"`
	Duration time.Duration `yaml:"Duration"`
	Burst    int           `yaml:"Burst"`
}

type APILimit struct {
	Upload []Bucket `yaml:"Upload"`
	Email  []Bucket `yaml:"Email"`
}

type Auto struct {
	Retry                     Retry         `yaml:"Retry"`
	DeleteExpiredFileDuration time.Duration `yaml:"DeleteExpiredFileDuration"`
}

type Retry struct {
	Duration time.Duration `yaml:"Duration"`
	MaxTimes int           `yaml:"MaxTimes"`
}

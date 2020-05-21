package server

import (
	"Common/infrastructure/persistence"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Server struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	AllowOrigin  string
}

var ServerSetting = &Server{}

type Database struct {
	Driver   string
	User     string
	Password string
	Host     string
	Name     string
	Port     string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Port        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Not Found :[ environment file ]")
	}
}

func Setup() {

	ServerSetting.RunMode = os.Getenv("APP_MODE")
	ServerSetting.HttpPort = os.Getenv("APP_PORT")
	ServerSetting.ReadTimeout = 60 * time.Second
	ServerSetting.WriteTimeout = 60 * time.Second
	ServerSetting.AllowOrigin = os.Getenv("APP_ALLOW_ORIGIN")
	DatabaseSetting.Driver = os.Getenv("DB_DRIVER")
	DatabaseSetting.Host = os.Getenv("DB_HOST")
	DatabaseSetting.User = os.Getenv("DB_USER")
	DatabaseSetting.Password = os.Getenv("DB_PASSWORD")
	DatabaseSetting.Name = os.Getenv("DB_NAME")
	DatabaseSetting.Port = os.Getenv("DB_PORT")

	RedisSetting.Host = os.Getenv("REDIS_HOST")
	RedisSetting.Port = os.Getenv("REDIS_PORT")
	RedisSetting.Password = os.Getenv("REDIS_PASSWORD")
	SetupDatabase()

}

func SetupDatabase() {
	services, err := persistence.NewRepositories(
		DatabaseSetting.Driver,
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Port,
		DatabaseSetting.Host,
		DatabaseSetting.Name)
	if err != nil {
		panic(err)
	}
	_ = services.AutoMigrate()
	defer services.Close()
}

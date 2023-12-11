package config

import (
	"os"
	"strconv"
	"time"
)

type App struct {
	Host        string
	Port        int
	Debug       bool
	Env         string
	ReadTimeout time.Duration
	FrontendURL string
}

var app = &App{}

func AppCfg() *App {
	return app
}

func LoadApp() {
	app.Host = os.Getenv("APP_HOST")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	app.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	app.Env = os.Getenv("APP_ENV")
	timeOut, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	app.ReadTimeout = time.Duration(timeOut) * time.Second
	app.FrontendURL = os.Getenv("APP_FRONTEND_URL")
}

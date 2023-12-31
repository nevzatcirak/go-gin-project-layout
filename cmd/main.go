package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nevzatcirak/go-gin-poc/api/route"
	"github.com/nevzatcirak/go-gin-poc/bootstrap"
	"github.com/nevzatcirak/go-gin-poc/middleware"
	"io"
	"os"
)

func setLogOutput() {
	f, _ := os.Create("go-gin-poc.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	//setLogOutput()
	app := bootstrap.App()

	env := app.Env

	db := app.DB
	defer app.CloseDBConnection()

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())
	route.Setup(env, db, server)

	var serverAddress string
	if env.Host != "" {
		serverAddress = env.Host
	}
	if env.Port == "" {
		serverAddress += ":80"
	} else {
		serverAddress += ":" + env.Port
	}
	server.Run(serverAddress)
}

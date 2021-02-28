package main

import (
	"os"
	"words-microservice/config"
	"words-microservice/controller"
	"words-microservice/db"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var Datastore *db.MongoDatastore

func init() {
	log.SetOutput(os.Stdout)
	var cfg config.Config
	readFile(&cfg)
	readEnv(&cfg)
	log.Infof("%+v", cfg)
	Datastore = db.NewDatastore(&cfg)
	controller.Datastore = Datastore
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong!")
	})

	r.POST("/words", controller.SearchForWordsHandler)

	r.Run(":3000")
}

func processError(err error) {
	log.Error(err)
	os.Exit(2)
}

func readFile(cfg *config.Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *config.Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

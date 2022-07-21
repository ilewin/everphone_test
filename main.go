package main

import (
	"github.com/sirupsen/logrus"
	"github.com/transparentideas/everphone_test/api"
	"github.com/transparentideas/everphone_test/db/mongost"
	"github.com/transparentideas/everphone_test/utils"
)

func main() {
	conf := utils.GetConfig()
	logrus.Info("Starting server on port", conf.App_Port)
	store, err := mongost.NewMongoStore()
	if err != nil {
		logrus.Fatal("Cannot initialise database connection:", err)
	}
	server := api.NewServer(store)
	err = server.Start(conf.App_Host + ":" + conf.App_Port)
	if err != nil {
		logrus.Fatal("cannot start server:", err)
	}
}

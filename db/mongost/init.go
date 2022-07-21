package mongost

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/transparentideas/everphone_test/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	DB *mongo.Database
}

// Close use this method to close database connection
func (r *MongoStore) Close() {
	logrus.Warning("Closing all db connections")
}

func NewMongoStore() (*MongoStore, error) {
	conf := utils.GetConfig()

	host := conf.Mongo_Uri
	dbName := conf.Mongo_Db
	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &MongoStore{DB: client.Database(dbName)}, nil
}

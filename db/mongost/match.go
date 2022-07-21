package mongost

import (
	"context"

	"github.com/transparentideas/everphone_test/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *MongoStore) MatchGiftToEmployee(ctx context.Context, arg db.MatchGiftToEmployeeParams) ([]db.Employee, error) {
	emplsgifts := make([]db.Employee, 0)
	matchStage := bson.D{{"$match", bson.D{{"name", arg.Name}}}}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "gifts"}, {"localField", "interests"}, {"foreignField", "categories"}, {"as", "gifts"}}}}
	showLoadedCursor, err := r.DB.Collection("employees").Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		panic(err)
	}

	if err = showLoadedCursor.All(ctx, &emplsgifts); err != nil {
		panic(err)
	}
	return emplsgifts, nil
}

func (r *MongoStore) MatchAll(ctx context.Context, arg db.MatchAllParams) ([]db.Employee, error) {
	emplsgifts := make([]db.Employee, 0)
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "gifts"}, {"localField", "interests"}, {"foreignField", "categories"}, {"as", "gifts"}}}}
	showLoadedCursor, err := r.DB.Collection("employees").Aggregate(ctx, mongo.Pipeline{lookupStage})
	if err != nil {
		panic(err)
	}

	if err = showLoadedCursor.All(ctx, &emplsgifts); err != nil {
		panic(err)
	}
	return emplsgifts, nil
}

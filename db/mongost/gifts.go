package mongost

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/transparentideas/everphone_test/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoStore) ListGifts(ctx context.Context, arg db.ListGiftsParams) ([]db.Gift, error) {
	giftsList := make([]db.Gift, 0)
	cursor, err := r.DB.Collection("gifts").Find(ctx, bson.M{})
	if err != nil {
		logrus.Error("Cannot find gifts:", err)
		return giftsList, err
	}

	for cursor.Next(ctx) {
		var gft db.Gift
		err = cursor.Decode(&gft)
		if err != nil {
			logrus.Print(err)
		}
		giftsList = append(giftsList, gft)
	}

	return giftsList, nil
}

func (r *MongoStore) GetGift(ctx context.Context, arg db.GetGiftParams) (*db.Gift, error) {
	var gft db.Gift
	err := r.DB.Collection("gifts").FindOne(ctx, bson.M{"_id": arg.Id}).Decode(&gft)
	if err != nil {
		logrus.Error("Cannot find gift:", err)
		return nil, err
	}

	return &gft, nil
}

func (r *MongoStore) AddGift(ctx context.Context, arg db.AddGiftParams) (*db.Gift, error) {

	gft := db.Gift{
		Id:         primitive.NewObjectID(),
		Name:       arg.Name,
		Categories: arg.Categories,
	}
	_, err := r.DB.Collection("gifts").InsertOne(ctx, gft)
	if err != nil {
		logrus.Error("Cannot add employee:", err)
		return &gft, err
	}

	return &gft, nil
}

func (r *MongoStore) UpdateGift(ctx context.Context, arg db.UpdateGiftParams) (*db.Gift, error) {
	var gft db.Gift
	err := r.DB.Collection("gifts").FindOneAndUpdate(ctx, bson.M{"_id": arg.Id}, bson.M{"$set": bson.M{"name": arg.Name, "categories": arg.Categories}}).Decode(&gft)
	if err != nil {
		logrus.Error("Cannot update gift:", err)
		return &gft, err
	}

	return &gft, nil
}

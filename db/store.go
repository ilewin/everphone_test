package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddEmployeeParams struct {
	Name      string   `bson:"_id" json:"name"`
	Interests []string `bson:"interests" json:"interests"`
}

type GetEmployeeByNameParams struct {
	Name string `uri:"name" bson:"name" json:"name" validate:"required"`
}

type ListEmployeesParams struct{}

type UpdateEmployeeParams struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Interests []string           `bson:"interests" json:"interests"`
}

type ListGiftsParams struct{}

type GetGiftParams struct {
	Id primitive.ObjectID `uri:"id" bson:"_id" json:"id"`
}

type AddGiftParams struct {
	Name       string   `bson:"name" json:"name"`
	Categories []string `bson:"categories" json:"categories"`
}

type UpdateGiftParams struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Categories []string           `bson:"categories" json:"categories"`
}

type MatchGiftToEmployeeParams struct {
	Name string `uri:"name" bson:"name" json:"name"`
}

type MatchAllParams struct {
}

type Store interface {
	Close()
	AddEmployee(ctx context.Context, arg AddEmployeeParams) (*Employee, error)
	GetEmployeeByName(ctx context.Context, arg GetEmployeeByNameParams) (*Employee, error)
	ListEmployees(ctx context.Context, arg ListEmployeesParams) ([]Employee, error)
	UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (*Employee, error)

	AddGift(ctx context.Context, arg AddGiftParams) (*Gift, error)
	GetGift(ctx context.Context, arg GetGiftParams) (*Gift, error)
	ListGifts(ctx context.Context, arg ListGiftsParams) ([]Gift, error)
	UpdateGift(ctx context.Context, arg UpdateGiftParams) (*Gift, error)

	MatchGiftToEmployee(ctx context.Context, arg MatchGiftToEmployeeParams) ([]Employee, error)
	MatchAll(ctx context.Context, arg MatchAllParams) ([]Employee, error)
}

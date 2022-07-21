package mongost

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/transparentideas/everphone_test/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoStore) ListEmployees(ctx context.Context, arg db.ListEmployeesParams) ([]db.Employee, error) {
	employeesList := make([]db.Employee, 0)
	cursor, err := r.DB.Collection("employees").Find(ctx, bson.M{})
	if err != nil {
		logrus.Error("Cannot find employees:", err)
		return employeesList, err
	}

	for cursor.Next(ctx) {
		var empl db.Employee
		err = cursor.Decode(&empl)
		if err != nil {
			logrus.Print(err)
		}
		employeesList = append(employeesList, empl)
	}

	return employeesList, nil
}

func (r *MongoStore) GetEmployeeByName(ctx context.Context, args db.GetEmployeeByNameParams) (*db.Employee, error) {
	var empl db.Employee
	err := r.DB.Collection("employees").FindOne(ctx, bson.M{"name": args.Name}).Decode(&empl)
	if err != nil {
		logrus.Error("Cannot find employee:", err)
		return nil, err
	}

	return &empl, nil
}

func (r *MongoStore) UpdateEmployee(ctx context.Context, arg db.UpdateEmployeeParams) (*db.Employee, error) {
	var empl db.Employee
	err := r.DB.Collection("employees").FindOneAndUpdate(ctx, bson.M{"_id": arg.Id}, bson.M{"$set": bson.M{"name": arg.Name, "interests": arg.Interests}}).Decode(&empl)
	if err != nil {
		logrus.Error("Cannot update employee:", err)
		return &empl, err
	}

	return &empl, nil
}

func (r *MongoStore) AddEmployee(ctx context.Context, arg db.AddEmployeeParams) (*db.Employee, error) {

	empl := db.Employee{
		Id:        primitive.NewObjectID(),
		Name:      arg.Name,
		Interests: arg.Interests,
	}
	_, err := r.DB.Collection("employees").InsertOne(ctx, empl)
	if err != nil {
		logrus.Error("Cannot add employee:", err)
		return &empl, err
	}

	return &empl, nil
}

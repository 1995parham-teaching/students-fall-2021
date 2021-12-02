package store

import (
	"context"
	"fmt"

	"githuh.com/cng-by-example/students/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const Collection = "students"

type MongoDB struct {
	db *mongo.Database
}

func NewMongoDBStore(db *mongo.Database) MongoDB {
	return MongoDB{
		db: db,
	}
}

func (m MongoDB) Save(s model.Student) error {
	_, err := m.db.Collection(Collection).InsertOne(context.TODO(), s)
	if err != nil {
		return fmt.Errorf("mongodb insert failed %w", err)
	}

	return nil
}

func (m MongoDB) LoadByID(id string) (model.Student, error) {
	return model.Student{}, nil
}

func (m MongoDB) Load() ([]model.Student, error) {
	var students []model.Student

	records, err := m.db.Collection(Collection).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("mongo find failed %w", err)
	}

	for records.Next(context.TODO()) {
		var student model.Student

		if err := records.Decode(&student); err != nil {
			return students, fmt.Errorf("mongo record decoding failed %w", err)
		}

		students = append(students, student)
	}

	return students, nil
}

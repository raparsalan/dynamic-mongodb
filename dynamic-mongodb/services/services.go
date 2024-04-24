package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DataService interface {
	CreateData(data bson.M) error
	GetData(indetify string, name string) (bson.M, error)
	GetAll() ([]bson.M, error)
	UpdateData(indetify string, name string, data bson.M) error
	DeleteData(indetify string, name string) error
}

type MongoDBDataService struct {
	Collection *mongo.Collection
	Context    context.Context
}

func NewMongoDBDataService(uri, dbName, collectionName string) (DataService, error) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database(dbName).Collection(collectionName)

	return &MongoDBDataService{
		Collection: collection,
		Context:    ctx,
	}, nil
}

func (s *MongoDBDataService) CreateData(data bson.M) error {
	_, err := s.Collection.InsertOne(s.Context, data)
	return err
}

func (s *MongoDBDataService) GetData(identify string, name string) (bson.M, error) {
	var data bson.M
	filter := bson.M{identify: name}
	err := s.Collection.FindOne(s.Context, filter).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *MongoDBDataService) GetAll() ([]bson.M, error) {
	var dataList []bson.M
	cursor, err := s.Collection.Find(s.Context, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.Context)

	for cursor.Next(s.Context) {
		var data bson.M
		err := cursor.Decode(&data)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return dataList, nil
}

func (s *MongoDBDataService) UpdateData(identify string, name string, data bson.M) error {
	filter := bson.M{identify: name}
	update := bson.M{"$set": data}
	_, err := s.Collection.UpdateOne(s.Context, filter, update)
	return err
}

func (s *MongoDBDataService) DeleteData(identify string, name string) error {
	filter := bson.M{identify: name}
	_, err := s.Collection.DeleteOne(s.Context, filter)

	return err
}

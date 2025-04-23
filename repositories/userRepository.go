package repositories

import (
	dbconfiguration "GO_PROJECT/dbConfiguration"
	"GO_PROJECT/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = dbconfiguration.GetCollection(dbconfiguration.DB, "User")

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user model.User) (*mongo.InsertOneResult, error) {
	// Generate new ID if not set
	if user.ID == primitive.NilObjectID {
		user.ID = primitive.NewObjectID()
	}
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRepository) GetUser(id primitive.ObjectID) (model.User, error) {
	var user model.User
	err := userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (r *UserRepository) GetAllUser() ([]model.User, error) {
	var users []model.User
	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user model.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	defer cursor.Close(context.TODO())
	return users, nil
}

func (r *UserRepository) UpdateUser(id primitive.ObjectID, user model.User) (*mongo.UpdateResult, error) {
	result, err := userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": id})
	return result, err
}

func (r *UserRepository) DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return result, err
}

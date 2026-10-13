package repository

import (
	"context"
	"os"

	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/model"
	"github.com/rodrigobastos42/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (ur *userRepository) UpdateUser(id string, user model.UserDomainInterface) *internalerrors.RestErr {
	collectionName := os.Getenv("MONGODB_DATABASE")

	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(user)
	update := bson.M{
		"$set": value,
	}
	userID, _ := primitive.ObjectIDFromHex(id)

	_, err := collection.UpdateByID(context.Background(), userID, update)
	if err != nil {
		return internalerrors.NewInternalServerError(err.Error())
	}
	return nil
}

package repository

import (
	"context"
	"os"

	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/model"
	"github.com/rodrigobastos42/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *internalerrors.RestErr) {
	collectionName := os.Getenv("MONGODB_DATABASE")

	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(user)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, internalerrors.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return *converter.ConvertEntityToDomain(*value), nil
}

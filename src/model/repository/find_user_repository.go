package repository

import (
	"context"
	"fmt"
	"os"

	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/model"
	"github.com/rodrigobastos42/crud-go/src/model/repository/entity"
	"github.com/rodrigobastos42/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (ur *userRepository) FindUserByID(id string) (*model.UserDomainInterface, *internalerrors.RestErr) {
	collectionName := os.Getenv("MONGODB_DATABASE")

	collection := ur.databaseConnection.Collection(collectionName)

	user := &entity.UserEntity{}
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	err := collection.FindOne(context.Background(), filter).Decode(user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			return nil, internalerrors.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		return nil, internalerrors.NewInternalServerError(errorMessage)
	}

	return converter.ConvertEntityToDomain(*user), nil
}

func (ur *userRepository) FindUserByEmail(email string) (*model.UserDomainInterface, *internalerrors.RestErr) {
	collectionName := os.Getenv("MONGODB_DATABASE")

	collection := ur.databaseConnection.Collection(collectionName)

	user := &entity.UserEntity{}
	filter := bson.M{"email": email}

	err := collection.FindOne(context.Background(), filter).Decode(user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			return nil, internalerrors.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		return nil, internalerrors.NewInternalServerError(errorMessage)
	}

	return converter.ConvertEntityToDomain(*user), nil
}

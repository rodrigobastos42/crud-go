package repository

import (
	"context"
	"os"

	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (ur *userRepository) DeleteUser(id string) *internalerrors.RestErr {
	collectionName := os.Getenv("MONGODB_DATABASE")

	collection := ur.databaseConnection.Collection(collectionName)

	userID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": userID}

	collection.FindOneAndDelete(context.Background(), filter)
	return nil
}

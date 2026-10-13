package repository

import (
	internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"
	"github.com/rodrigobastos42/crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *internalerrors.RestErr)
	FindUserByID(id string) (*model.UserDomainInterface, *internalerrors.RestErr)
	FindUserByEmail(email string) (*model.UserDomainInterface, *internalerrors.RestErr)
	UpdateUser(id string, user model.UserDomainInterface) *internalerrors.RestErr
	DeleteUser(id string) *internalerrors.RestErr
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func NewUserRepository(databaseConnection *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: databaseConnection,
	}
}

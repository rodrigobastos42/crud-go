package model

import internalerrors "github.com/rodrigobastos42/crud-go/src/configuration/internal_errors"

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(id string)

	EncryptPassword()
	GenerateToken() (string, *internalerrors.RestErr)
}

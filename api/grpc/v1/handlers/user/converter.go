package user

import (
	"github.com/aradwann/eenergy/entities"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user *entities.User) *User {
	return &User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

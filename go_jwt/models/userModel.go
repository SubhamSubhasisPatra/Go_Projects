package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"id"`
	FirstName    *string            `json:"first_name" validate:"required, min=2 , max=100"`
	LastName     *string            `json:"Last_name" validate:"required, min=2 , max=100"`
	Password     *string            `json:"password" validate:"required, min=6 , max=50"`
	Email        *string            `json:"email" validate:"email, required"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    *time.Time         `json:"created_at"`
	UpdatedAt    *time.Time         `json:"updated_at"`
	UserId       *string            `json:"user_id"`
}

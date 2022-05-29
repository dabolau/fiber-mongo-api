package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// 用户数据模型
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" validate:"required"`
	Password string             `json:"password" validate:"required"`
	Email    string             `json:"email" validate:"required"`
	Phone    string             `json:"phone" validate:"required"`
}

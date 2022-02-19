package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserId     string             `json:"userId,omitempty" validate:"required"`
	Content    string             `json:"content,omitempty" validate:"required"`
	CreateDate primitive.DateTime `json:"createDate,omitempty"`
}

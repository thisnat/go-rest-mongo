package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	Id         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserId     string             `bson:"userId" json:"userId,omitempty" validate:"required"`
	Content    string             `bson:"content" json:"content,omitempty" validate:"required"`
	CreateDate primitive.DateTime `bson:"createDate" json:"createDate,omitempty"`
}

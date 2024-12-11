package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Title  string              `bson:"title" json:"title"`
	Author string              `bson:"author" json:"author"`
	UserID *primitive.ObjectID `bson:"userid,omitempty" json:"userid"`
}

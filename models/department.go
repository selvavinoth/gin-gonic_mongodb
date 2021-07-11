package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Department struct {
	Id          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	IsActive    bool          `json:"isActive" bson:"isActive"`
	CreatedOn   time.Time     `json:"createdOn" bson:"createdOn"`
	ModifiedOn  time.Time     `json:"modifiedOn" bson:"modifiedOn"`
}

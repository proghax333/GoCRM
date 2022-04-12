package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
	IsActive         bool   `json:"is_active" bson:"is_active"`
}

var Users *mgm.Collection

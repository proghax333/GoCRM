package models

import (
	"github.com/kamva/mgm/v3"
)

type Customer struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"username" bson:"username"`
	SSN              string `json:"password" bson:"password"`
	IsActive         bool   `json:"is_active" bson:"is_active"`
}

var Customers *mgm.Collection

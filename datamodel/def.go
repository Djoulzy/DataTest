package datamodel

import (
	"fmt"
	"time"
)

type BaseItem struct {
	ID        int       `bson:"_id" json:"_id,omitempty"`
	Crea_date time.Time `bson:"crea_date" json:"crea_date"`
	Views     int       `bson:"views" json:"views"`
}

type RichItem struct {
	BaseItem
	Height int
	Width  int
	Tags   string
}

type AnotherItem struct {
	Name    string
	Surname string
	Age     int
}

func (AI AnotherItem) String() string {
	return fmt.Sprintf("My name is %s %s and I'm %d years old", AI.Name, AI.Surname, AI.Age)
}

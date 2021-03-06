package model

import (
	

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Disability struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (disability *Disability) Create() *Disability {
	if err := db.C("disabilities").Insert(disability); err != nil {
		panic(err)
	}
	return disability
}

func Disabilities() []Disability {
	var disabilities, disabilitiesArranged []Disability

	if err := db.C("disabilities").Find(nil).All(&disabilities); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	var disabilityOther Disability

	for _, disability := range disabilities {
		if disability.Id.Hex() == "594cb622472e11263c329906" {
			disabilityOther = disability
			continue
		}
		disabilitiesArranged = append(disabilitiesArranged, disability)
	}
	disabilitiesArranged = append(disabilitiesArranged, disabilityOther)
	return disabilitiesArranged
}

func DisabilityById(id bson.ObjectId) *Disability {
	disability := new(Disability)

	if err := db.C("disabilities").FindId(id).One(&disability); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return disability
}

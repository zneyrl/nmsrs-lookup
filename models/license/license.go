package license

import (
	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type License struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]License, error) {
	licns := []License{}

	if err := db.Licenses.Find(bson.M{}).All(&licns); err != nil {
		return nil, err
	}
	return licns, nil
}

func (licn *License) Insert() (string, error) {
	licn.ID = bson.NewObjectId()

	if err := db.Licenses.Insert(licn); err != nil {
		return "", err
	}
	return licn.ID.Hex(), nil
}

func Find(id string) (*License, error) {
	var licn License

	if !bson.IsObjectIdHex(id) {
		return &licn, models.ErrInvalidObjectID
	}

	if err := db.Licenses.FindId(bson.ObjectIdHex(id)).One(&licn); err != nil {
		return &licn, err
	}
	return &licn, nil
}

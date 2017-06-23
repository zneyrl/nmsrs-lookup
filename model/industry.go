package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func industrySeeder() {
	data := []string{
		"ACTIVITIES OF PRIVATE HOUSEHOLDS AS EMPLOYERS AND UNDIFFENTIATED PRODUCTION ACTIVITIES OF PRIVATE",
		"AGRICULTURE",
		"CONSTRUCTION",
		"EDUCATION",
		"ELECTRICITY, GAS AND WATER SUPPLY",
		"EXTRA - TERRITORIAL ORGANIZATIONS AND BODIES",
		"FINANCIAL INTERMEDIATION",
		"FISHING",
		"HEALTH AND SOCIAL WORK",
		"HOTELS AND RESTAURANTS",
		"MANUFACTURING",
		"MINING AND QUARRYING",
		"OTHER COMMUNITY, SOCIAL AND PERSONAL SERVICE ACTIVITIES",
		"PUBLIC ADMINISTRATION AND DEFENSE",
		"REAL ESTATE, RENTING AND BUSINESS ACTIVITIES",
		"TRANSPORT, STORAGE AND COMMUNICATION",
		"WHOLESALE AND RETAIL TRADE",
	}

	for _, name := range data {
		industry := Industry{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		industry.Create()
	}
}

type Industry struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func (industry *Industry) Create() *Industry {
	if err := db.C("industries").Insert(industry); err != nil {
		panic(err)
	}
	defer db.Close()
	return industry
}

func (industry Industry) Index(q string) []Industry {
	industries := []Industry{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("industries").Find(query).All(&industries); err != mgo.ErrNotFound {
		panic(err)
	} else if err == mgo.ErrNotFound {
		return nil
	}
	defer db.Close()
	return industries
}

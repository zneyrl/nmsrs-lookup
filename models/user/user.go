package user

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/zneyrl/nmsrs-lookup/db"
	"github.com/zneyrl/nmsrs-lookup/env"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

var (
	ErrInvalidObjectID    = errors.New("Invalid object ID")
	ErrActionNotPermitted = errors.New("Action not permitted")
	ErrEmailTaken         = errors.New("Email has already been taken")
	contentDir            = "content/users"
)

type User struct {
	ID              bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name            string        `schema:"name" json:"name" bson:"name,omitempty" validate:"required,min=2"`
	Email           string        `schema:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
	Password        string        `schema:"password" json:"password" bson:"password,omitempty" validate:"required,min=6"`
	ConfirmPassword string        `schema:"confirm_password" json:"confirm_password" bson:",omitempty" validate:"required,eqfield=Password"`
	IsAdmin         bool          `schema:"is_admin" json:"is_admin" bson:"isAdmin"`
	PhotoIsSet      bool          `schema:"photo_is_set" json:"photo_is_set" bson:"photoIsSet"`
	CreatedAt       int64         `schema:"created_at" json:"created_at" bson:"createdAt,omitempty"`
	UpdatedAt       int64         `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
}

func All() ([]User, error) {
	users := []User{}
	if err := db.Users.Find(bson.M{"email": bson.M{"$ne": env.AdminEmail}}).All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (usr *User) Insert() (string, error) {
	usr.ID = bson.NewObjectId()
	usr.ConfirmPassword = ""
	usr.Password = str.Bcrypt(usr.Password)
	now := time.Now().Unix()
	usr.CreatedAt = now
	usr.UpdatedAt = now

	if err := db.Users.Insert(usr); err != nil {
		return "", err
	}
	id := usr.ID.Hex()
	return id, MakeReadMeFile(id)
}

func Find(id string) (User, error) {
	var usr User

	if !bson.IsObjectIdHex(id) {
		return usr, ErrInvalidObjectID
	}

	if err := db.Users.FindId(bson.ObjectIdHex(id)).One(&usr); err != nil {
		return usr, err
	}
	return usr, nil
}

func (usr *User) Delete() error {
	id := usr.ID.Hex()

	if IsAdminUser(id) {
		return ErrActionNotPermitted
	}

	if err := db.Users.RemoveId(usr.ID); err != nil {
		return err
	}
	dir := filepath.Join(contentDir, id)
	_, err := os.Stat(dir)

	if err == nil {
		if err := os.RemoveAll(dir); err != nil {
			return err
		}
	}
	return nil
}

func DeleteMany(ids []string) error {
	for _, id := range ids {
		usr, err := Find(id)

		if err != nil {
			return err
		}

		if err := usr.Delete(); err != nil {
			return err
		}
	}
	return nil
}

func MakeReadMeFile(id string) error {
	file := filepath.Join(contentDir, id, "README.md")

	dir := filepath.Dir(file)
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
	content := fmt.Sprintf("id: %s\n", id)

	if err := ioutil.WriteFile(file, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

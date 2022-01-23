package database

import (
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

func Connect(cName string) (*mgo.Session, *mgo.Collection) {
	url := viper.GetString("mongodb.url")
	session, err := mgo.Dial(url)
	if err != nil {
	}
	session.SetMode(mgo.Monotonic, true)
	return session, session.DB("StudentManager").C(cName)
}

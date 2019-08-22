package mongops

import (
	"gopkg.in/mgo.v2"
)

//CRUD interface holds behavior of all operations made to database and is to be implemented.


// CreateConnection():Function creates a dialup connection to given host and returns the collection object.
func CreateConnection(host string, database string, collection string) (*mgo.Collection, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	return session.DB(database).C(collection), err
}

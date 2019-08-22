package mongops

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Interface to be mocked.
type CRUD interface {
	InsertRecord(mgo.Collection) error
	  
}

type Connection interface{
	Insert(...interface{})(error)
}

// Document holds Json objects reieved and to be inserted to collection.
type Document struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Rating int           `bson:"rating" json:"rating"`
	Cast   []string      `bson:"cast" json:"cast"`
}

// InsertValue adds new document to the given Collection.
func (d *Document) InsertRecord(con mgo.Collection) error {
	err := con.Insert(d)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// DisplayAll Returns all the documents in the collection
func (d *Document) GetAll(con mgo.Collection) (ResultSet []Document, err error) {
	err = con.Find(bson.M{}).All(&ResultSet)
	if err != nil {
		return nil, err
	}
	return ResultSet, nil
}

// DeleteRecord deletes the calling document object from the Collection.
func (d *Document) DeleteRecord(con mgo.Collection) error {
	err := con.Remove(d)
	if err != nil {
		return err
	}
	return nil
}

// FindRecord
func (d Document) FindRecord(query map[string]interface{}, con mgo.Collection) (result Document, err error) {
	err = con.Find(query).One(&result)
	if err != nil {
		fmt.Println(err)
		return Document{}, err
	}
	return result, nil
}

//UpdateRecord updates selected field argument in document.
func (d Document) UpdateRecord(modifyField bson.M, con mgo.Collection) error {
	err := con.Update(bson.M{"_id": d.ID}, bson.M{"$set": modifyField})
	if err != nil {
		return err
	}
	return nil
}

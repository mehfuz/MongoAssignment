package mongops

import (
	"testing"
    // "github.com/golang/mock/gomock"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
)

var session, _ = mgo.Dial("mongodb://localhost:27017")
var Collection = session.DB("POSTS").C("staging")


func TestInsert(test *testing.T) {
	// var err error
	test_doc := Document{
		Name: "dummy", Rating: 3, Cast: []string{"hi", "low"},
	}

	err2 := test_doc.InsertRecord(*Collection)
	if err2 != nil {
		test.Error("Case failed..")
	}
	// mgctrl := gomock.NewController(test)
	// defer mgctrl.Finish()
	// mockRuner := NewMockCRUD(mgctrl)

	// mockRuner.EXPECT().InsertRecord(*Collection).Return(err)





}

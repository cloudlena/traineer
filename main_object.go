package traineer

import "gopkg.in/mgo.v2/bson"

// mainObject is a generic object with some metadata attributes.
type mainObject struct {
	ID          bson.ObjectId
	Name        string
	Description string
	Tags        []string
}

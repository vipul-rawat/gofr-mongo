package mongo

import "go.mongodb.org/mongo-driver/mongo"

type insertManyRes struct {
	*mongo.InsertManyResult
}

func (i *insertManyRes) InsertedIDs() []interface{} {
	return i.InsertManyResult.InsertedIDs
}

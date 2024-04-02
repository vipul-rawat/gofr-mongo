package mongo

import "go.mongodb.org/mongo-driver/mongo"

type insertOneRes struct {
	*mongo.InsertOneResult
}

func (i *insertOneRes) InsertedID() interface{} {
	return i.InsertOneResult.InsertedID
}

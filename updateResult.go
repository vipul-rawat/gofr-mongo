package mongo

import "go.mongodb.org/mongo-driver/mongo"

type updateRes struct {
	*mongo.UpdateResult
}

func (u *updateRes) MatchedCount() int64 {
	return u.UpdateResult.MatchedCount
}

func (u *updateRes) ModifiedCount() int64 {
	return u.UpdateResult.ModifiedCount
}

func (u *updateRes) UpsertedCount() int64 {
	return u.UpdateResult.ModifiedCount
}

func (u *updateRes) UpsertedID() interface{} {
	return u.UpdateResult.UpsertedID
}

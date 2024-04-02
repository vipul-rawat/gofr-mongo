package mongo

type BulkWriteResult interface {
	InsertedCount() int64
	MatchedCount() int64
	ModifiedCount() int64
	DeletedCount() int64
	UpsertedCount() int64
	UpsertedIDs() map[int64]interface{}
}

type InsertOneResult interface {
	InsertedID() interface{}
}

type InsertManyResult interface {
	InsertedIDs() []interface{}
}

type UpdateResult interface {
	MatchedCount() int64
	ModifiedCount() int64
	UpsertedCount() int64
	UpsertedID() interface{}
}

type WriteModel interface {
	writeModel()
}

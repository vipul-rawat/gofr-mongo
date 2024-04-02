package mongo

import "go.mongodb.org/mongo-driver/mongo"

type bulkWriteRes struct {
	*mongo.BulkWriteResult
}

func (b *bulkWriteRes) InsertedCount() int64 {
	return b.BulkWriteResult.InsertedCount
}

func (b *bulkWriteRes) MatchedCount() int64 {
	return b.BulkWriteResult.MatchedCount
}

func (b *bulkWriteRes) ModifiedCount() int64 {
	return b.BulkWriteResult.ModifiedCount
}

func (b *bulkWriteRes) DeletedCount() int64 {
	return b.BulkWriteResult.DeletedCount
}

func (b *bulkWriteRes) UpsertedCount() int64 {
	return b.BulkWriteResult.UpsertedCount
}

func (b *bulkWriteRes) UpsertedIDs() map[int64]interface{} {
	return b.BulkWriteResult.UpsertedIDs
}

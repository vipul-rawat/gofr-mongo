package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gofr.dev/pkg/gofr/datasource"
)

type Client struct {
	*mongo.Database

	logger  Logger
	metrics Metrics
}

var DB = func(conf datasource.Config, logger datasource.Logger, metrics datasource.Metrics) datasource.Mongo {
	logger.Logf("using gofr-mongo as external DB for mongo")

	m, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.Get("MONGO_URI")))
	if err != nil {
		logger.Errorf("error connecting to mongoDB, err:%v", err)

		return nil
	}

	return &Client{
		Database: m.Database(conf.Get("MONGO_DATABASE")),
		logger:   logger,
		metrics:  metrics,
	}
}

func (c *Client) Clone(col string, opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	return c.Database.Collection(col).Clone(opts...)
}

func (c *Client) BulkWrite(ctx context.Context, collection string, models []datasource.WriteModel) (datasource.BulkWriteResult, error) {
	mongoM := make([]mongo.WriteModel, 0)
	for _, v := range models {
		mongoM = append(mongoM, v.(mongo.WriteModel))
	}

	res, err := c.Database.Collection(collection).BulkWrite(ctx, mongoM)

	return datasource.BulkWriteResult{
		InsertedCount: res.InsertedCount,
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		DeletedCount:  res.DeletedCount,
		UpsertedCount: res.UpsertedCount,
		UpsertedIDs:   res.UpsertedIDs,
	}, err
}

func (c *Client) InsertOne(ctx context.Context, collection string, document interface{}) (datasource.InsertOneResult, error) {
	res, err := c.Database.Collection(collection).InsertOne(ctx, document)

	return datasource.InsertOneResult{InsertedID: res.InsertedID}, err
}

func (c *Client) InsertMany(ctx context.Context, col string, documents []interface{}) (datasource.InsertManyResult, error) {
	res, err := c.Database.Collection(col).InsertMany(ctx, documents)

	return datasource.InsertManyResult{InsertedIDs: res.InsertedIDs}, err
}

func (c *Client) DeleteOne(ctx context.Context, col string, filter interface{}) (int64, error) {
	res, err := c.Database.Collection(col).DeleteOne(ctx, filter)

	return res.DeletedCount, err
}

func (c *Client) DeleteMany(ctx context.Context, col string, filter interface{}) (int64, error) {
	res, err := c.Database.Collection(col).DeleteMany(ctx, filter)

	return res.DeletedCount, err
}

func (c *Client) UpdateByID(ctx context.Context, col string, id interface{}, update interface{}) (datasource.UpdateResult, error) {
	res, err := c.Database.Collection(col).UpdateByID(ctx, id, update)

	return datasource.UpdateResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
		UpsertedID:    res.UpsertedID,
	}, err
}

func (c *Client) UpdateOne(ctx context.Context, col string, filter interface{}, update interface{}) (datasource.UpdateResult, error) {
	res, err := c.Database.Collection(col).UpdateOne(ctx, filter, update)

	return datasource.UpdateResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
		UpsertedID:    res.UpsertedID,
	}, err
}

func (c *Client) UpdateMany(ctx context.Context, collection string, filter interface{}, update interface{}) (datasource.UpdateResult, error) {
	res, err := c.Database.Collection(collection).UpdateMany(ctx, filter, update)

	return datasource.UpdateResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
		UpsertedID:    res.UpsertedID,
	}, err
}

func (c *Client) ReplaceOne(ctx context.Context, col string, filter interface{}, replacement interface{}) (datasource.UpdateResult, error) {
	res, err := c.Database.Collection(col).ReplaceOne(ctx, filter, replacement)

	return datasource.UpdateResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
		UpsertedID:    res.UpsertedID,
	}, err
}

func (c *Client) Aggregate(ctx context.Context, col string, pipeline interface{}) (datasource.Cursor, error) {
	return c.Database.Collection(col).Aggregate(ctx, pipeline)
}

func (c *Client) CountDocuments(ctx context.Context, col string, filter interface{}) (int64, error) {
	return c.Database.Collection(col).CountDocuments(ctx, filter)
}

func (c *Client) EstimatedDocumentCount(ctx context.Context, col string) (int64, error) {
	return c.Database.Collection(col).EstimatedDocumentCount(ctx)
}

func (c *Client) Distinct(ctx context.Context, col string, fieldName string, filter interface{}) ([]interface{}, error) {
	return c.Database.Collection(col).Distinct(ctx, fieldName, filter)
}

func (c *Client) Find(ctx context.Context, col string, filter interface{}) (datasource.Cursor, error) {
	cu, err := c.Database.Collection(col).Find(ctx, filter)

	return cu, err
}

func (c *Client) FindOne(ctx context.Context, col string, filter interface{}) datasource.SingleResult {
	return c.Database.Collection(col).FindOne(ctx, filter)
}

func (c *Client) FindOneAndDelete(ctx context.Context, col string, filter interface{}) datasource.SingleResult {
	return c.Database.Collection(col).FindOneAndDelete(ctx, filter)
}

func (c *Client) FindOneAndReplace(ctx context.Context, col string, filter interface{}, replacement interface{}) datasource.SingleResult {
	return c.Database.Collection(col).FindOneAndReplace(ctx, filter, replacement)
}

func (c *Client) FindOneAndUpdate(ctx context.Context, col string, filter interface{}, update interface{}) datasource.SingleResult {
	return c.Database.Collection(col).FindOneAndUpdate(ctx, filter, update)
}

func (c *Client) Watch(ctx context.Context, col string, pipeline interface{}) (*mongo.ChangeStream, error) {
	return c.Database.Collection(col).Watch(ctx, pipeline)
}

func (c *Client) Indexes(col string) mongo.IndexView {
	return c.Database.Collection(col).Indexes()
}

func (c *Client) SearchIndexes(col string) mongo.SearchIndexView {
	return c.Database.Collection(col).SearchIndexes()
}

func (c *Client) Drop(ctx context.Context, col string) error {
	return c.Database.Collection(col).Drop(ctx)
}

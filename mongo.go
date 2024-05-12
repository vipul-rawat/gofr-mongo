package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	*mongo.Database

	logger  Logger
	metrics Metrics
}

func New(conf Config, logger Logger, metrics Metrics) *Client {
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

func (c *Client) InsertOne(ctx context.Context, collection string, document interface{}) (interface{}, error) {
	c.logger.Debug("InsertOne")

	return c.Database.Collection(collection).InsertOne(ctx, document)
}

func (c *Client) Find(ctx context.Context, collection string, filter interface{}, results interface{}) error {
	c.logger.Debug("Find")

	cur, err := c.Database.Collection(collection).Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, results); err != nil {
		return err
	}

	return nil
}

func (c *Client) FindOne(ctx context.Context, collection string, filter interface{}, result interface{}) error {
	c.logger.Debug("FindOne")

	res := c.Database.Collection(collection).FindOne(ctx, filter)
	if res.Err() != nil {
		return res.Err()
	}

	b, err := res.Raw()
	if err != nil {
		return err
	}

	return bson.Unmarshal(b, result)
}

func (c *Client) InsertMany(ctx context.Context, collection string, documents []interface{}) ([]interface{}, error) {
	c.logger.Debug("InsertMany")

	res, err := c.Database.Collection(collection).InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (c *Client) DeleteOne(ctx context.Context, collection string, filter interface{}) (int64, error) {
	c.logger.Debug("DeleteOne")

	res, err := c.Database.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func (c *Client) DeleteMany(ctx context.Context, collection string, filter interface{}) (int64, error) {
	res, err := c.Database.Collection(collection).DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func (c *Client) UpdateByID(ctx context.Context, collection string, id interface{}, update interface{}) (int64, error) {
	c.logger.Debug("UpdateByID")

	res, err := c.Database.Collection(collection).UpdateByID(ctx, id, update)

	return res.ModifiedCount, err
}

func (c *Client) UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) error {
	c.logger.Debug("UpdateOne")

	_, err := c.Database.Collection(collection).UpdateOne(ctx, filter, update)

	return err
}

func (c *Client) UpdateMany(ctx context.Context, collection string, filter interface{}, update interface{}) (int64, error) {
	c.logger.Debug("updateMany")

	res, err := c.Database.Collection(collection).UpdateMany(ctx, filter, update)

	return res.ModifiedCount, err
}

func (c *Client) CountDocuments(ctx context.Context, collection string, filter interface{}) (int64, error) {
	c.logger.Debug("CountDocuments")

	return c.Database.Collection(collection).CountDocuments(ctx, filter)
}

func (c *Client) Drop(ctx context.Context, collection string) error {
	return c.Database.Collection(collection).Drop(ctx)
}

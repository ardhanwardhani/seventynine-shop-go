package mongo

import (
	"context"
	"seventynine-shop-go/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(db *mongo.Database) domain.ProductRepository {
	return &MongoProductRepository{collection: db.Collection("products")}
}

func (r *MongoProductRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *MongoProductRepository) GetByID(id int) (*domain.Product, error) {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	var product domain.Product
	err := r.collection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *MongoProductRepository) Create(product *domain.Product) error {
	_, err := r.collection.InsertOne(context.Background(), product)
	return err
}

func (r *MongoProductRepository) Update(product *domain.Product) error {
	filter := bson.D{primitive.E{Key: "id", Value: product.ID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: product.Name},
			{Key: "stock", Value: product.Stock},
		}},
	}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *MongoProductRepository) Delete(id int) error {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

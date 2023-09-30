package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"golang-mongodb-repository/models"
)

type ProductRepositoryDB struct {
	ProductCollection *mongo.Collection
}

type ProductRepository interface {
	Insert(product models.Product) (bool, error)
	GetAll() ([]models.Product, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (repo *ProductRepositoryDB) Insert(product models.Product) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product.ID = primitive.NewObjectID()
	result, err := repo.ProductCollection.InsertOne(ctx, product)

	if err != nil {
		return false, err
	}

	if result.InsertedID == nil {
		return false, errors.New("failed to add product")
	}

	return true, nil
}

func (repo *ProductRepositoryDB) GetAll() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := repo.ProductCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error getting products:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			log.Println("Error decoding product:", err)
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error iterating over products:", err)
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := repo.ProductCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return false, err
	}

	if result.DeletedCount <= 0 {
		return false, errors.New("no product found to delete")
	}

	return true, nil
}

func NewProductRepositoryDB(dbClient *mongo.Collection) ProductRepositoryDB {
	return ProductRepositoryDB{ProductCollection: dbClient}
}

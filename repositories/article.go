package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"kms/db"
	"kms/models"
)

type MongoArticleRepository struct {
	collection *mongo.Collection
}

func NewMongoArticleRepositories() *MongoArticleRepository {
	return &MongoArticleRepository{db.MongoDb.Collection("tb_article")}
}

func (r *MongoArticleRepository) CreateArticle(article *models.Article) error {
	result, err := r.collection.InsertOne(context.Background(), article)
	if err == nil {
		fmt.Println("created article", result.InsertedID)
	}
	return err
}

func (r *MongoArticleRepository) UpdateArticle(article *models.Article) error {
	result, err := r.collection.UpdateOne(context.Background(), bson.M{"id": article.ID}, bson.M{"$set": article})
	if err == nil {
		fmt.Println("update article", result.UpsertedID)
	}
	return err
}

func (r *MongoArticleRepository) DeleteArticle(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err == nil {
		fmt.Println("update article", id)
	}
	return err
}

func (r *MongoArticleRepository) FindArticleById(id string) (*models.Article, error) {
	var article models.Article
	err := r.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&article)
	return &article, err
}

func (r *MongoArticleRepository) FindAllArticle() ([]*models.Article, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var articles []*models.Article
	for cursor.Next(context.Background()) {
		var article models.Article
		if err := cursor.Decode(&article); err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	return articles, nil
}

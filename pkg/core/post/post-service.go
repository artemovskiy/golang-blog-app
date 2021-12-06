package post

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PostService struct {
	mongo *mongo.Client
}

func NewPostService() (*PostService, error) {
	service := PostService{}
	err := service.init()
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *PostService) init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017/blog").SetAuth(options.Credential{
		Username: "root",
		Password: "root",
	}))
	if err != nil {
		return err
	}
	r.mongo = client
	return nil
}

func (r *PostService) Get(id string) (*Post, error) {
	var result Post
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = r.getPostsCollection().FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	result.Id = id
	return &result, nil
}

func (r *PostService) Create(data *CreatePostData) (*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := r.getPostsCollection().InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}

	post := Post{
		Id:   res.InsertedID.(primitive.ObjectID).Hex(),
		Text: data.Text,
	}

	fmt.Println("creating post: " + data.Text)

	return &post, nil
}

func (r *PostService) getPostsCollection() *mongo.Collection {
	return r.mongo.Database("blog").Collection("posts")
}

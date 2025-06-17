package repository

import (
	"context"
	"time"

	"github.com/C0deNe0/blog-go/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository interface {
	Create(ctx context.Context, post domain.Post) error
	List(ctx context.Context) ([]domain.Post, error)
	GetById(ctx context.Context, id string) (*domain.Post, error)
	Update(ctx context.Context, post domain.Post) error
	Delete(ctx context.Context, id string) error
}

type mongoPostRepo struct {
	coll *mongo.Collection
}

// Create implements PostRepository.
func (m *mongoPostRepo) Create(ctx context.Context, post domain.Post) error {
	post.CreatedAt = time.Now()

	_, err := m.coll.InsertOne(ctx, post)
	return err
}

// Delete implements PostRepository.
func (m *mongoPostRepo) Delete(ctx context.Context, id string) error {
	_, err := m.coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// GetById implements PostRepository.
func (m *mongoPostRepo) GetById(ctx context.Context, id string) (*domain.Post, error) {
	var post domain.Post

	err := m.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// List implements PostRepository.
func (m *mongoPostRepo) List(ctx context.Context) ([]domain.Post, error) {
	cursor, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var posts []domain.Post

	for cursor.Next(ctx) {
		var post domain.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Update implements PostRepository.
func (m *mongoPostRepo) Update(ctx context.Context, post domain.Post) error {
	_, err := m.coll.UpdateOne(ctx,
		bson.M{"_id": post.Id},
		bson.M{"$set": bson.M{
			"title":     post.Title,
			"content":   post.Content,
			"author_id": post.AuthorID,
			"tags":      post.Tags,
		}})

	return err
}

func NewPostRepository(db *mongo.Database) PostRepository {
	return &mongoPostRepo{
		coll: db.Collection("posts"),
	}
}

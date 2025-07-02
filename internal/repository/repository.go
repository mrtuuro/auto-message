package repository

import (
	"context"
	"time"

	"github.com/mrtuuro/auto-messager/internal/apperror"
	"github.com/mrtuuro/auto-messager/internal/code"
	"github.com/mrtuuro/auto-messager/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MessageRepository interface {
	NextUnsent(ctx context.Context, limit int) ([]model.Message, error)
	MarkSent(ctx context.Context, id, messageId string, t time.Time) error
	ListSent(ctx context.Context, limit, offset int) ([]model.Message, error)
}

type mongoMessageRepository struct {
	collection *mongo.Collection
}

func NewMongoMessageRepository(coll *mongo.Collection) MessageRepository {
	return &mongoMessageRepository{collection: coll}
}

func (r *mongoMessageRepository) NextUnsent(ctx context.Context, limit int) ([]model.Message, error) {
	cur, err := r.collection.
		Find(ctx,
		bson.M{"sent": false},
		options.Find().SetSort(bson.M{"created_at": 1}).SetLimit(int64(limit)))
	if err != nil {
		return nil, apperror.NewAppError(
			code.ErrSystemDBFailure,
			err,
			code.GetErrorMessage(code.ErrSystemDBFailure),
			)
	}
	var msgs []model.Message
	if err = cur.All(ctx, &msgs); err != nil {
		return nil, apperror.NewAppError(
			code.ErrSystemDBFailure,
			err,
			code.GetErrorMessage(code.ErrSystemDBFailure),
			)
	}
	return msgs, nil
}
func (r *mongoMessageRepository) MarkSent(ctx context.Context, id, messageId string, t time.Time) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.UpdateByID(ctx, oid,
		bson.M{"$set": bson.M{"sent": true, "sent_at": t, "message_id": messageId}})
	if err != nil {
		return apperror.NewAppError(
			code.ErrSystemDBFailure,
			err,
			code.GetErrorMessage(code.ErrSystemDBFailure),
			)

	}
	return nil
}
func (r *mongoMessageRepository) ListSent(ctx context.Context, limit, offset int) ([]model.Message, error) {
	cur, err := r.collection.Find(ctx,
		bson.M{"sent": true},
		options.Find().
			SetSort(bson.M{"sent_at": -1}).
			SetSkip(int64(offset)).
			SetLimit(int64(limit)))
	if err != nil {
		return nil, apperror.NewAppError(
			code.ErrSystemDBFailure,
			err,
			code.GetErrorMessage(code.ErrSystemDBFailure),
			)
	}
	var msgs []model.Message
	if err = cur.All(ctx, &msgs); err != nil {
		return nil, apperror.NewAppError(
			code.ErrSystemDBFailure,
			err,
			code.GetErrorMessage(code.ErrSystemDBFailure),
			)
	}
	return msgs, nil
}

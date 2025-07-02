package service

import (
	"context"
	"time"

	"github.com/mrtuuro/auto-messager/internal/dispatcher"
	"github.com/mrtuuro/auto-messager/internal/model"
	"github.com/mrtuuro/auto-messager/internal/repository"
)

type MessageService interface {
	ListSent(ctx context.Context, limit, offset int) ([]model.Message, error)
	Flush(ctx context.Context, limit int) error
}

type messageService struct {
	repo repository.MessageRepository
	send *dispatcher.Dispatcher
}

func NewMessageService(repo repository.MessageRepository, d *dispatcher.Dispatcher) MessageService {
	return &messageService{
		repo: repo,
		send: d,
	}
}

func (s *messageService) Flush(ctx context.Context, limit int) error {
	msgs, err := s.repo.NextUnsent(ctx, limit)
	if err != nil {
		return err
	}
	for _, m := range msgs {
		if len(m.Content) > 160 {
			continue // or mark failed
		}
		msgID, err := s.send.Send(ctx, m.To, m.Content)
		if err != nil {
			continue
		}
		_ = s.repo.MarkSent(ctx, m.ID.Hex(), msgID, time.Now().UTC())
	}
	return nil

}

func (s *messageService) ListSent(ctx context.Context, limit, offset int) ([]model.Message, error) {
	msgs, err := s.repo.ListSent(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

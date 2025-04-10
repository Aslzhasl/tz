package service

import (
	"context"
	"time"
	"tz/model"
	"tz/repository"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string) (*model.Task, error)
	GetTask(ctx context.Context, id string) (*model.Task, error)
	ListTasks(ctx context.Context) ([]model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, title, description string) (*model.Task, error) {
	task := &model.Task{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := s.repo.Create(ctx, task)
	return task, err
}

func (s *taskService) GetTask(ctx context.Context, id string) (*model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *taskService) ListTasks(ctx context.Context) ([]model.Task, error) {
	return s.repo.List(ctx)
}

func (s *taskService) UpdateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	task.UpdatedAt = time.Now()
	err := s.repo.Update(ctx, task)
	return task, err
}

func (s *taskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

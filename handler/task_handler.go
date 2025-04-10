package handler

import (
	"context"
	pb "tz/api/pb"
	"tz/model"
	"tz/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskHandler struct {
	pb.UnimplementedTaskServiceServer
	Service service.TaskService
}

func (h *TaskHandler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	task, err := h.Service.CreateTask(ctx, req.Title, req.Description)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create task: %v", err)
	}
	return &pb.TaskResponse{Task: toProto(task)}, nil
}

func (h *TaskHandler) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	task, err := h.Service.GetTask(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "task not found: %v", err)
	}
	return &pb.TaskResponse{Task: toProto(task)}, nil
}

func (h *TaskHandler) ListTasks(ctx context.Context, _ *pb.Empty) (*pb.ListTasksResponse, error) {
	tasks, err := h.Service.ListTasks(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not list tasks: %v", err)
	}
	var pbTasks []*pb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, toProto(&t))
	}
	return &pb.ListTasksResponse{Tasks: pbTasks}, nil
}

func (h *TaskHandler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	task := &model.Task{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
	updated, err := h.Service.UpdateTask(ctx, task)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not update task: %v", err)
	}
	return &pb.TaskResponse{Task: toProto(updated)}, nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := h.Service.DeleteTask(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete task: %v", err)
	}
	return &pb.DeleteTaskResponse{Message: "deleted successfully"}, nil
}

func toProto(t *model.Task) *pb.Task {
	return &pb.Task{
		Id:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   t.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

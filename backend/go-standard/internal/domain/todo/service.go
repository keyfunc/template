package todo

import (
	"context"
)

type Service struct {
	Repo *Repository
}
type ServiceDeps struct {
	Repo *Repository
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		Repo: deps.Repo,
	}

}

func (s *Service) List(ctx context.Context, query ListTodoReq) (*ListTodoRes, error) {
	res, err := s.Repo.QueryTodo(ctx, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Create(ctx context.Context, body CreateTodoReq) (Todo, error) {
	res, err := s.Repo.CreateTodo(ctx, body)
	if err != nil {
		return Todo{}, err
	}
	return res, nil
}

func (s *Service) Update(ctx context.Context, id int64, body UpdateTodoReq) (Todo, error) {
	res, err := s.Repo.UpdateTodo(ctx, id, body)
	if err != nil {
		return Todo{}, err
	}
	return res, nil
}

func (s *Service) Del(ctx context.Context, id int64) error {
	err := s.Repo.DelTodo(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

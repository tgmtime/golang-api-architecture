package usecase

import (
	"context"

	"github.com/tgmtime/golang-api-architecture/config"
	"github.com/tgmtime/golang-api-architecture/domain/filter"
	model "github.com/tgmtime/golang-api-architecture/domain/model"
	"github.com/tgmtime/golang-api-architecture/domain/repository"
	"github.com/tgmtime/golang-api-architecture/usecase/dto"
)

type ColorUsecase struct {
	base *BaseUsecase[model.Color, dto.CreateColor, dto.UpdateColor, dto.Color]
}

func NewColorUsecase(cfg *config.Config, repository repository.ColorRepository) *ColorUsecase {
	return &ColorUsecase{
		base: NewBaseUsecase[model.Color, dto.CreateColor, dto.UpdateColor, dto.Color](cfg, repository),
	}
}

// Create
func (u *ColorUsecase) Create(ctx context.Context, req dto.CreateColor) (dto.Color, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *ColorUsecase) Update(ctx context.Context, id int, req dto.UpdateColor) (dto.Color, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *ColorUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *ColorUsecase) GetById(ctx context.Context, id int) (dto.Color, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *ColorUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.Color], error) {
	return s.base.GetByFilter(ctx, req)
}

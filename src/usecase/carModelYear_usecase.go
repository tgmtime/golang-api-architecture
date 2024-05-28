package usecase

import (
	"context"

	"github.com/tgmtime/golang-api-architecture/config"
	"github.com/tgmtime/golang-api-architecture/domain/filter"
	model "github.com/tgmtime/golang-api-architecture/domain/model"
	"github.com/tgmtime/golang-api-architecture/domain/repository"
	"github.com/tgmtime/golang-api-architecture/usecase/dto"
)

type CarModelYearUsecase struct {
	base *BaseUsecase[model.CarModelYear, dto.CreateCarModelYear, dto.UpdateCarModelYear, dto.CarModelYear]
}

func NewCarModelYearUsecase(cfg *config.Config, repository repository.CarModelYearRepository) *CarModelYearUsecase {
	return &CarModelYearUsecase{
		base: NewBaseUsecase[model.CarModelYear, dto.CreateCarModelYear, dto.UpdateCarModelYear, dto.CarModelYear](cfg, repository),
	}
}

// Create
func (u *CarModelYearUsecase) Create(ctx context.Context, req dto.CreateCarModelYear) (dto.CarModelYear, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CarModelYearUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelYear) (dto.CarModelYear, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelYearUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelYearUsecase) GetById(ctx context.Context, id int) (dto.CarModelYear, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelYearUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelYear], error) {
	return s.base.GetByFilter(ctx, req)
}

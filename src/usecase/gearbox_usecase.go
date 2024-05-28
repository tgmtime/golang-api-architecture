package usecase

import (
	"context"

	"github.com/tgmtime/golang-api-architecture/config"
	"github.com/tgmtime/golang-api-architecture/domain/filter"
	model "github.com/tgmtime/golang-api-architecture/domain/model"
	"github.com/tgmtime/golang-api-architecture/domain/repository"
	"github.com/tgmtime/golang-api-architecture/usecase/dto"
)

type GearboxUsecase struct {
	base *BaseUsecase[model.Gearbox, dto.Name, dto.Name, dto.IdName]
}

func NewGearboxUsecase(cfg *config.Config, repository repository.GearboxRepository) *GearboxUsecase {
	return &GearboxUsecase{
		base: NewBaseUsecase[model.Gearbox, dto.Name, dto.Name, dto.IdName](cfg, repository),
	}
}

// Create
func (u *GearboxUsecase) Create(ctx context.Context, req dto.Name) (dto.IdName, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *GearboxUsecase) Update(ctx context.Context, id int, req dto.Name) (dto.IdName, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *GearboxUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *GearboxUsecase) GetById(ctx context.Context, id int) (dto.IdName, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *GearboxUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.IdName], error) {
	return s.base.GetByFilter(ctx, req)
}

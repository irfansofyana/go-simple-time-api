package service

import (
	"timezone-api/domain"
	"timezone-api/error"
)

type TimeService interface {
	GetTimes(tZones string) (*domain.TimesDTO, *error.AppError)
}

type DefaultTimeService struct {
	repo domain.TimeRepository
}

func (d DefaultTimeService) GetTimes(tZones string) (*domain.TimesDTO, *error.AppError) {
	t, err := d.repo.GetAllTimes(tZones)
	if err != nil {
		return nil, err
	}
	return &domain.TimesDTO{Times: t}, nil
}

func NewDefaultTimeService(repo domain.TimeRepository) DefaultTimeService {
	return DefaultTimeService{repo}
}

package domain

import (
	"strings"
	"time"
	"timezone-api/error"
)

type TimeRepository interface {
	GetAllTimes(tZones string) (*TimesDomain, *error.AppError)
}

type DefaultTimeRepository struct{}

func (d DefaultTimeRepository) GetAllTimes(tZones string) (*TimesDomain, *error.AppError) {
	tZonesArr := strings.Split(tZones, ",")
	currTimes := make(map[string]string)

	for _, tZone := range tZonesArr {
		if len(tZone) == 0 {
			continue
		}

		timeLoc, err := time.LoadLocation(tZone)
		if err != nil {
			return nil, error.NewInvalidTimezoneError()
		}
		currTimes[tZone] = time.Now().In(timeLoc).String()
	}

	if len(currTimes) == 0 {
		return &TimesDomain{CurrentTime: time.Now().In(time.UTC).String()}, nil
	}

	return &TimesDomain{OtherTimes: currTimes}, nil
}

func NewDefaultTimeRepository() DefaultTimeRepository {
	var tRepo DefaultTimeRepository
	return tRepo
}

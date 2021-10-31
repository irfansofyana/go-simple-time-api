package app

import (
	"net/http"
	"timezone-api/service"
)

type TimeHandler struct {
	service service.TimeService
}

func (th *TimeHandler) handleTimeAPI(w http.ResponseWriter, r *http.Request) {
	tZones := r.URL.Query().Get("tz")
	timesNow, err := th.service.GetTimes(tZones)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, timesNow.SendAPIResponse())
	}
}

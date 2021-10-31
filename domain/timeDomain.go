package domain

type TimesDomain struct {
	CurrentTime string            `json:"current_time,omitempty"`
	OtherTimes  map[string]string `json:",omitempty"`
}

type TimesDTO struct {
	Times *TimesDomain
}

func (t *TimesDTO) SendAPIResponse() map[string]string {
	if len(t.Times.CurrentTime) > 0 {
		return map[string]string{"current_time": t.Times.CurrentTime}
	}
	return t.Times.OtherTimes
}

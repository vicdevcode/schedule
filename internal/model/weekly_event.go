package model

type CreateWeeklyEvent struct {
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	WeekDay   string `json:"week_day"`
	WeekType  string `json:"week_type"`
}

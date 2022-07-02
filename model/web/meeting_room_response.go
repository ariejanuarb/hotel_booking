package web

type MeetingRoomResponse struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Capacity     int     `json:"capacity"`
	Facility     string  `json:"facility"`
	PricePerHour float64 `json:"price_per_hour"`
	PricePerDay  float64 `json:"price_per_day"`
	DailyRevenue float64 `json:"daily_revenue"`
}

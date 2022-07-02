package web

type MeetingRoomUpdateRequest struct {
	Id           int     `validate:"required"`
	Name         string  `validate:"required,max=100,min=1" json:"name"`
	Capacity     int     `validate:"required" json:"capacity"`
	Facility     string  `validate:"required,max=100,min=1" json:"facility"`
	PricePerHour float64 `validate:"required" json:"price_per_hour"`
	PricePerDay  float64 `validate:"required" json:"price_per_day"`
	DailyRevenue float64 `validate:"required" json:"daily_revenue"`
}

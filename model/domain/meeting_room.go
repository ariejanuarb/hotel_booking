package domain

type MeetingRoom struct {
	Id           int
	Name         string
	Capacity     int
	Facility     string
	PricePerHour float64
	PricePerDay  float64
	DailyRevenue float64
}

package domain

import "time"

type Invoice struct {
	Invoice_id   int
	Invoice_Date string
	Tax          string
	Price        string
	Total        string
	Created_at   time.Time
	Updated_at   time.Time
}

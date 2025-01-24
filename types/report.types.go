package types

import "time"

type SummaryResponse struct {
	TotalTicketsSold int64   `json:"total_tickets_sold"`
	TotalRevenue     float64 `json:"total_revenue"`
}

type CinemaSummaryResponse struct {
	CinemaName     string  `json:"cinema_name"`
	CinemaLocation string  `json:"cinema_location"`
	TicketsSold    int64   `json:"tickets_sold"`
	TotalRevenue   float64 `json:"total_revenue"`
}

type MovieSummaryResponse struct {
	MovieTitle     string  `json:"movie_title"`
	TicketsSold    int64   `json:"tickets_sold"`
	TotalRevenue   float64 `json:"total_revenue"`
	AverageRevenue float64 `json:"average_revenue_per_ticket"`
}

type DailySummaryResponse struct {
	Date         time.Time `json:"date"`
	TicketsSold  int64     `json:"tickets_sold"`
	TotalRevenue float64   `json:"total_revenue"`
}

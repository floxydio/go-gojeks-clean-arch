package dto

import (
	"gojeksrepo/ent/trip"
	"time"
)

type TripRequest struct {
	UserID      string      `json:"user_id" form:"user_id"`
	PickupLat   float64     `json:"pickup_lat,omitempty" form:"pickup_lat"`
	PickupLong  float64     `json:"pickup_long,omitempty" form:"pickup_long"`
	DropLat     float64     `json:"drop_lat,omitempty" form:"drop_lat"`
	DropLong    float64     `json:"drop_long,omitempty" form:"drop_long"`
	Status      trip.Status `json:"status,omitempty" form:"status"`
	DistanceKm  float64     `json:"distance_km,omitempty" form:"distance_km"`
	Numeric     string      `json:"numeric,omitempty" form:"numeric"`
	IsPaid      bool        `json:"is_paid,omitempty" form:"is_paid"`
	CreatedAt   time.Time   `json:"created_at,omitempty" form:"created_at"`
	StartedAt   time.Time   `json:"started_at,omitempty" form:"started_at"`
	CompletedAt time.Time   `json:"completed_at,omitempty" form:"completed_at"`
}

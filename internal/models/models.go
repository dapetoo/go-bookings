package models

import "time"

// Reservation holds reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Users is the user model object
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Room is the room model object
type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restrictions is the restriction model object
type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservations is the reservation model object
type Reservations struct {
	ID        int
	RoomID    int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Rooms
}

// RoomRestrictions is the room restriction model object
type RoomRestrictions struct {
	ID            int
	RoomID        int
	ReservationID int
	RestrictionID int
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Reservation   Reservations
	Room          Rooms
	Restriction   Restrictions
}

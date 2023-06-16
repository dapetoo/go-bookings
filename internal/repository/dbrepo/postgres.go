package dbrepo

import (
	"context"
	"github.com/dapetoo/go-bookings/internal/models"
	"time"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservations InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservations(res models.Reservation) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
    	values($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, query, res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

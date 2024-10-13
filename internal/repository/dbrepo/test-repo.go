package dbrepo

import (
	"time"

	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// INSERTA A RESERVATION INTO DATABASE
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	return 1, nil
}

// nsertRoomRestriction inserta a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

// SearchAvailabilityByDatesByRoomID return true if availability exists for roomID,and false if no availability
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, rooID int) (bool, error) {

	return false, nil
}

// earchAvailabilityForAllRooms returns a slice of all available rooms if any for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}

// GetRoomByID gets a room by its id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {

	var room models.Room

	return room, nil

}
func (m *testDBRepo) GetUserById(id int) (models.User, error) {
	var u models.User
	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {

	return nil

}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 1, "", nil
}

// AllReservation Return a slice of all reservations
func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil

}

// AllNewReservation Return a slice of new reservations
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

// returns 1 reservation by id
func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {

	var res models.Reservation

	return res, nil
}

// updates a reservation in the database
func (m *testDBRepo) UpdateReservation(u models.Reservation) error {

	return nil
}

func (m *testDBRepo) DeleteReservation(id int) error {

	return nil
}

// updates processed for a reservation by id
func (m *testDBRepo) UpdateProcessedForReservarion(id, processed int) error {

	return nil

}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}

func (m *testDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {

	var restrictions []models.RoomRestriction

	return restrictions, nil
}

// inserts a room restriction
func (m *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {

	return nil

}

// delets a room restriction
func (m *testDBRepo) DeleteBlockByID(id int) error {

	return nil

}

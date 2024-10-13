package repository

import (
	"time"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int,error) 
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start,end time.Time,rooID int) (bool,error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room,error)
	GetUserById(id int) (models.User,error)
	Authenticate(email,testPassword string) (int,string,error)
	UpdateUser(u models.User) error
	AllReservations()([]models.Reservation,error)
	AllNewReservations() ([]models.Reservation, error)
	GetReservationByID(id int) (models.Reservation,error)
	UpdateReservation(u models.Reservation) error
	DeleteReservation(id int) error
	UpdateProcessedForReservarion(id,processed int)error
	AllRooms()([]models.Room,error)
	GetRestrictionsForRoomByDate(roomID int,start,end time.Time)([]models.RoomRestriction,error)
	InsertBlockForRoom(id int, startDate time.Time) error
	DeleteBlockByID(id int) error
}

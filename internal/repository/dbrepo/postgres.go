package dbrepo

import (
	"context"
	"log"
	"time"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// INSERTA A RESERVATION INTO DATABASE
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var newID int

	stmt := `insert into reservations(first_name,last_name,email,phone,start_date,
	         end_date,room_id,created_at,updated_at)
			 values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id `
	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return 0, err
	}

	return newID, nil
}

// nsertRoomRestriction inserta a room restriction into the database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `insert into room_restrictions(start_date,end_date,room_id,reservation_id,restriction_id,created_at,updated_at) 
		values($1,$2,$3,$4,$5,$6,$7)`
	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID return true if availability exists for roomID,and false if no availability
func (m *postgresDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, rooID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		select 
			count(id)
		from
			room_restrictions
		where
		room_id=$1
		and
			$2 < end_date and $3 > start_date;`
	var numRows int
	row := m.DB.QueryRowContext(ctx, query, rooID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}
	if numRows == 0 {
		return true, nil
	}
	return false, nil
}
//earchAvailabilityForAllRooms returns a slice of all available rooms if any for a given date range
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var rooms []models.Room
	query := `
			SELECT r.id, r.room_name 
			FROM 
				rooms r 
			WHERE 
				r.id NOT IN (
				SELECT 
					rr.room_id 
				FROM 
					room_restrictions rr 
				WHERE 
					rr.start_date <= $1 
				AND 
					rr.end_date >= $2
				);
		
	`
	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)

		if err = rows.Err(); err != nil {
			log.Fatal("Error Scanning rows", err)
		}
	}
	return rooms, nil
}
//GetRoomByID gets a room by its id 
func(m *postgresDBRepo)GetRoomByID(id int) (models.Room,error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var room models.Room
	query:=`
		select id,room_name,created_at,updated_at from rooms where id=$1	
	`

	row:=m.DB.QueryRowContext(ctx,query,id)
	err:=row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err!=nil{
		return room,err
	}
	return room,nil

}
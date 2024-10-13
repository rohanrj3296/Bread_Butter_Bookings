package dbrepo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/models"
	"golang.org/x/crypto/bcrypt"
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

// earchAvailabilityForAllRooms returns a slice of all available rooms if any for a given date range
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

// GetRoomByID gets a room by its id
func (m *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var room models.Room
	query := `
		select id,room_name,created_at,updated_at from rooms where id=$1	
	`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		return room, err
	}
	return room, nil

}

// GetUserById gets a user by id
func (m *postgresDBRepo) GetUserById(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		select id,first_name,last_name,email,password,access_level,created_at,updated_at
		from users where id=$1
		`
	row := m.DB.QueryRowContext(ctx, query, id)
	var u models.User
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return u, err
	} else {
		return u, nil
	}
}

// updates a user in the database
func (m *postgresDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		update users set first_name=$1,last_name=$2,email=$3,access_level=$4,updated_at=$5

	`
	_, err := m.DB.ExecContext(ctx, query, u.FirstName, u.LastName, u.Email, u.AccessLevel, time.Now())
	if err != nil {
		return err
	}
	return nil
}

// Authenticate authenticates a user
func (m *postgresDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int
	var hashedPassword string
	row := m.DB.QueryRowContext(ctx, "select id , password from users where email=$1", email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect Password entered!")
	} else if err != nil {
		return 0, "", err
	}
	return id, hashedPassword, nil
}

// AllReservation Return a slice of all reservations
func (m *postgresDBRepo) AllReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservations []models.Reservation
	queru := `select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date,
	r.end_date, r.room_id, r.created_at, r.updated_at,r.processed,
	rm.id, rm.room_name
from reservations r
left join rooms rm on (r.room_id = rm.id)
order by r.start_date asc`

	rows, err := m.DB.QueryContext(ctx, queru)
	if err != nil {
		return reservations, err
	}

	defer rows.Close()
	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Processed,
			&i.Room.ID,
			&i.Room.RoomName,
		)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}
	if err = rows.Err(); err != nil {
		return reservations, err
	}
	return reservations, nil
}

// AllNewReservation Return a slice of new reservations
func (m *postgresDBRepo) AllNewReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservations []models.Reservation
	queru := `select r.id, r.first_name, r.last_name, r.email, r.phone, r.start_date,
	r.end_date, r.room_id, r.created_at, r.updated_at,r.processed,
	rm.id, rm.room_name
from reservations r
left join rooms rm on (r.room_id = rm.id)
where processed=0
order by r.start_date asc`

	rows, err := m.DB.QueryContext(ctx, queru)
	if err != nil {
		return reservations, err
	}

	defer rows.Close()
	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Processed,
			&i.Room.ID,
			&i.Room.RoomName,
		)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}
	if err = rows.Err(); err != nil {
		return reservations, err
	}
	return reservations, nil
}

// returns 1 reservation by id
func (m *postgresDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var res models.Reservation
	query := `
	select r.id,r.first_name,r.last_name,r.email,r.phone,r.start_date,
	r.end_date,r.room_id,r.created_at,r.updated_at,r.processed,
	rm.id,rm.room_name
	from reservations r 
	left join rooms rm on (r.room_id=rm.id)
	where r.id=$1
	`
	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Phone,
		&res.StartDate,
		&res.EndDate,
		&res.RoomID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.Processed,
		&res.Room.ID,
		&res.Room.RoomName,
	)
	if err != nil {
		return res, err
	}
	return res, nil
}

// updates a reservation in the database
func (m *postgresDBRepo) UpdateReservation(u models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		update reservations set first_name=$1,last_name=$2,email=$3,phone=$4,updated_at=$5
		where id = $6

	`
	_, err := m.DB.ExecContext(ctx, query, u.FirstName, u.LastName, u.Email, u.Phone, time.Now(), u.ID)
	if err != nil {
		return err
	}
	return nil
}

// deletes a reservation by id
func (m *postgresDBRepo) DeleteReservation(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		delete from reservations where id=$1
	
	
	`
	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// updates processed for a reservation by id
func (m *postgresDBRepo) UpdateProcessedForReservarion(id, processed int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		update reservations set processed=$1 where id=$2
	
	
	`
	_, err := m.DB.ExecContext(ctx, query, processed, id)
	if err != nil {
		log.Printf("Error updating reservation %d: %v", id, err)
		return err
	}
	return nil
}


func(m *postgresDBRepo) AllRooms()([]models.Room,error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var rooms []models.Room
	query:=`select id,room_name,created_at,updated_at from rooms order by room_name`
	rows,err:=m.DB.QueryContext(ctx,query)
	if err!=nil{
		return rooms,err
	}
	defer rows.Close()
	for rows.Next(){
		var rm models.Room
		err:=rows.Scan(
			&rm.ID,
			&rm.RoomName,
			&rm.CreatedAt,
			&rm.UpdatedAt,
		)
		if err!=nil{
			return rooms,err
		}
		rooms = append(rooms, rm)

	}
	if err = rows.Err();err!=nil{
		return rooms,err
	}
	return rooms,nil
}
//returns restrictions for a room by date range
func (m *postgresDBRepo) GetRestrictionsForRoomByDate(roomID int,start,end time.Time)([]models.RoomRestriction,error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var restrictions []models.RoomRestriction
	query:=`
		select id , coalesce(reservation_id,0),restriction_id,room_id,start_date,end_date
		from room_restrictions where $1<end_date and $2 >=start_date and room_id=$3
		`
	rows,err:=m.DB.QueryContext(ctx,query,start,end,roomID)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()

	for rows.Next(){
		var r models.RoomRestriction
		err:=rows.Scan(
			&r.ID,
			&r.ReservationID,
			&r.RestrictionID,
			&r.RoomID,
			&r.StartDate,
			&r.EndDate,
		)
		if err!=nil{
			return nil,err
		}
		restrictions=append(restrictions, r)
	}
	if err=rows.Err(); err!=nil{
		return nil,err
	}
	return restrictions,nil
	}
	
	
	
//inserts a room restriction
func (m *postgresDBRepo) InsertBlockForRoom(id int,startDate time.Time) error{
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query:=`
		insert into room_restrictions (start_date,end_date,room_id,restriction_id,
		created_at,updated_at) values ($1,$2,$3,$4,$5,$6)
		`
	_,err:=m.DB.ExecContext(ctx,query,startDate,startDate.AddDate(0,0,1),id,2,time.Now(),time.Now())
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil

}

//delets a room restriction
func (m *postgresDBRepo) DeleteBlockByID(id int) error{
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query:=`
		delete from room_restrictions where id=$1
		`
	_,err:=m.DB.ExecContext(ctx,query,id)
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil

}
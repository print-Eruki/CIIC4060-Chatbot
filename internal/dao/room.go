package dao

import (
	"database/sql"
	"fmt"

	"github.com/print-Eruki/CIIC4060-chatbot/internal/model"
)

type RoomDAO struct {
	DB *sql.DB
}

func NewRoomDAO(db *sql.DB) *RoomDAO {
	return &RoomDAO{DB: db}
}

func (dao *RoomDAO) GetRooms() ([]model.Room, error) {
	query := `
	SELECT 
		rid, building, room_number, capacity
	FROM 
		public.room;
	`

	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []model.Room
	for rows.Next() {
		var room model.Room

		err := rows.Scan(&room.Rid, &room.Building, &room.Room_number, &room.Capacity)
		if err != nil {
			fmt.Println("Error when scanning, error:", err)
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func (dao *RoomDAO) GetRoomByID(id uint64) (*model.Room, error) {
	query := `
	SELECT 
		rid, building, room_number, capacity
	FROM 
		public.room 
	WHERE 
		rid = $1;
	`

	var room model.Room

	err := dao.DB.QueryRow(query, id).Scan(
		&room.Rid,
		&room.Building,
		&room.Room_number,
		&room.Capacity,
	)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

// Creates a room record, modifies the newRoom param in memory
func (dao *RoomDAO) CreateRoom(newRoom *model.Room) error {
	query := `
	INSERT INTO public.room 
		(building, room_number, capacity)
	VALUES 
		($1, $2, $3)
	RETURNING 
		rid, building, room_number, capacity;
	`

	err := dao.DB.QueryRow(query,
		newRoom.Building,
		newRoom.Room_number,
		newRoom.Capacity,
	).Scan(
		&newRoom.Rid,
		&newRoom.Building,
		&newRoom.Room_number,
		&newRoom.Capacity,
	)

	return err
}

// Updates the room with param id and modifies the updatedRoom in memory
func (dao *RoomDAO) UpdateRoom(updatedRoom *model.Room, id uint64) error {
	query := `
	UPDATE public.room
	SET 
		building = $1, 
		room_number = $2, 
		capacity = $3
	WHERE 
		rid = $4
	RETURNING 
		rid, building, room_number, capacity;
	`

	err := dao.DB.QueryRow(query,
		updatedRoom.Building,
		updatedRoom.Room_number,
		updatedRoom.Capacity,
		id,
	).Scan(
		&updatedRoom.Rid,
		&updatedRoom.Building,
		&updatedRoom.Room_number,
		&updatedRoom.Capacity,
	)

	return err
}

func (dao *RoomDAO) DeleteRoom(id uint64) (model.Room, error) {
	query := `
	DELETE FROM 
		public.room
	WHERE 
		rid = $1
	RETURNING 
		rid, building, room_number, capacity;
	`
	var deletedRoom model.Room
	fmt.Println(id)
	err := dao.DB.QueryRow(query, id).Scan(
		&deletedRoom.Rid,
		&deletedRoom.Building,
		&deletedRoom.Room_number,
		&deletedRoom.Capacity,
	)

	if err != nil {
		return model.Room{}, err
	}

	return deletedRoom, nil
}

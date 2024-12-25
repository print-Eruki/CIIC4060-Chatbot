package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/print-Eruki/CIIC4060-chatbot/internal/model"
)

type MeetingDAO struct {
	DB *sql.DB
}

func NewMeetingDAO(db *sql.DB) *MeetingDAO {
	return &MeetingDAO{DB: db}
}

func (dao *MeetingDAO) GetMeetings() ([]model.Meeting, error) {
	query := `
	SELECT 
		mid, ccode, starttime, endtime, cdays
	FROM 
		public.meeting;
	`
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meetings []model.Meeting
	for rows.Next() {
		var meeting model.Meeting
		//create temp Time, later convert to the string format wanted
		var tempStartTime time.Time
		var tempEndTime time.Time

		err := rows.Scan(
			&meeting.Mid,
			&meeting.Ccode,
			&tempStartTime,
			&tempEndTime,
			&meeting.Cdays,
		)
		if err != nil {
			fmt.Println("Error when scanning, errr")
		}
		// convert the Time into only HH:MM:SS formatted string
		meeting.Starttime = tempStartTime.Format("15:04:05")
		meeting.Endtime = tempEndTime.Format("15:04:05")
		meetings = append(meetings, meeting)
	}
	return meetings, nil
}

func (dao *MeetingDAO) GetMeetingByID(id uint64) (*model.Meeting, error) {
	query := `
	SELECT
		mid, ccode, starttime, endtime, cdays
	FROM 
		public.meeting
	WHERE 
		mid = $1;
	`

	var meeting model.Meeting
	var tempStartTime time.Time
	var tempEndTime time.Time

	err := dao.DB.QueryRow(query, id).Scan(
		&meeting.Mid,
		&meeting.Ccode,
		&tempStartTime,
		&tempEndTime,
		&meeting.Cdays,
	)
	if err != nil {
		return &model.Meeting{}, err
	}
	meeting.Starttime = tempStartTime.Format("15:04:05")
	meeting.Endtime = tempEndTime.Format("15:04:05")

	return &meeting, nil
}

func (dao *MeetingDAO) CreateMeeting(newMeeting *model.Meeting) error {
	query := `
	INSERT INTO public.meeting
		(ccode, starttime, endtime, cdays)
	VALUES ($1, $2, $3, $4)
	RETURNING 
		mid, ccode, starttime, endtime, cdays
	`
	var tempStartTime time.Time
	var tempEndTime time.Time

	err := dao.DB.QueryRow(query,
		newMeeting.Ccode,
		newMeeting.Starttime,
		newMeeting.Endtime,
		newMeeting.Cdays,
	).Scan(
		&newMeeting.Mid,
		&newMeeting.Ccode,
		&tempStartTime,
		&tempEndTime,
		&newMeeting.Cdays,
	)

	newMeeting.Starttime = tempStartTime.Format("15:04:05")
	newMeeting.Endtime = tempEndTime.Format("15:04:05")

	return err
}

func (dao *MeetingDAO) UpdateMeeting(updatedMeeting *model.Meeting, id uint64) error {
	query := `
	UPDATE public.meeting
	SET
		ccode = $1,
		starttime = $2,
		endtime = $3,
		cdays = $4
	WHERE 
		mid = $5
	RETURNING 
		mid, ccode, starttime, endtime, cdays;
	`
	var tempStartTime time.Time
	var tempEndTime time.Time

	err := dao.DB.QueryRow(query,
		updatedMeeting.Ccode,
		updatedMeeting.Starttime,
		updatedMeeting.Endtime,
		updatedMeeting.Cdays,
		id,
	).Scan(
		&updatedMeeting.Mid,
		&updatedMeeting.Ccode,
		&tempStartTime,
		&tempEndTime,
		&updatedMeeting.Cdays,
	)

	updatedMeeting.Starttime = tempStartTime.Format("15:04:05")
	updatedMeeting.Endtime = tempEndTime.Format("15:04:05")

	return err
}

func (dao *MeetingDAO) DeleteMeeting(id uint64) (*model.Meeting, error) {
	query := `
	DELETE FROM
		public.meeting
	WHERE 
		mid = $1
	RETURNING
		mid, ccode, starttime, endtime, cdays;
	`

	var deletedMeeting model.Meeting
	var tempStartTime time.Time
	var tempEndTime time.Time

	err := dao.DB.QueryRow(query, id).Scan(
		&deletedMeeting.Mid,
		&deletedMeeting.Ccode,
		&tempStartTime,
		&tempEndTime,
		&deletedMeeting.Cdays,
	)

	if err != nil {
		return &model.Meeting{}, err
	}

	deletedMeeting.Starttime = tempStartTime.Format("15:04:05")
	deletedMeeting.Endtime = tempEndTime.Format("15:04:05")

	return &deletedMeeting, nil
}

package dao

import (
	"database/sql"
	"fmt"

	"github.com/print-Eruki/CIIC4060-chatbot/internal/model"
)

type SectionDAO struct {
	DB *sql.DB
}

func NewSectionDAO(db *sql.DB) *SectionDAO {
	return &SectionDAO{DB: db}
}

func (dao *SectionDAO) GetSections() ([]model.Section, error) {
	query := `
	SELECT 
		sid, roomid, mid, cid, semester, years, capacity
	FROM
		public.section;
	`

	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sections []model.Section
	for rows.Next() {
		var section model.Section

		err = rows.Scan(
			&section.Sid,
			&section.Roomid,
			&section.Mid,
			&section.Cid,
			&section.Semester,
			&section.Years,
			&section.Capacity,
		)
		if err != nil {
			fmt.Println(err)
		}

		sections = append(sections, section)
	}
	return sections, nil
}

func (dao *SectionDAO) GetSectionByID(id uint64) (*model.Section, error) {
	query := `
	SELECT 
		sid, roomid, mid, cid, semester, years, capacity
	FROM 
		public.section 
	WHERE 
		sid = $1;
	`

	var section model.Section

	err := dao.DB.QueryRow(query, id).Scan(
		&section.Sid,
		&section.Roomid,
		&section.Mid,
		&section.Cid,
		&section.Semester,
		&section.Years,
		&section.Capacity,
	)

	if err != nil {
		return nil, err
	}

	return &section, nil
}

// Creates a section record, modifies the newSection param in memory
func (dao *SectionDAO) CreateSection(newSection *model.Section) error {
	query := `
	INSERT INTO public.section 
		(roomid, mid, cid, semester, years, capacity)
	VALUES 
		($1, $2, $3, $4, $5, $6)
	RETURNING 
		sid, roomid, mid, cid, semester, years, capacity;
	`

	err := dao.DB.QueryRow(query,
		newSection.Roomid,
		newSection.Mid,
		newSection.Cid,
		newSection.Semester,
		newSection.Years,
		newSection.Capacity,
	).Scan(
		&newSection.Sid,
		&newSection.Roomid,
		&newSection.Mid,
		&newSection.Cid,
		&newSection.Semester,
		&newSection.Years,
		&newSection.Capacity,
	)

	return err
}

// Updates the section with param id and modifies the updatedSection in memory
func (dao *SectionDAO) UpdateSection(updatedSection *model.Section, id uint64) error {
	query := `
	UPDATE public.section
	SET 
		roomid = $1, 
		mid = $2, 
		cid = $3, 
		semester = $4, 
		years = $5, 
		capacity = $6
	WHERE 
		sid = $7
	RETURNING 
		sid, roomid, mid, cid, semester, years, capacity;
	`

	err := dao.DB.QueryRow(query,
		updatedSection.Roomid,
		updatedSection.Mid,
		updatedSection.Cid,
		updatedSection.Semester,
		updatedSection.Years,
		updatedSection.Capacity,
		id,
	).Scan(
		&updatedSection.Sid,
		&updatedSection.Roomid,
		&updatedSection.Mid,
		&updatedSection.Cid,
		&updatedSection.Semester,
		&updatedSection.Years,
		&updatedSection.Capacity,
	)

	return err
}

// Deletes the section with param id, returns the deleted record
func (dao *SectionDAO) DeleteSection(id uint64) (model.Section, error) {
	query := `
	DELETE FROM 
		public.section
	WHERE 
		sid = $1
	RETURNING 
		sid, roomid, mid, cid, semester, years, capacity;
	`
	var deletedSection model.Section
	fmt.Println(id)
	err := dao.DB.QueryRow(query, id).Scan(
		&deletedSection.Sid,
		&deletedSection.Roomid,
		&deletedSection.Mid,
		&deletedSection.Cid,
		&deletedSection.Semester,
		&deletedSection.Years,
		&deletedSection.Capacity,
	)

	// Didnt found a section with param id
	if err != nil {
		return model.Section{}, err
	}

	return deletedSection, nil
}

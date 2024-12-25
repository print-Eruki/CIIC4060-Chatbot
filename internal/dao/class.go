package dao

import (
	"database/sql"
	"fmt"

	"github.com/print-Eruki/CIIC4060-chatbot/internal/model"
)

type ClassDAO struct {
	Db *sql.DB
}

func NewClassDAO(db *sql.DB) *ClassDAO {
	return &ClassDAO{Db: db}
}

// Selects * from class, if there is an error when running the query it returns the error
func (dao *ClassDAO) GetClasses() ([]model.Class, error) {
	query := `
	SELECT 
		ccode, cname, cid, cred, cdesc, csyllabus, term, years 
	FROM 
		public.class;
	`

	rows, err := dao.Db.Query(query)
	if err != nil {
		//return a nil and the error, let the handler take care of it
		return nil, err
	}
	defer rows.Close()

	var classes []model.Class
	for rows.Next() {
		var class model.Class

		err := rows.Scan(&class.Ccode, &class.Cname, &class.Cid, &class.Cred,
			&class.Cdesc, &class.Csyllabus, &class.Term, &class.Years)
		if err != nil {
			fmt.Println("Error when scanning, error:", err)
		}
		classes = append(classes, class)
	}
	return classes, nil
}

func (dao *ClassDAO) GetClassByID(id uint64) (*model.Class, error) {
	query := `
	SELECT 
		ccode, cname, cid, cred, cdesc, csyllabus, term, years 
	FROM 
		public.class 
	WHERE 
		cid = $1;
	`

	var class model.Class

	err := dao.Db.QueryRow(query, id).Scan(
		&class.Ccode,
		&class.Cname,
		&class.Cid,
		&class.Cred,
		&class.Cdesc,
		&class.Csyllabus,
		&class.Term,
		&class.Years,
	)

	if err != nil {
		return nil, err
	}

	return &class, nil
}

// Creates a class record, modifies the newClass param in memory
func (dao *ClassDAO) CreateClass(newClass *model.Class) error {
	query := `
	INSERT INTO public.class 
		(ccode, cdesc, cname, cred, csyllabus, term, years)
	VALUES 
		($1, $2, $3, $4, $5, $6, $7)
	RETURNING 
		ccode, cname, cid, cred, cdesc, csyllabus, term, years;
	`

	err := dao.Db.QueryRow(query,
		newClass.Ccode,
		newClass.Cdesc,
		newClass.Cname,
		newClass.Cred,
		newClass.Csyllabus,
		newClass.Term,
		newClass.Years,
	).Scan(
		&newClass.Ccode,
		&newClass.Cname,
		&newClass.Cid,
		&newClass.Cred,
		&newClass.Cdesc,
		&newClass.Csyllabus,
		&newClass.Term,
		&newClass.Years,
	)

	return err
}

// Updates the class with param id and modifies the updatedClass in memory
func (dao *ClassDAO) UpdateClass(updatedClass *model.Class, id uint64) error {
	query := `
	UPDATE public.class
	SET 
		ccode = $1, 
		cdesc = $2, 
		cname = $3, 
		cred = $4, 
		csyllabus = $5, 
		term = $6, 
		years = $7
	WHERE 
		cid = $8
	RETURNING 
		ccode, cname, cid, cred, cdesc, csyllabus, term, years;
	`

	err := dao.Db.QueryRow(query,
		updatedClass.Ccode,
		updatedClass.Cdesc,
		updatedClass.Cname,
		updatedClass.Cred,
		updatedClass.Csyllabus,
		updatedClass.Term,
		updatedClass.Years,
		id,
	).Scan(
		&updatedClass.Ccode,
		&updatedClass.Cname,
		&updatedClass.Cid,
		&updatedClass.Cred,
		&updatedClass.Cdesc,
		&updatedClass.Csyllabus,
		&updatedClass.Term,
		&updatedClass.Years,
	)

	return err
}

// Deletes the class with param id, returns the deleted record
func (dao *ClassDAO) DeleteClass(id uint64) (model.Class, error) {
	query := `
	DELETE FROM 
		public.class
	WHERE 
		cid = $1
	RETURNING 
		ccode, cname, cid, cred, cdesc, csyllabus, term, years;
	`
	var deletedClass model.Class
	fmt.Println(id)
	err := dao.Db.QueryRow(query, id).Scan(
		&deletedClass.Ccode,
		&deletedClass.Cname,
		&deletedClass.Cid,
		&deletedClass.Cred,
		&deletedClass.Cdesc,
		&deletedClass.Csyllabus,
		&deletedClass.Term,
		&deletedClass.Years,
	)

	// Didnt found a class with param id
	if err != nil {
		return model.Class{}, err
	}

	return deletedClass, nil
}

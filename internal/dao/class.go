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
	query := "SELECT ccode, cname, cid, cred, cdesc, csyllabus, term, years FROM public.class;"

	rows, err := dao.Db.Query(query)
	if err != nil {
		//return a nil and the error, let the handler take care of it
		return nil, err
	}
	defer rows.Close()

	var classes []model.Class
	for rows.Next() {
		var class model.Class

		err := rows.Scan(&class.Cid, &class.Ccode, &class.Cname, &class.Cred, &class.Cdesc, &class.Csyllabus, &class.Term, &class.Years)
		if err != nil {
			fmt.Println("Error when scanning, error:", err)
		}
		classes = append(classes, class)
	}
	return classes, nil
}

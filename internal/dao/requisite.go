package dao

import (
	"database/sql"
	"fmt"

	"github.com/print-Eruki/CIIC4060-chatbot/internal/model"
)

type RequisiteDAO struct {
	DB *sql.DB
}

func NewRequisiteDAO(db *sql.DB) *RequisiteDAO {
	return &RequisiteDAO{DB: db}
}

func (dao *RequisiteDAO) GetRequisites() ([]model.Requisite, error) {
	query := `
	SELECT
		classid, reqid, prereq
	FROM 
		public.requisite;
	`

	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var requisites []model.Requisite
	for rows.Next() {
		var r model.Requisite
		err := rows.Scan(
			&r.Classid,
			&r.Reqid,
			&r.Prereq,
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		requisites = append(requisites, r)
	}
	return requisites, nil
}

// first param is the class id, second param is the requisite id
func (dao *RequisiteDAO) GetRequisiteByID(cIdParam uint64, rIdParam uint64) (model.Requisite, error) {
	query := `
	SELECT 
		classid, reqid, prereq
	FROM 
		public.requisite
	WHERE
		classid = $1 and reqid = $2;
	`

	var requisite model.Requisite
	err := dao.DB.QueryRow(query, cIdParam, rIdParam).Scan(
		&requisite.Classid,
		&requisite.Reqid,
		&requisite.Prereq,
	)
	if err != nil {
		return model.Requisite{}, err
	}

	return requisite, nil
}

func (dao *RequisiteDAO) CreateRequisite(newRequisite *model.Requisite) error {
	query := `
	INSERT INTO public.requisite
		(classid, reqid, prereq)
	VALUES
		($1, $2, $3)
	RETURNING
		classid, reqid, prereq;
	`

	err := dao.DB.QueryRow(query,
		newRequisite.Classid,
		newRequisite.Reqid,
		newRequisite.Prereq,
	).Scan(
		&newRequisite.Classid,
		&newRequisite.Reqid,
		&newRequisite.Prereq,
	)

	return err
}

func (dao *RequisiteDAO) UpdateRequisite(r *model.Requisite, cIdParam uint64, rIdParam uint64) error {
	query := `
	UPDATE public.requisite
	SET
		classid = $1,
		reqid = $2,
		prereq = $3
	WHERE
		classid = $4 and reqid = $5
	RETURNING
		classid, reqid, prereq;
	`

	err := dao.DB.QueryRow(query,
		r.Classid,
		r.Reqid,
		r.Prereq,
		cIdParam,
		rIdParam,
	).Scan(
		&r.Classid,
		&r.Reqid,
		&r.Prereq,
	)

	return err
}

func (dao *RequisiteDAO) DeleteRequisite(cIdParam uint64, rIdParam uint64) (model.Requisite, error) {
	query := `
	DELETE FROM 
		public.requisite
	WHERE 
		classid = $1 and reqid = $2
	RETURNING 
		classid, reqid, prereq;
	`
	var deletedRequisite model.Requisite

	err := dao.DB.QueryRow(query, cIdParam, rIdParam).Scan(
		&deletedRequisite.Classid,
		&deletedRequisite.Reqid,
		&deletedRequisite.Prereq,
	)
	if err != nil {
		return model.Requisite{}, err
	}

	return deletedRequisite, nil

}

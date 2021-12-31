package repositories

import (
	"challenge_fullstack/src/models"
	"challenge_fullstack/src/utilities"
	"context"
	"database/sql"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type RepoDevelopers struct {
	ctx  *context.Context
	conn *sql.Conn
}

func NewRepoDevelopers(ctx *context.Context, conn *sql.Conn) *RepoDevelopers {
	return &RepoDevelopers{ctx, conn}
}

func (db RepoDevelopers) Insert(d *models.Developers) error {
	res, err := db.conn.ExecContext(
		*db.ctx,
		`INSERT INTO developers(dev_name, dev_gender, dev_birth, dev_hobby, lvl_id)
		 VALUES(?, ?, ?, ?, ?)`,
		d.Name, d.Gender, d.Birth, d.Hobby, d.Levels.ID,
	)
	if err != nil {
		requestError := tryDatabaseError(err)
		return &requestError
	}
	id, err := res.LastInsertId()
	if err != nil {
		requestError := tryDatabaseError(err)
		return &requestError
	}
	d.ID = uint32(id)

	return err
}

func (db RepoDevelopers) Update(d *models.Developers, id string) error {
	_, err := db.conn.ExecContext(
		*db.ctx,
		`UPDATE developers SET dev_name = ?, dev_gender = ?, dev_birth = ?, dev_hobby = ?, lvl_id = ?
		WHERE dev_id = ?`,
		d.Name, d.Gender, d.Birth, d.Hobby, d.Levels.ID,
		id,
	)
	if err != nil {
		requestError := tryDatabaseError(err)
		return &requestError
	}

	return nil
}

func (db RepoDevelopers) Delete(id string) error {
	_, err := db.conn.ExecContext(
		*db.ctx,
		`DELETE FROM developers WHERE dev_id = ?`,
		id,
	)
	if err != nil {
		requestError := tryDatabaseError(err)
		return &requestError
	}

	return nil
}

//Query retorna uma consulta do banco de dados
func (db RepoDevelopers) Query(q *utilities.Query) error {
	var developerModel models.Developers

	q.Model = &developerModel

	q.SqlQuery = `SELECT D.dev_id, D.dev_name, D.dev_gender, D.dev_birth,
	              D.dev_hobby, L.lvl_id, L.lvl_name
                  FROM developers D, levels L`
	q.AddWhere("D.lvl_id = L.lvl_id", "", nil)
	q.IsPagination = true
	q.NickName = `D`
	q.OrderByDefault = `dev_name`

	levelId := q.C.Query("level_id")
	if levelId != "" {
		q.AddWhere("L.lvl_id = ?", levelId, reflect.TypeOf((*uint32)(nil)).Elem())
	}

	q.Mount()

	rows, err := db.conn.QueryContext(
		*db.ctx,
		q.SqlResult, q.Args...,
	)

	if err != nil {
		requestError := tryDatabaseError(err)
		return &requestError
	}

	defer rows.Close()

	for rows.Next() {
		var developer models.Developers
		err = rows.Scan(&developer.ID, &developer.Name, &developer.Gender, &developer.Birth,
			&developer.Hobby, &developer.Levels.ID, &developer.Levels.Name)
		if err != nil {
			requestError := tryDatabaseError(err)
			return &requestError
		}

		q.Response.ResultSet = append(q.Response.ResultSet, developer)
	}

	q.Context = db.ctx
	q.Conn = db.conn

	q.Count()

	return nil
}

func (db RepoDevelopers) QueryByID(id string) (models.Developers, error) {

	row := db.conn.QueryRowContext(
		*db.ctx,
		`SELECT d.dev_id, d.dev_name, d.dev_gender, d.dev_birth, d.dev_hobby,
		        l.lvl_id, l.lvl_name
		FROM developers d, levels l
		WHERE d.lvl_id = l.lvl_id
		  AND d.dev_id = ?`, id,
	)
	var developer models.Developers
	err := row.Scan(&developer.ID, &developer.Name, &developer.Gender, &developer.Birth, &developer.Hobby,
		&developer.Levels.ID, &developer.Levels.Name)
	if err != nil {
		return developer, &models.RequestError{fiber.StatusNotFound, err}
	}
	return developer, err
}

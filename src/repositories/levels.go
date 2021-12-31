package repositories

import (
	"challenge_fullstack/src/models"
	"challenge_fullstack/src/utilities"
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type RepoLevels struct {
	ctx  *context.Context
	conn *sql.Conn
}

func NewRepoLevels(ctx *context.Context, conn *sql.Conn) *RepoLevels {
	return &RepoLevels{ctx, conn}
}

func (db *RepoLevels) Insert(l *models.Levels) error {
	res, err := db.conn.ExecContext(
		*db.ctx,
		`INSERT INTO levels(lvl_name)
		 VALUES(?)`,
		l.Name,
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
	l.ID = uint32(id)

	return nil
}

func (db *RepoLevels) Update(l *models.Levels, id string) error {
	_, err := db.conn.ExecContext(
		*db.ctx,
		`UPDATE levels SET lvl_name = ?
		WHERE lvl_id = ?`,
		l.Name,
		id,
	)
	if err != nil {
		requestError := models.RequestError{}
		requestError.StatusCode = fiber.StatusInternalServerError
		requestError.Err = err
		return &requestError
	}

	return nil
}

func (db *RepoLevels) Delete(id string) error {
	_, err := db.conn.ExecContext(
		*db.ctx,
		`DELETE FROM levels WHERE lvl_id = ?`,
		id,
	)
	if err != nil {
		requestError := tryDatabaseError(err)
		return &requestError
	}

	return nil
}

//Query retorna uma consulta do banco de dados
func (db *RepoLevels) Query(q *utilities.Query) error {
	var levelModel models.Levels

	q.Model = &levelModel

	q.SqlQuery = `SELECT L.lvl_id, L.lvl_name
                  FROM levels L`
	q.IsPagination = true
	q.NickName = `L`
	q.OrderByDefault = `lvl_name`

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
		var level models.Levels
		err = rows.Scan(&level.ID, &level.Name)
		if err != nil {
			requestError := tryDatabaseError(err)
			return &requestError
		}
		q.Response.ResultSet = append(q.Response.ResultSet, level)
	}
	q.Context = db.ctx
	q.Conn = db.conn

	q.Count()

	return nil
}

func (db *RepoLevels) QueryByID(id string) (models.Levels, error) {

	row := db.conn.QueryRowContext(
		*db.ctx,
		`SELECT lvl_id, lvl_name
		 FROM levels l
		 WHERE lvl_id = ?`, id,
	)
	var level models.Levels

	err := row.Scan(&level.ID, &level.Name)

	if err != nil {
		return level, &models.RequestError{fiber.StatusNotFound, err}
	}
	return level, err
}

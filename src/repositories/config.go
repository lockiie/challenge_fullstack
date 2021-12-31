package repositories

import (
	"challenge_fullstack/src/models"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func tryDatabaseError(err error) (requestError models.RequestError) {
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if ok {
			switch driverErr.Number {
			case 1451:
				{
					requestError.Err = errors.New("Record has dependent data")
					requestError.StatusCode = fiber.StatusNotImplemented
					return requestError
				}
			case 1452:
				{
					requestError.Err = errors.New("Dependent record not found")
					requestError.StatusCode = fiber.StatusNotImplemented
					return requestError

				}
			default:
				{
					requestError.Err = errors.New(err.Error())
					requestError.StatusCode = fiber.StatusInternalServerError
					return requestError
				}
			}
		}
	}
	requestError.Err = errors.New(err.Error())
	requestError.StatusCode = fiber.StatusInternalServerError
	return requestError

}

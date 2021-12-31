package controllers

import (
	"challenge_fullstack/src/db"
	"challenge_fullstack/src/models"
	"challenge_fullstack/src/repositories"
	"challenge_fullstack/src/utilities"
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func CreateDevelopers(c *fiber.Ctx) error {
	d := models.Developers{}

	err := json.Unmarshal(c.Body(), &d)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(messageInvalidJson + err.Error()))
	}

	if err = d.Validators(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	repoDeveloper := repositories.NewRepoDevelopers(&ctx, conn)

	if err := repoDeveloper.Insert(&d); err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(d)
}

func UpdateDevelopers(c *fiber.Ctx) error {
	d := models.Developers{}

	err := json.Unmarshal(c.Body(), &d)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(messageInvalidJson + err.Error()))
	}

	if err = d.Validators(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}

	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	id := c.Params(utilities.ParamIDValue)

	repoDeveloper := repositories.NewRepoDevelopers(&ctx, conn)

	_, err = repoDeveloper.QueryByID(id)

	if err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}

	if err := repoDeveloper.Update(&d, id); err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}

	return c.SendStatus(fiber.StatusNoContent)

}

func DeleteDevelopers(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	id := c.Params(utilities.ParamIDValue)

	repoDeveloper := repositories.NewRepoDevelopers(&ctx, conn)

	_, err = repoDeveloper.QueryByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.SendError(err.Error()))
	}

	if err := repoDeveloper.Delete(id); err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))

	}

	return c.SendStatus(fiber.StatusNoContent)
}

func QueryDevelopers(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()
	repoDevelopers := repositories.NewRepoDevelopers(&ctx, conn)

	var query utilities.Query

	query.C = c

	err = repoDevelopers.Query(&query)

	if err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(query.Response)
}

func QueryByIDDevelopers(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	repoDeveloper := repositories.NewRepoDevelopers(&ctx, conn)

	id := c.Params(utilities.ParamIDValue)

	brand, err := repoDeveloper.QueryByID(id)

	if err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(brand)
}

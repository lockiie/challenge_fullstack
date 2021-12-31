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

func CreateLevels(c *fiber.Ctx) error {
	d := models.Levels{}

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

	repoDeveloper := repositories.NewRepoLevels(&ctx, conn)

	if err := repoDeveloper.Insert(&d); err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(d)
}

func UpdateLevels(c *fiber.Ctx) error {
	d := models.Levels{}

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

	repoLevels := repositories.NewRepoLevels(&ctx, conn)

	_, err = repoLevels.QueryByID(id)

	if err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}

	if err := repoLevels.Update(&d, id); err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}

	return c.SendStatus(fiber.StatusNoContent)

}

func DeleteLevels(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	id := c.Params(utilities.ParamIDValue)

	repoLevels := repositories.NewRepoLevels(&ctx, conn)

	_, err = repoLevels.QueryByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.SendError(err.Error()))
	}

	if err := repoLevels.Delete(id); err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func QueryLevels(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()
	repoLevels := repositories.NewRepoLevels(&ctx, conn)

	var query utilities.Query

	query.C = c

	err = repoLevels.Query(&query)

	if err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(query.Response)
}

func QueryByIDLevels(c *fiber.Ctx) error {
	var ctx = context.Background()
	conn, err := db.Pool.Conn(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.SendError(err.Error()))
	}
	defer conn.Close()

	repoLevels := repositories.NewRepoLevels(&ctx, conn)

	id := c.Params(utilities.ParamIDValue)

	level, err := repoLevels.QueryByID(id)

	if err != nil {
		errRequestError := err.(*models.RequestError)
		return c.Status(errRequestError.StatusCode).JSON(models.SendError(errRequestError.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(level)
}

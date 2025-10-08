package main

import (
	"_lab/data-pump/common"
	// "_lab/data-pump/db"
	"_lab/data-pump/helpers"
	"_lab/data-pump/models"
	"encoding/json"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

const (
	_ = iota
	dbFault
	chainFault
	serverFault
	otherFault
)

type response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Code    int    `json:"error_code,omitempty"`
	Message string `json:"error_message,omitempty"`
}

func main() {
	client, err := ethclient.Dial(common.ProviderURL)
	if err != nil {
		log.Fatalf("\033[31m[ERR]\033[0m Failed to connect client: %v", err)
	}

	registry := common.NewRegistry()
	instance := registry.Instance(client, common.ContractAddress)

	// dbConn, err := db.Connect()
	// if err != nil {
	// 	log.Fatal("\033[31m[ERR]\033[0m Failed to connect the database")
	// }

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.Get("/ping", func(c *fiber.Ctx) error {
				return c.SendString("pong")
			})
			// v1.Get("/random", func(c *fiber.Ctx) error {
			// 	return fetchRandomData(c, registry, instance, dbConn)
			// })
			v1.Get("/single/:id", func(c *fiber.Ctx) error {
				return fetchSingleData(c, registry, instance)
			})
		}
	}

	log.Fatalln(app.Listen(":" + os.Getenv("API_PORT")))
}

func fetchRandomData(c *fiber.Ctx, registry *common.Registry, instance *bind.BoundContract, dbConn *gorm.DB) error {
	var dbEntry models.Entry
	if err := dbConn.Order("RANDOM()").Preload("Ownership").Preload("Properties").First(&dbEntry).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Success: false,
			Code:    dbFault,
			Message: err.Error(),
		})
	}

	data, err := bind.Call(instance, nil, registry.PackGetLatestProperty(dbEntry.CardNumber), registry.UnpackGetLatestProperty)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Success: false,
			Code:    chainFault,
			Message: err.Error(),
		})
	}

	var bcEntry models.Entry
	if err := json.Unmarshal([]byte(data), &bcEntry); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Success: false,
			Code:    serverFault,
			Message: err.Error(),
		})
	}

	if !helpers.IsSubset(dbEntry, bcEntry) {
		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Success: false,
			Code:    otherFault,
			Message: "Data mismatch detected",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response{
		Success: true,
		Data:    bcEntry,
	})
}

func fetchSingleData(c *fiber.Ctx, registry *common.Registry, instance *bind.BoundContract) error {
	id := c.Params("id")
	data, err := bind.Call(instance, nil, registry.PackGetLatestProperty(id), registry.UnpackGetLatestProperty)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Success: false,
			Code:    chainFault,
			Message: err.Error(),
		})
	}

	if len(data) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(response{
			Success: false,
			Code:    otherFault,
			Message: "No data found",
		})
	}

	var bcEntry models.Entry
	if err := json.Unmarshal([]byte(data), &bcEntry); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response{
			Success: false,
			Code:    serverFault,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response{
		Success: true,
		Data:    bcEntry,
	})
}

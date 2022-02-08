package controllers

import (
	"echo-mongo-api/configs"
	"echo-mongo-api/models"
	"echo-mongo-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var playerCollection *mongo.Collection = configs.GetCollection(configs.DB, "players")
var validate = validator.New()

// CreatePlayer ...
func CreatePlayer(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var player models.Player
	defer cancel()

	// validate the request body
	if err := c.Bind(&player); err != nil {
		return c.JSON(http.StatusBadRequest, responses.PlayerResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&player); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.PlayerResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newPlayer := models.Player{
		Id: primitive.NewObjectID(),
		Name: player.Name,
		Age: player.Age,
		CreateAt: time.Now(),
	}

	result, err := playerCollection.InsertOne(ctx, newPlayer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.PlayerResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

// GetAPlayer ...
func GetAPlayer(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	playerId := c.Param("playerId")
	var player models.Player
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(playerId)

	err := playerCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&player)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.PlayerResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": player}})
}

// UpdatePlayer ...
func UpdatePlayer(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	playerId := c.Param("playerId")
	var player models.Player
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(playerId)

	//validate the request body
	if err := c.Bind(&player); err != nil {
		return c.JSON(http.StatusBadRequest, responses.PlayerResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&player); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.PlayerResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": player.Name, "Age": player.Age}

	result, err := playerCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//get updated player details
	var updatedPlayer models.Player
	if result.MatchedCount == 1 {
		err := playerCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedPlayer)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
	}

	return c.JSON(http.StatusOK, responses.PlayerResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": updatedPlayer}})
}

// GetAllPlayers ...
func GetAllPlayers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var players []models.Player
	defer cancel()

	results, err := playerCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singlePlayer models.Player
		if err = results.Decode(&singlePlayer); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		players = append(players, singlePlayer)
	}

	return c.JSON(http.StatusInternalServerError, responses.PlayerResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": players}})
}
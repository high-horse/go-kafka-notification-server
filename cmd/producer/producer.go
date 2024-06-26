package main

import (
	"encoding/json"
	"errors"
	"fmt"
	models "kafka-notify/pkg"
	"log"
	"net/http"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

const (
	ProducerPort       = ":8080"
	KafkaServerAddress = "localhost:9092"
	KafkaTopic         = "notifications"
)

// HELPER FUNCTION
var ErrUserNotFoundInProducer = errors.New("user not found")

func findUserById(id int, users []models.User) (models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, ErrUserNotFoundInProducer
}

func getIDFromRequest(fromValue string, ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.PostForm(fromValue))
	if err != nil {
		return 0, fmt.Errorf(
			"failed to parse ID from form value %s: %w", fromValue, err,
		)
	}
	return id, nil
}


// KAFKA functions
func sendKafkaMesage(
	producer sarama.SyncProducer,
	users []models.User,
	ctx *gin.Context,
	fromID , toID int,
) error {
	message := ctx.PostForm("message")

	fromUser, err := findUserById(fromID, users)
	if err != nil {
		return err
	}

	toUser, err := findUserById(toID, users)
	if err != nil {
		return err
	}

	notification := models.Notification {
		From: fromUser,
		To: toUser,
		Message: message,
	}

	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage {
		Topic: KafkaTopic,
		Key: sarama.StringEncoder(strconv.Itoa(toUser.ID)),
		Value: sarama.StringEncoder(notificationJSON),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}

func sendMessageHandler(
	producer sarama.SyncProducer,
	users []models.User,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fromID, err := getIDFromRequest("fromID", ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		toID, err := getIDFromRequest("toID", ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err  = sendKafkaMesage(producer, users, ctx, fromID, toID)
		if errors.Is(err, ErrUserNotFoundInProducer) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "notification sent successfully"})
	}
}

func setupProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{KafkaServerAddress}, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}



func main() {
	users := []models.User {
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
		{ID: 3, Name: "Jill"},
		{ID: 4, Name: "Jack"},
		{ID: 5, Name: "Judy"},
	}

	producer, err := setupProducer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}

	defer producer.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/send", sendMessageHandler(producer, users))

	fmt.Printf("Kafka PRODUCER 📨 started at http://localhost%s\n", ProducerPort)

	if err := router.Run(ProducerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
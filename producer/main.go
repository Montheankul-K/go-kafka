package main

import (
	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"producer/controllers"
	"producer/services"
	"strings"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewAccountController(accountService)

	app := fiber.New()

	app.Post("/openAccount", accountController.OpenAccount)
	app.Post("/depositFund", accountController.DepositFund)
	app.Post("/withdrawFund", accountController.WithdrawFund)
	app.Post("closeAccount", accountController.CloseAccount)

	app.Listen(":8000")
}

/*
func main() {
	servers := []string{"localhost:9093"}
	producer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: "message",
		Value: sarama.StringEncoder("hello world"),
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("partition = %v, offset = %v\n", partition, offset)
}
*/

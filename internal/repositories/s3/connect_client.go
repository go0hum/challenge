package s3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ConnectClient() (client *s3.Client, err error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal("Error cargando config: " + err.Error())
	}

	return s3.NewFromConfig(cfg), nil
}

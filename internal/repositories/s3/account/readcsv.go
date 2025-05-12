package account

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (r *Repository) ReadCSV(ctx context.Context) (*s3.GetObjectOutput, error) {
	resp, err := r.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.Bucket),
		Key:    aws.String(r.Key),
	})
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	log.Println("Lectura con exito")

	return resp, nil
}

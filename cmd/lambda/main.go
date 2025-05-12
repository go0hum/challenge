package main

import (
	"challenge/cmd/lambda/handlers/account"
	"challenge/internal/repositories/mailhog"
	accountEmailRepository "challenge/internal/repositories/mailhog/account"
	"challenge/internal/repositories/mysql"
	accountDatabaseRepository "challenge/internal/repositories/mysql/account"
	"challenge/internal/repositories/s3"
	accountRepository "challenge/internal/repositories/s3/account"
	accountService "challenge/internal/services/account"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	/*err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}*/

	fmt.Println("Inicia main")

	client, err := s3.ConnectClient()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Carga  s3")

	connect, err := mysql.ConnectClient(os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASS"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Carga mysql")

	emailConnect := mailhog.ConnectClient(os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_SMTP_LINK"))

	fmt.Println("Carga email")

	accountRepository := &accountRepository.Repository{
		Client: client,
		Bucket: os.Getenv("AWS_S3_BUCKET"),
		Key:    os.Getenv("AWS_S3_FILE"),
	}

	fmt.Println("configura s3")

	accountDatabaseRepository := &accountDatabaseRepository.RepositoryMysql{
		Db: connect,
	}

	accountEmailRepository := &accountEmailRepository.RepositoryEmail{
		EmailConnect: emailConnect,
		Theme:        os.Getenv("EMAIL_THEME"),
		From:         os.Getenv("EMAIL_FROM"),
		To:           []string{os.Getenv("EMAIL_TO")},
		Subject:      os.Getenv("EMAIL_SUBJECT"),
		SmtpLink:     os.Getenv("EMAIL_SMTP_LINK"),
	}

	accountService := &accountService.Service{
		Account:  accountRepository,
		Database: accountDatabaseRepository,
		Email:    accountEmailRepository,
		Logo:     os.Getenv("EMAIL_THEME_LOGO"),
	}

	fmt.Println("Preregistro")

	accountHandler := &account.Handler{
		AccountService: accountService,
	}

	fmt.Println("LLamada a lambda")

	//accountHandler.HandleRequest(context.Background())
	lambda.Start(accountHandler.HandleRequest)
}

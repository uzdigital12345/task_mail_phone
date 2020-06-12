package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
	"github.com/jmoiron/sqlx"

	"github.com/uzdigital12345/task_mail_phone/storage"
	"github.com/uzdigital12345/task_mail_phone/pkg/sms"
	"github.com/uzdigital12345/task_mail_phone/pkg/mail"
)

const (
	PostgresDatabase = "alief_tech"
	PostgresUser     = "delever"
	PostgresPassword = "delever"
)

// NewSqlx ...
func NewDatabaseConnection() error {

	psqlString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		PostgresUser,
		PostgresPassword,
		PostgresDatabase)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		fmt.Println("Error while connecting database: %v", err)
		return err
	}

	storage.NewStoragePg(connDB)

	return nil
}

func main() {

	var client http.Client
	err := NewDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	s :=sms.New(client)
	m := mail.New("Send mail")

	informationBody := "Your purchases \nProduct: \n   Lavash \n Price: 2$ "

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "sms",
			Action: func(c *cli.Context) error {
				s.SendToPhone(informationBody)
				return nil
			},
		},
		{
			Name: "mail",
			Action: func(c *cli.Context) error {
				err := m.SendToMail(informationBody)
				if err !=nil {
					fmt.Println("Error while sending sms Error: ",err)
					err := storage.StorageI.Error(storage.StoragePg{}).SaveError(err.Error())
					if err !=nil {
						log.Fatal("Error while saving database",err)
					}
				}
				return nil
			},
		},
	}
	app.Run(os.Args)

}

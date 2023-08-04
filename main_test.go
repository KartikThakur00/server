package main

import (
	"database/sql"
	"log"
	"react-gin/server/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)



func Test_POST(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("error init mock", err)
	}

	defer db.Close()

	svc := Service{
		database: db,
	}

	from := time.Now()
	to := time.Now().AddDate(1, 0, 0)

	employee := models.Employee{
		Id:     1,
		Name:   "Test",
		Gender: "male",
		FromDate:   from,
		ToDate: to,
		Phone:  1234567890,
		Resume: "resume.pdf",
		Email:  "demo@gmail.com",
	}
	mock.ExpectExec(`INSERT INTO`)

}

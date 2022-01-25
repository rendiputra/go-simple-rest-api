package databases

import (
	"github.com/rendiputra/go-simple-rest-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@/db_go_auth?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal koneksi dengan database!", err.Error())
	}

	log.Println("Sukses koneksi dengan database!")
	db.AutoMigrate(&models.Task{})
	DBConn = db
}

package main

import(
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User{
	gorm.Model
	Email string `json"Email"`
}

func InitDatabase(){
	DBConn, err = gorm.Open("postgres","rest.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer DBConn.close()
	DBConn.AutoMigrate(&User{})
}


func main() {
	app := fiber.New()
	InitDatabase()
}
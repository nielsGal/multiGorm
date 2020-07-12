package main

import(
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct{
	gorm.Model
	Email string `json:"Email"`
}
var (
	DBConn *gorm.DB
)

func InitDatabase(){
	var err error
	DBConn, err = gorm.Open("postgres","host=localhost port=5432 user=niels dbname=test_database password=galjaard sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}
	defer DBConn.Close()
	DBConn.AutoMigrate(&User{})
}


func main() {
	app := fiber.New()
	InitDatabase()
	app.Listen(3000)
}
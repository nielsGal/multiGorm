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
	DBConn, err = gorm.Open("postgres","host=localhost port=5432 user=postgres dbname=postgres password=example sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}
	
	DBConn.AutoMigrate(&User{})
}

func getUsers(c *fiber.Ctx){
	db := DBConn
	var users []User
	db.Find(&users)
	c.JSON(users)
}

func main() {
	app := fiber.New()
	app.Get("/users",getUsers)
	InitDatabase()
	defer DBConn.Close()
	app.Listen(3000)
}
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

func CreateUser(c *fiber.Ctx){
	db := DBConn
	user := new(User)
	if err := c.BodyParser(user); err != nil{
		c.Status(422).Send("some isse with this request")
	}
	if result := db.Create(user); result.Error != nil{
		c.Status(422).Send("some isse with this request")
	}
	c.JSON(user)
}

func main() {
	app := fiber.New()
	app.Post("/user",CreateUser)
	InitDatabase()
	defer DBConn.Close()
	app.Listen(3001)
	
}
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

	"github.com/DaniilShd/test2/internal/driver"
	"github.com/DaniilShd/test2/internal/handlers"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// const (
// 	portNumber, er := os.LookupEnv("PORT_NUMBER")
// 	passwordDB     = "root"
// 	host           = "localhost"
// 	nameDB         = "persons"
// 	userDB         = "postgres"
// 	portDB         = "5432"
// )

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	engine := html.New("../templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//app.Use(adaptor.HTTPMiddleware(logMiddleware))

	app.Post("/", handlers.Repo.AddPerson)
	app.Get("/", handlers.Repo.GetPersons)
	app.Delete("/:id<int;min(0)>", handlers.Repo.DeletePersonByID)
	app.Put("/:id<int;min(0)>", handlers.Repo.ChangePersonByID)

	portNumber, er := os.LookupEnv("PORT_NUMBER")
	fmt.Println(portNumber)
	if !er {
		log.Fatal("No env")
	}

	log.Fatal(app.Listen(portNumber))
}

func run() (*driver.DB, error) {

	//connect to database
	log.Println("Connecting to database...")

	host, er := os.LookupEnv("HOST")
	if !er {
		log.Fatal("No env")
	}
	portDB, er := os.LookupEnv("PORT_DB")
	fmt.Println(portDB)
	if !er {
		log.Fatal("No env")
	}
	nameDB, er := os.LookupEnv("NAME_DB")
	fmt.Println(nameDB)
	if !er {
		log.Fatal("No env")
	}
	userDB, er := os.LookupEnv("USER_DB")
	fmt.Println(userDB)
	if !er {
		log.Fatal("No env")
	}
	passwordDB, er := os.LookupEnv("PASSWORD_DB")
	fmt.Println(passwordDB)
	if !er {
		log.Fatal("No env")
	}

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", host, portDB, nameDB, userDB, passwordDB)
	db, err := driver.ConnectSQL(dsn)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.SQL.Close()
	log.Println("Connected to database!")

	handlers.NewHandlers(handlers.NewRepository(db))

	return db, nil
}

package di //dependency injection
import (
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/Vishal-2029/BookCrud-App/config"
	_ "github.com/Vishal-2029/BookCrud-App/docs"
	"github.com/Vishal-2029/BookCrud-App/pkg/db"
	"github.com/Vishal-2029/BookCrud-App/pkg/handlers"
	"github.com/Vishal-2029/BookCrud-App/pkg/repo"
	"github.com/Vishal-2029/BookCrud-App/pkg/services"
	"github.com/Vishal-2029/BookCrud-App/utility"
	"github.com/gofiber/fiber/v2"
)

func Init() {
	config.LoadConfig()

	app := fiber.New()

	log := utility.InitLogger()
	log.Println("Starting Book API Server...")

	dbconn := db.Connect()
	log.Println("Database connected successfully.")

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Request: %s %s", c.Method(), c.Path())
		return c.Next()
	})

	BookRepo := repo.NewBookRepo(dbconn)
	BookService := services.NewBookservices(BookRepo)
	handlers.NewBookHandler(app.Group("/"), BookService)

	log.Fatal(app.Listen(":8080"))
}

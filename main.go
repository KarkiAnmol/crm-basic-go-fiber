package main

import (
	"fmt"

	"github.com/KarkiAnmol/crm-basic-go-fiber/database"
	"github.com/KarkiAnmol/crm-basic-go-fiber/lead"
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/lead/id", lead.GetLead)
	app.Get("/lead", lead.GetLeads)
	app.Post("/lead/id", lead.UpdateLead)
	app.Delete("/lead/id", lead.DeleteLead)
}
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		fmt.Println("Error: unable to connect to database")
		return
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}
func main() {
	app := fiber.New()
	initDatabase()
	SetUpRoutes(app)
	app.Listen(":3000")
	defer database.DBConn.Close()
}

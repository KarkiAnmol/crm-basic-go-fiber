package lead

import (
	"log"

	"github.com/KarkiAnmol/crm-basic-go-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
}

func GetLeads(c *fiber.Ctx) {
	var leads []Lead
	db := database.DBConn
	db.Find(&leads)
	c.JSON(leads)

}
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}
func UpdateLead(c *fiber.Ctx) {
	db := database.DBConn
	var lead Lead
	if err := c.BodyParser(lead); err != nil {
		log.Fatal(err)
	}
	db.Create(&lead)
	c.JSON(lead)
}
func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		log.Fatal("No lead found")
	}
	db.Delete(&lead)
	c.Send("Lead deleted Successfully")
}

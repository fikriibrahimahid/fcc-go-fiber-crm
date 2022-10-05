package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(c *fiber.Ctx) {

}

func GetLead(c *fiber.Ctx) {

}

func NewLead(c *fiber.Ctx) {

}

func DeleteLead(c *fiber.Ctx) {

}

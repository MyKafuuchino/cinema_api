package route

import (
	"cinema_api/controller"
	"cinema_api/database"
	"cinema_api/middleware"
	"cinema_api/repository"
	"cinema_api/service"
	"github.com/gofiber/fiber/v2"
)

func SetupTicketRouter(r fiber.Router) {
	db := database.Db
	ticketRepo := repository.NewTicketRepository(db)
	userRepo := repository.NewUserRepository(db)
	screeningRepo := repository.NewScreeningRepository(db)

	ticketService := service.NewTicketService(ticketRepo, userRepo, screeningRepo)
	ticketController := controller.NewTicketController(ticketService)

	ticketRouter := r.Group("/tickets")
	ticketRouter.Get("", ticketController.GetAllTickets)
	ticketRouter.Get("/:id", ticketController.GetTicketById)
	ticketRouter.Post("", middleware.ProtectRouteByRole("ADMIN"), ticketController.CreateTicket)
	ticketRouter.Put("/:id", middleware.ProtectRouteByRole("ADMIN"), ticketController.UpdateTicketById)
	ticketRouter.Delete("/:id", middleware.ProtectRouteByRole("ADMIN"), ticketController.DeleteTicket)

}

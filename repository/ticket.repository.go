package repository

import (
	"cinema_api/model"
	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *model.Ticket) error
	FindAll() ([]model.Ticket, error)
	FindById(id uint) (*model.Ticket, error)

	FindByUserId(id uint) ([]model.Ticket, error)
	FindByScreeningId(id uint) ([]model.Ticket, error)
	Update(ticket *model.Ticket) error
	Delete(ticket *model.Ticket) error
}

type ticketRepository struct {
	db *gorm.DB
}

func (r ticketRepository) Create(ticket *model.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r ticketRepository) FindAll() ([]model.Ticket, error) {
	var tickets []model.Ticket
	if err := r.db.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r ticketRepository) FindById(id uint) (*model.Ticket, error) {
	var ticket model.Ticket
	if err := r.db.Where("id = ?", id).First(&ticket).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r ticketRepository) FindByUserId(id uint) ([]model.Ticket, error) {
	var ticket []model.Ticket
	if err := r.db.Where("user_id = ?", id).Find(&ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}

func (r ticketRepository) FindByScreeningId(id uint) ([]model.Ticket, error) {
	var ticket []model.Ticket
	if err := r.db.Where("screening_id = ?", id).Find(&ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}

func (r ticketRepository) Update(ticket *model.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r ticketRepository) Delete(ticket *model.Ticket) error {
	return r.db.Delete(ticket).Error
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db: db}
}

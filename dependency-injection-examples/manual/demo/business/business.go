package business

import (
	"demo/database"
)

// BusinessLogic interface
type BusinessLogic interface {
	ProcessData() string
}

// Business struct
type Business struct {
	db database.DatabaseAccess
}

// Constructor
func NewBusiness(db database.DatabaseAccess) *Business {
	return &Business{db: db}
}

// Implement ProcessData()
func (b Business) ProcessData() string {
	return "Business logic processed " + b.db.GetData()
}

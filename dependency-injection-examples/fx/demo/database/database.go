package database

// DatabaseAccess interface
type DatabaseAccess interface {
	GetData() string
}

// Database struct
type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

// Implement GetData()
func (db Database) GetData() string {
	return "Data from database"
}

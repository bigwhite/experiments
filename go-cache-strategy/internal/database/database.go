package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	// Needed for LogEntry
	"cachestrategysdemo/internal/models" // Adjust import path based on your module name

	_ "github.com/mattn/go-sqlite3"
)

// Database defines the interface for database operations.
type Database interface {
	GetUser(ctx context.Context, userID string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	BulkUpdateUsers(ctx context.Context, users []*models.User) error // For Write-Behind
	InsertLogEntry(ctx context.Context, entry *models.LogEntry) error
	GetLogByID(ctx context.Context, logID string) (*models.LogEntry, error)
	Close() error
}

// --- SQLite Database Implementation ---
type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB(filepath string) (Database, error) {
	// Use WAL mode for better concurrency, timeout for busy handling
	db, err := sql.Open("sqlite3", filepath+"?_journal_mode=WAL&_busy_timeout=5000")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Ping to ensure connection is alive
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Create users table
	createUsersTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        name TEXT
    );`
	if _, err = db.Exec(createUsersTableSQL); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	// Create logs table
	createLogsTableSQL := `CREATE TABLE IF NOT EXISTS logs (
		id TEXT PRIMARY KEY,
		timestamp DATETIME,
		level TEXT,
		message TEXT
	);`
	if _, err = db.Exec(createLogsTableSQL); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create logs table: %w", err)
	}

	log.Println("Database initialized and tables checked/created.")
	return &SQLiteDB{db: db}, nil
}

func (s *SQLiteDB) GetUser(ctx context.Context, userID string) (*models.User, error) {
	query := `SELECT id, name FROM users WHERE id = ?`
	row := s.db.QueryRowContext(ctx, query, userID)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Not found
		}
		// log.Printf("[DB] Error getting user %s: %v\n", userID, err)
		return nil, fmt.Errorf("db query error: %w", err)
	}
	return user, nil
}

func (s *SQLiteDB) UpdateUser(ctx context.Context, user *models.User) error {
	query := `INSERT OR REPLACE INTO users (id, name) VALUES (?, ?)`
	_, err := s.db.ExecContext(ctx, query, user.ID, user.Name)
	if err != nil {
		// log.Printf("[DB] Error updating user %s: %v\n", user.ID, err)
		return fmt.Errorf("db update error: %w", err)
	}
	return nil
}

func (s *SQLiteDB) BulkUpdateUsers(ctx context.Context, users []*models.User) error {
	if len(users) == 0 {
		return nil
	}
	ids := make([]string, len(users))
	for i, u := range users {
		ids[i] = u.ID
	}
	log.Printf("==> [DB] Bulk updating %d users (IDs: %v)...\n", len(users), ids)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("==> [DB] Error starting transaction for bulk update: %v\n", err)
		return fmt.Errorf("db begin tx error: %w", err)
	}
	// Ensure rollback happens if commit fails or isn't reached
	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	query := `INSERT OR REPLACE INTO users (id, name) VALUES (?, ?)`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("==> [DB] Error preparing statement for bulk update: %v\n", err)
		return fmt.Errorf("db prepare error: %w", err)
	}
	defer stmt.Close()

	successCount := 0
	for _, user := range users {
		_, err = stmt.ExecContext(ctx, user.ID, user.Name)
		if err != nil {
			log.Printf("==> [DB] Error updating user %s in bulk: %v\n", user.ID, err)
			// Decide: continue or abort? Continuing for demo simplicity.
		} else {
			successCount++
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("==> [DB] Error committing transaction for bulk update: %v\n", err)
		return fmt.Errorf("db commit error: %w", err)
	}
	committed = true // Mark as committed to prevent deferred rollback
	log.Printf("==> [DB] Successfully bulk updated %d/%d users for IDs: %v\n", successCount, len(users), ids)
	return nil
}

func (s *SQLiteDB) InsertLogEntry(ctx context.Context, entry *models.LogEntry) error {
	// log.Printf("[DB] Inserting log directly (ID: %s)\n", entry.ID)
	query := `INSERT INTO logs (id, timestamp, level, message) VALUES (?, ?, ?, ?)`
	_, err := s.db.ExecContext(ctx, query, entry.ID, entry.Timestamp, entry.Level, entry.Message)
	if err != nil {
		// log.Printf("[DB] Error inserting log %s: %v\n", entry.ID, err)
		return fmt.Errorf("db insert log error: %w", err)
	}
	return nil
}

func (s *SQLiteDB) GetLogByID(ctx context.Context, logID string) (*models.LogEntry, error) {
	query := `SELECT id, timestamp, level, message FROM logs WHERE id = ?`
	row := s.db.QueryRowContext(ctx, query, logID)
	entry := &models.LogEntry{}
	err := row.Scan(&entry.ID, &entry.Timestamp, &entry.Level, &entry.Message)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Not found
		}
		// log.Printf("[DB] Error getting log %s: %v\n", logID, err)
		return nil, fmt.Errorf("db query log error: %w", err)
	}
	return entry, nil
}

func (s *SQLiteDB) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

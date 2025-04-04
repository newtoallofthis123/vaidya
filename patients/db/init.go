package db

import (
	"database/sql"

	"github.com/Masterminds/squirrel"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
	pq squirrel.StatementBuilderType
}

func NewStore(connPath string) (Store, error) {
	db, err := sql.Open("postgres", connPath)
	if err != nil {
		return Store{}, err
	}

	return Store{db: db, pq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)}, nil
}

func (s *Store) InitTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS patients(
		id text primary key,
		name text not null,
		age int not null,
		gender text not null,
		address text,
		identity text,
		phone text,
		conditions text,
		problems text,
		description text not null,
		medicines text DEFAULT '',
		diagnosis text DEFAULT '',
		next_session text DEFAULT '',
		created_at timestamp default now()
	);
	`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

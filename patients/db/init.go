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
		first_name text not null,
		last_name text,
		age int not null,
		gender text not null,
		address text,
		identity text,
		phone text,
		email text,
		description text not null,
		recurring boolean default false,
		created_at timestamp default now()
	);

	CREATE TABLE IF NOT EXISTS sessions(
		id text primary key,
		doctor_id text REFERENCES doctors(id),
		curr_time timestamp not null,
		recommended boolean default false
	);

	CREATE TABLE IF NOT EXISTS records(
		id text primary key,
		patient_id text REFERENCES patients(id),
		session_id text REFERENCES sessions(id),
		problems text not null,
		medicines text not null,
		conditions text not null,
		diagnosis text not null,
		next_session text REFERENCES sessions(id) DEFAULT null
	);
	`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

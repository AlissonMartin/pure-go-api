package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password= sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
		id serial primary key,
		name varchar(255),
		balance serial 
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := "insert into accounts (name, balance) values ($1, $2)"
	_, err := s.db.Query(query, acc.Name, acc.Balance)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err :=s.db.Query("select * from accounts")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)

		err := rows.Scan(&account.ID, &account.Name, &account.Balance)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	return accounts, nil	
}


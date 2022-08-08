package repository

import (
	"database/sql"
	"errors"
	"time"
)

// SQLiteRepository the type for a repository that connects to sqlite database
type SQLiteRepository struct {
	Conn *sql.DB
}

// NewSQLiteRepository returns a new repository with a connection to sqlite
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

// Migrate creates the table(s) we need
func (repo *SQLiteRepository) Migrate() error {
	query := `
	create table if not exists holdings(
		id integer primary key autoincrement,
		amount real not null,
		purchase_date integer not null,
		purchase_price integer not null);
	`

	_, err := repo.Conn.Exec(query)
	return err
}

// InsertHolding inserts one record into the database
func (repo *SQLiteRepository) InsertHolding(holdings Holdings) (*Holdings, error) {
	stmt := "insert into holdings (amount, purchase_date, purchase_price) values (?, ?, ?)"

	res, err := repo.Conn.Exec(stmt, holdings.Amount, holdings.PurchaseDate.Unix(), holdings.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	holdings.ID = id
	return &holdings, nil
}

// AllHoldings returns all holdings, by purchase date
func (repo *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := "select id, amount, purchase_date, purchase_price from holdings order by purchase_date"
	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Holdings
	for rows.Next() {
		var h Holdings
		var unixTime int64
		err := rows.Scan(
			&h.ID,
			&h.Amount,
			&unixTime,
			&h.PurchasePrice,
		)
		if err != nil {
			return nil, err
		}
		h.PurchaseDate = time.Unix(unixTime, 0)
		all = append(all, h)
	}

	return all, nil
}

func (repo *SQLiteRepository) GetHoldingByID(id int) (*Holdings, error) {
	row := repo.Conn.QueryRow("select id, amount, purchase_date, purchase_price from holdings where id = ?", id)

	var h Holdings
	var unixTime int64
	err := row.Scan(
		&h.ID,
		&h.Amount,
		&unixTime,
		&h.PurchasePrice,
	)

	if err != nil {
		return nil, err
	}

	h.PurchaseDate = time.Unix(unixTime, 0)

	return &h, nil
}

func (repo *SQLiteRepository) UpdateHolding(id int64, updated Holdings) error {
	if id == 0 {
		return errors.New("invalid updated id")
	}

	stmt := "Update holdings set amount = ?, purchase_date = ?, purchase_price = ? where id = ?"
	res, err := repo.Conn.Exec(stmt, updated.Amount, updated.PurchaseDate.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

func (repo *SQLiteRepository) DeleteHolding(id int64) error {
	res, err := repo.Conn.Exec("delete from holdings where id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errDeleteFailed
	}
	
	return nil
}
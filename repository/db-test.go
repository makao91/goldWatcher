package repository

import "time"

type TestRepository struct {}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

// Migrate creates the table(s) we need
func (repo *TestRepository) Migrate() error {
	return nil
}

// InsertHolding inserts one record into the database
func (repo *TestRepository) InsertHolding(holdings Holdings) (*Holdings, error) {
	return &holdings, nil
}

// AllHoldings returns all holdings, by purchase date
func (repo *TestRepository) AllHoldings() ([]Holdings, error) {
	var all []Holdings
	h := Holdings{
		Amount: 1,
		PurchaseDate: time.Now(),
		PurchasePrice: 1000,
	}
	all = append(all, h)

	h = Holdings{
		Amount: 2,
		PurchaseDate: time.Now(),
		PurchasePrice: 2000,
	}

	all = append(all, h)
	
	return all, nil
}

func (repo *TestRepository) GetHoldingByID(id int) (*Holdings, error) {
	h := Holdings{
		Amount: 1,
		PurchaseDate: time.Now(),
		PurchasePrice: 1000,
	}

	return &h, nil
}

func (repo *TestRepository) UpdateHolding(id int64, updated Holdings) error {
	return nil
}

func (repo *TestRepository) DeleteHolding(id int64) error {
	return nil
}
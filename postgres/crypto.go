package postgres

import (
	"database/sql"

	"github.com/caiomp87/crypto-votes-entities/models"
)

type ICrypto interface {
	FindAll() ([]*models.Crypto, error)
	FindByID(id int) (*models.Crypto, error)
	Create(cypto *models.Crypto) (*models.Crypto, error)
	UpdateByID(id int, crypto *models.Crypto) (*models.Crypto, error)
	DeleteByID(id int) error
}

type cryptoService struct {
	db *sql.DB
}

func NewCryptoService() ICrypto {
	return &cryptoService{}
}

func (cs *cryptoService) FindAll() ([]*models.Crypto, error) {
	rows, err := cs.db.Query("SELECT * FROM cryptos")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cryptos := make([]*models.Crypto, 0)
	for rows.Next() {
		var crypto *models.Crypto
		if err := rows.Scan(&crypto); err != nil {
			return nil, err
		}
		cryptos = append(cryptos, crypto)
	}

	return cryptos, nil
}

func (cs *cryptoService) FindByID(id int) (*models.Crypto, error) {
	row := cs.db.QueryRow(`SELECT * FROM cryptos WHERE id=$1`, id)

	var crypto *models.Crypto
	if err := row.Scan(&crypto); err != nil {
		return nil, err
	}

	return crypto, nil
}

func (cs *cryptoService) Create(crypto *models.Crypto) (*models.Crypto, error) {
	insertStatement := `INSERT INTO cryptos ("Name", "Network") VALUES($1, $2)`

	rows, err := cs.db.Query(insertStatement, crypto.Name, crypto.Network)
	if err != nil {
		return nil, err
	}

	var createdCrypto *models.Crypto
	if err := rows.Scan(&createdCrypto); err != nil {
		return nil, err
	}

	return createdCrypto, nil
}

func (cs *cryptoService) UpdateByID(id int, crypto *models.Crypto) (*models.Crypto, error) {
	return nil, nil
}

func (cs *cryptoService) DeleteByID(id int) error {
	return nil
}

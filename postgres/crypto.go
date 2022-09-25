package postgres

import (
	"database/sql"

	"github.com/caiomp87/crypto-votes-entities/models"
)

var CryptoDatastore ICrypto

type ICrypto interface {
	FindAll() ([]*models.Crypto, error)
	FindByID(id int) (*models.Crypto, error)
	Create(crypto *models.Crypto) error
	UpdateByID(id int, crypto *models.Crypto) (*models.Crypto, error)
	DeleteByID(id int) error
}

type cryptoService struct {
	db *sql.DB
}

func NewCryptoService(db *sql.DB) ICrypto {
	return &cryptoService{
		db,
	}
}

func (cs *cryptoService) FindAll() ([]*models.Crypto, error) {
	rows, err := cs.db.Query("SELECT * FROM cryptos")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cryptos := make([]*models.Crypto, 0)
	for rows.Next() {
		crypto := models.Crypto{}
		if err := rows.Scan(&crypto.ID, &crypto.Name, &crypto.Network, &crypto.UpVotes, &crypto.DownVotes); err != nil {
			return nil, err
		}
		cryptos = append(cryptos, &crypto)
	}

	return cryptos, nil
}

func (cs *cryptoService) FindByID(id int) (*models.Crypto, error) {
	row := cs.db.QueryRow(`SELECT * FROM cryptos WHERE id=$1`, id)

	crypto := models.Crypto{}
	if err := row.Scan(&crypto.ID, &crypto.Name, &crypto.Network, &crypto.UpVotes, &crypto.DownVotes); err != nil {
		return nil, err
	}

	return &crypto, nil
}

func (cs *cryptoService) Create(crypto *models.Crypto) error {
	insertStatement := `INSERT INTO cryptos ("name", "network") VALUES($1, $2)`

	_, err := cs.db.Exec(insertStatement, crypto.Name, crypto.Network)
	if err != nil {
		return err
	}

	return nil
}

func (cs *cryptoService) UpdateByID(id int, crypto *models.Crypto) (*models.Crypto, error) {
	return nil, nil
}

func (cs *cryptoService) DeleteByID(id int) error {
	return nil
}

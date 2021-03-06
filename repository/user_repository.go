package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/juanjoss/off-users-service/ports"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository() *UserRepository {
	host := os.Getenv("DB_HOST")
	driver := os.Getenv("DB_DRIVER")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	source := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, dbPort, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Connect(driver, source)
	if err != nil {
		log.Fatalf("unable to connect to DB: %v", err.Error())
	}

	repo := &UserRepository{
		db: db,
	}

	return repo
}

/*
	Inserts the new user and its ssds.
*/
func (ur *UserRepository) Register(request ports.RegisterRequest) error {
	var id int

	row := ur.db.QueryRow(
		`INSERT INTO users (first_name, last_name, email, password) 
		VALUES ($1, $2, $3, $4)
		RETURNING id`,
		request.User.FirstName,
		request.User.LastName,
		request.User.Email,
		request.User.Password,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}

	for _, ssd := range request.SSDs {
		ssd.UserId = id
		_, err := ur.db.NamedExec(
			`INSERT INTO ssds (user_id, mac_address)
			VALUES (:user_id, :mac_address)`,
			ssd,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
	Inserts a product and its quantity into a ssd. Upserts for existing records.
*/
func (ur *UserRepository) AddProductToSSD(request ports.AddProductToSsdRequest) error {
	_, err := ur.db.Exec(
		`INSERT INTO product_ssds (ssd_id, barcode, quantity) 
		VALUES ($1, $2, $3)
		ON CONFLICT DO NOTHING`,
		request.SsdId, request.Barcode, request.Quantity,
	)
	if err != nil {
		return err
	}

	return nil
}

/*
	Returns a random ssd.
*/
func (ur *UserRepository) RandomSSD() (ports.GetRandomSsdResponse, error) {
	response := ports.GetRandomSsdResponse{}

	err := ur.db.Get(&response, `SELECT * FROM ssds ORDER BY RANDOM() LIMIT 1`)
	if err != nil {
		return response, err
	}

	return response, nil
}

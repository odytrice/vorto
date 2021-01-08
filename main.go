package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// Delivery data transfer object
type Delivery struct {
	ID         int       `json:"id"`
	SupplierID int       `json:"supplierId"`
	DriverID   int       `json:"driverId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func getConnectionString() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "<secretpassword>")
	dbname := getEnv("DB_NAME", "vorto")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func fetchInvalidDeliveries() ([]Delivery, error) {

	deliveries := make([]Delivery, 0)

	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		msg := "Failed to open a DB connection: " + err.Error()
		return deliveries, errors.New(msg)
	}
	defer db.Close()

	query := `
		SELECT DISTINCT d.* from public.delivery d
		JOIN public.supplier_bean_type sbt
		ON d.supplier_id = sbt.supplier_id
		WHERE sbt.bean_type_id
		NOT IN (
			SELECT bean_type_id FROM carrier_bean_type cbt
			JOIN public.driver dr
			ON cbt.carrier_id = dr.carrier_id
			WHERE dr.Id = d.driver_id
		)
	`

	rows, err := db.Query(query)
	if err != nil {
		return deliveries, err
	}
	defer rows.Close()

	for rows.Next() {
		var delivery Delivery
		rows.Scan(&delivery.ID, &delivery.SupplierID, &delivery.DriverID, &delivery.CreatedAt, &delivery.UpdatedAt)
		deliveries = append(deliveries, delivery)
	}
	return deliveries, nil
}

// Fetches Invalid Deliveries
func getInvalid(w http.ResponseWriter, r *http.Request) {
	deliveries, err := fetchInvalidDeliveries()
	if err != nil {
		//TODO
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	jsonBytes, err := json.Marshal(deliveries)
	if err != nil {
		//TODO
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func main() {
	http.HandleFunc("/deliveries/invalid", getInvalid)

	fmt.Println("Server is running on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

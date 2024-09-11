package repositories

import (
	"scoter-web-api/internal/config"
	"scoter-web-api/internal/models"
)

func GetAllScooters() ([]models.Scooter, error) {
	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, latitude, longitude, is_active FROM scooters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scooters []models.Scooter
	for rows.Next() {
		var scooter models.Scooter
		if err := rows.Scan(&scooter.ID, &scooter.Name, &scooter.Latitude, &scooter.Longitude, &scooter.IsActive); err != nil {
			return nil, err
		}
		scooters = append(scooters, scooter)
	}
	return scooters, nil
}

func CreateScooter(scooter *models.Scooter) error {
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO scooters (name, latitude, longitude, is_active) VALUES (?, ?, ?, ?)",
		scooter.Name, scooter.Latitude, scooter.Longitude, scooter.IsActive)
	return err
}

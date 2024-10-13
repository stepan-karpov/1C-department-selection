package db

import (
	"database/sql"
	"fmt"
	logs "main/server/logs_writer"

	_ "github.com/mattn/go-sqlite3"
)

func SubscribeUser(dbPath string, ip string, experiment_number int) error {
	logs.LogInfo(fmt.Sprintf("Adding user with IP: %s", ip))

	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error while connecting to database: %w", err)
	}
	defer database.Close()

	var count int
	checkUserSQL := `
	SELECT COUNT(*) FROM subscriptions WHERE ip = ? AND experiment_number = ?;`

	err = database.QueryRow(checkUserSQL, ip, experiment_number).Scan(&count)
	if err != nil {
		return fmt.Errorf("error while checking user existence: %w", err)
	}

	if count > 0 {
		logs.LogError(fmt.Sprintf("User with IP: %s and experiment_number: %d already exists", ip, experiment_number))
		return fmt.Errorf("user with IP: %s and experiment_number: %d already exists", ip, experiment_number)
	}

	insertUserSQL := `
	INSERT INTO subscriptions (ip, experiment_number) 
	VALUES (?, ?);`

	_, err = database.Exec(insertUserSQL, ip, experiment_number)
	if err != nil {
		return fmt.Errorf("error while inserting user: %w", err)
	}

	logs.LogInfo(fmt.Sprintf("User with IP: %s added successfully", ip))
	return nil
}

type Subscription struct {
 ID               int
 IP               string
 ExperimentNumber int
}

func GetAllSubscriptions(dbPath string) ([]Subscription, error) {
 database, err := sql.Open("sqlite3", dbPath)
 if err != nil {
  return nil, fmt.Errorf("error while connecting to database: %w", err)
 }
 defer database.Close()

 rows, err := database.Query("SELECT id, ip, experiment_number FROM subscriptions;")
 if err != nil {
  return nil, fmt.Errorf("error while retrieving subscriptions: %w", err)
 }
 defer rows.Close()

 var subscriptions []Subscription

 for rows.Next() {
  var sub Subscription
  if err := rows.Scan(&sub.ID, &sub.IP, &sub.ExperimentNumber); err != nil {
   return nil, fmt.Errorf("error while scanning subscriptions: %w", err)
  }
  subscriptions = append(subscriptions, sub)
 }

 if err := rows.Err(); err != nil {
  return nil, fmt.Errorf("error occurred during row iteration: %w", err)
 }

 return subscriptions, nil
}

package database

import (
	"database/sql"
    "log"

		_ "modernc.org/sqlite"
)

var db *sql.DB

// Ініціалізація бази даних
func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./telegram_bot.db")
	if err != nil {
		return err
	}

	// Створюємо таблицю для повідомлень
	createTableSQL := `CREATE TABLE IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        message TEXT
    );`
		_, err = db.Exec(createTableSQL)
		if err != nil {
			return err
		}

		log.Println("База даних успішно ініціалізована.")
		return nil
}

// Збереження повідомлення
func SaveMessage(userID int64, message string) error {
	query := `INSERT INTO messages (user_id, message) VALUES (?, ?)`
	_, err := db.Exec(query, userID, message)
	return err
}

// Отримання всіх повідомлень користувача
func GetMessages(userID int64) ([]string, error) {
	query := `SELECT message FROM messages WHERE user_id = ?`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var message string
		err = rows.Scan(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
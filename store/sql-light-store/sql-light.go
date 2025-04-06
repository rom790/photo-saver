package sqllightstore

import (
	"database/sql"
	// "fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SqlLightStore struct {
	db     *sql.DB
	dbPath string
}

func New(dbPath string) *SqlLightStore {
	return &SqlLightStore{
		dbPath: dbPath,
	}
}

func (s *SqlLightStore) Open() error {
	db, err := sql.Open("sqlite3", s.dbPath)
	// fmt.Println(s.dbPath, "fdsfas")
	if err != nil {
		return err
	}
	s.db = db
	return s.init()
}

func (s *SqlLightStore) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// func (s *SqlLightStore) init() error {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		username TEXT NOT NULL UNIQUE
// 	);
// 	CREATE TABLE IF NOT EXISTS photos (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		user_id INTEGER,
// 		file_path TEXT NOT NULL,
// 		FOREIGN KEY(user_id) REFERENCES users(id)
// 	);
// 	`

// 	// fmt.Println("test")
// 	_, err := s.db.Exec(query)
// 	// fmt.Println(err)

//		s.CheckSQLite()
//		return err
//	}
func (s *SqlLightStore) init() error {
	// Создание таблиц
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE
    );
    CREATE TABLE IF NOT EXISTS photos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        file_path TEXT NOT NULL,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	// // Вставка тестовых данных
	// _, err = s.db.Exec(`
	//     INSERT OR IGNORE INTO users (username) VALUES ('test_user');
	//     INSERT OR IGNORE INTO photos (user_id, file_path)
	//     VALUES (1, '/photos/test.jpg');
	// `)
	// if err != nil {
	// 	return fmt.Errorf("ошибка вставки тестовых данных: %v", err)
	// }

	// return s.CheckSQLite()
	return nil
}

// func (s *SqlLightStore) CheckSQLite() error {
// 	var result string
// 	query := "SELECT sqlite_version();"
// 	err := s.db.QueryRow(query).Scan(&result)
// 	if err != nil {
// 		return fmt.Errorf("не удалось выполнить запрос: %v", err)
// 	}
// 	fmt.Printf("Версия SQLite: %s\n", result)
// 	return nil
// }

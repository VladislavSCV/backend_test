import (
    "database/sql"
    _ "github.com/lib/pq" // Драйвер для PostgreSQL
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
    connStr := "user=youruser dbname=yourdb sslmode=disable password=yourpassword"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    DB = db
    return db, nil
}
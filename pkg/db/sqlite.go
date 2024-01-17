package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Sqlite3Client *sql.DB
)

func InitSqlite3() {
	db, err := sql.Open("sqlite3", "./cloud-db")
	if err != nil {
		panic(err)
	}
	Sqlite3Client = db
}

func NewSqlite3() Sqlite3 {
	if Sqlite3Client == nil {
		InitSqlite3()
	}
	return Sqlite3{
		DB: Sqlite3Client,
	}
}

type Sqlite3 struct {
	DB *sql.DB
}

func (db Sqlite3) SelectPage(ctx context.Context, filter interface{}, sort interface{}, skip, limit int64) (int64, interface{}, error) {
	return 0, nil, nil
}
func (db Sqlite3) SelectList(ctx context.Context, filter interface{}, sort interface{}) (interface{}, error) {
	return nil, nil
}
func (db Sqlite3) SelectOne(ctx context.Context, filter interface{}) (interface{}, error) {
	return 0, nil
}
func (db Sqlite3) SelectCount(ctx context.Context, filter interface{}) (int64, error) { return 0, nil }
func (db Sqlite3) UpdateOne(ctx context.Context, filter, update interface{}) (int64, error) {
	return 0, nil
}
func (db Sqlite3) UpdateMany(ctx context.Context, filter, update interface{}) (int64, error) {
	return 0, nil
}
func (db Sqlite3) Delete(ctx context.Context, filter interface{}) (int64, error) { return 0, nil }
func (db Sqlite3) InsertOne(ctx context.Context, model interface{}) (interface{}, error) {
	return nil, nil
}
func (db Sqlite3) InsertMany(ctx context.Context, models []interface{}) ([]interface{}, error) {
	return nil, nil
}
func (db Sqlite3) Aggregate(ctx context.Context, pipeline interface{}, result interface{}) error {
	return nil
}
func (db Sqlite3) CreateIndexes(ctx context.Context, indexes []mongo.IndexModel) error {
	return nil
}
func (db Sqlite3) GetCollection() *mongo.Collection {
	return nil
}

package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"hello/config"
	"log"

	"github.com/google/uuid"
	// ドライバのインストール
	_ "github.com/mattn/go-sqlite3"
)

// テーブルの作成
var Db *sql.DB

// エラーの宣言
var err error

// テーブル名の宣言
const (
	tableNameUser = "users"
	// todoテーブル追加
	tableNameTodo = "todos"
	// sessionテーブル作成
	tableNameSession = "sessions"
)

// テーブルはmain関数の前に作成
func init() {
	// データーベースとエラー。ドライバとデーターベース名
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}

	// コマンドの作成
	// 最後にテーブル名を渡す
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	// コマンドを呼び出し
	Db.Exec(cmdU)

	// コマンドを作成
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	// コマンド呼び出し
	Db.Exec(cmdT)

	// コマンド作成
	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	// コマンド実行
	Db.Exec(cmdS)
}

// UUID作成
// 返り値をuuidobj
func createUUID() (uuidobj uuid.UUID) {
	// NewUUIDを使用
	uuidobj, _ = uuid.NewUUID()
	// returnで返す
	return uuidobj
}

// パスワード作成
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	// returnで返す
	return cryptext
}

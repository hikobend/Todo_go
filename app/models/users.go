package models

import (
	"log"
	"time"
)

// struct作成
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

// Userの作成
// User型のメソッドとして作成
func (u *User) CreateUser() (err error) {
	// コマンド作成
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	// コマンドを実行
	// UUIDとPassWordを作成する必要がある
	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Userの取得
// 引数と返り値
func GetUser(id int) (user User, err error) {
	// userの定義
	user = User{}
	// コマンドの作成
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`

	// コマンドの実行
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	// user, errを返す
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

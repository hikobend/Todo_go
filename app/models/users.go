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

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
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

// 関数作成
func (u *User) UpdateUser() (err error) {
	// コマンド作成
	cmd := `update users set name = ?, email = ? where id = ?`
	// コマンド実行
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// 関数の作成
func (u *User) DeleteUser() (err error) {
	// コマンドの作成
	cmd := `delete from users where id = ?`
	// コマンドの実行
	// idが一致するものを削除する
	_, err = Db.Exec(cmd, u.ID)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at) values (?, ?, ?, ?)`
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)
	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err

}

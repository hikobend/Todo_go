package models

import (
	"log"
	"time"
)

// struct作成
type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// Userユーザーのメソッドとして作成
func (u *User) CreateTodo(content string) (err error) {
	// コマンド作成
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`
	// コマンド実行
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// 関数を作成
func GetTodo(id int) (todo Todo, err error) {
	// コマンド作成
	cmd := `select id, content, user_id, created_at from todos where id = ?`
	// todo宣言
	todo = Todo{}
	// コマンド実行
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	// todo errを返す
	return todo, err
}

func GetTodos() (todos []Todo, srr error) {
	cmd := `select id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodoByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

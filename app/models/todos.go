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

// 複数の関数を取得
func GetTodos() (todos []Todo, err error) {
	// コマンド作成
	cmd := `select id, content, user_id, created_at from todos`
	// コマンドを渡す
	rows, err := Db.Query(cmd)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	// それぞれを取り出す
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		// エラーハンドリング
		if err != nil {
			log.Fatalln(err)
		}
		// append
		todos = append(todos, todo)
	}
	rows.Close()

	// todos errを返す
	return todos, err
}

// 関数の作成
func (u *User) GetTodoByUser() (todos []Todo, err error) {
	// コマンドを作成。すべてのtodosを取得
	cmd := `select id, content, user_id, created_at from todos where user_id = ?`
	// コマンドを実行
	rows, err := Db.Query(cmd, u.ID)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	//Scan
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		// エラーハンドリング
		if err != nil {
			log.Fatalln(err)
		}
		// todosリストに追加する
		todos = append(todos, todo)
	}
	// 閉じる
	rows.Close()

	// todos errを返す
	return todos, err
}

// 関数の作成。Todoメソッド
func (t *Todo) UpdateTodo() error {
	// コマンドの作成
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	// コマンドの実行
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// 関数の作成。Todoメソッド
func (t *Todo) DeleteTodo() error {
	// コマンド作成
	cmd := `delete from todos where id = ?`
	// コマンドを実行
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	// return err
	return err
}

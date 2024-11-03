package task

import (
	"database/sql"
	"log"

	"github.com/yann-fk-21/todo-app/types"
)

type Store struct {
	db *sql.DB
	logger *log.Logger
}

func NewStore(db *sql.DB, logger *log.Logger) *Store{
	return &Store{
		db: db,
		logger: logger,
	}
}

func (st *Store) CreateTask(task types.Task) error {
	_, err := st.db.Query("INSERT INTO task(title, description, status, createdAt) VALUES (?, ?, ?, ?)", task.Title, task.Description, task.Status, task.CreatedAt)
    if err != nil {
		return err
	}
	return nil
}

func (st *Store) GetTasks() ([]types.Task, error) {
	rows, err := st.db.Query("SELECT * FROM task")
	if err != nil {
		st.logger.Println(err)
		st.logger.Fatal(err)
	}

	var tasks []types.Task

	defer rows.Close()

	for rows.Next() {
		task, err := scanIntoTask(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *task)
	}
	return tasks, nil
}

func scanIntoTask(rows *sql.Rows)(*types.Task, error) {
	 task := new(types.Task)
     err := rows.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return task, nil
}
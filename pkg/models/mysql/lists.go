package mysql

import (
	"database/sql"
	"errors"

	"github.com/cristiano-pacheco/list/pkg/models"
)

// ListModel type which wraps a sql.DB connection pool.
type ListModel struct {
	DB *sql.DB
}

// Insert a new list item
func (model *ListModel) Insert(name string) (int, error) {
	stmt := `INSERT INTO list (name) VALUES(?)`

	result, err := model.DB.Exec(stmt, name)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get select one list record from database
func (model *ListModel) Get(id int) (*models.List, error) {
	stmt := `SELECT * FROM list WHERE id = ?`

	row := model.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed List struct.
	list := &models.List{}

	err := row.Scan(&list.ID, &list.Name, &list.CreatedAt, &list.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return list, nil
}

// GetAll get all records from database
func (model *ListModel) GetAll() ([]*models.List, error) {
	stmt := `SELECT * FROM list`

	rows, err := model.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Initialize a pointer to a new zeroed List struct.
	lists := []*models.List{}

	for rows.Next() {
		list := &models.List{}

		err := rows.Scan(&list.ID, &list.Name, &list.CreatedAt, &list.UpdatedAt)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lists, nil
}

// Update the record in database
func (model *ListModel) Update(id int, name string) (bool, error) {
	stmt := `UPDATE list SET name = ? WHERE id = ?`

	_, err := model.DB.Exec(stmt, name, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, models.ErrNoRecord
		}

		return false, err
	}

	return true, nil
}

// Delete the record in database
func (model *ListModel) Delete(id int) (bool, error) {
	stmt := `DELETE FROM list WHERE id = ?`

	_, err := model.DB.Exec(stmt, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, models.ErrNoRecord
		}
		return false, err
	}

	return true, nil
}

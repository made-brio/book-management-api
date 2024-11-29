package repository

import (
	"book-management-api/internal/models"
	"database/sql"
)

func GetAllCategories(db *sql.DB) (result []models.Category, err error) {
	sql := "SELECT * FROM categories"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var category models.Category

		err = rows.Scan(&category.ID,
			&category.Name)
		if err != nil {
			return
		}

		result = append(result, category)
	}

	return
}

func CreateCategory(db *sql.DB, payload models.CreateCategoryRequest) (err error) {
	sql := "INSERT INTO categories name VALUES $1"
	_, err = db.Exec(sql, payload.Name)
	return
}

func GetCategoryByID(db *sql.DB, id int) (result models.Category, err error) {
	sqlQuery := "SELECT id, name FROM categories WHERE id = $1"

	// Eksekusi query dengan parameter id
	err = db.QueryRow(sqlQuery, id).Scan(&result.ID,
		&result.Name)

	// Tangani jika tidak ada data ditemukan
	if err == sql.ErrNoRows {
		return result, nil // Kembalikan struct kosong tanpa error
	} else if err != nil {
		return result, err // error lain jika terjadi
	}

	return result, nil
}

func DeleteCategory(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM categories WHERE id = $1"
	_, err = db.Exec(sql, id)
	return
}

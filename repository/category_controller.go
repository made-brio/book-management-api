package repository

import (
	"book-management-api/models"
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
			&category.Name,
			&category.CreatedAt,
			&category.CreatedBy,
			&category.ModifiedAt,
			&category.ModifiedBy)
		if err != nil {
			return
		}

		result = append(result, category)
	}

	return
}

func CreateCategory(db *sql.DB, category models.Category) (err error) {
	sql := "INSERT INTO categories (id, name) VALUES ($1, $2)"
	_, err = db.Exec(sql, category.ID, category.Name)
	return
}

func GetCategoryByID(db *sql.DB, id int) (result models.Category, err error) {
	sqlQuery := "SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id = $1"

	// Eksekusi query dengan parameter id
	err = db.QueryRow(sqlQuery, id).Scan(&result.ID,
		&result.Name,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.ModifiedAt,
		&result.ModifiedBy)

	// Tangani jika tidak ada data ditemukan
	if err == sql.ErrNoRows {
		return result, nil // Kembalikan struct kosong tanpa error
	} else if err != nil {
		return result, err // error lain jika terjadi
	}

	return result, nil
}

func DeleteCategory(db *sql.DB, category models.Category) (err error) {
	sql := "DELETE FROM categories WHERE id = $1"
	_, err = db.Exec(sql, category.ID)
	return
}

func GetBooksByCategory(db *sql.DB, categoryID int) (result []models.Book, err error) {
	sqlQuery := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by
		FROM books 
		WHERE category_id = $1`

	rows, err := db.Query(sqlQuery, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageURL,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CategoryID,
			&book.CreatedAt,
			&book.CreatedBy,
			&book.ModifiedAt,
			&book.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, book)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}

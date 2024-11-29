package repository

import (
	"book-management-api/internal/models"
	"database/sql"
)

func CreateBook(db *sql.DB, payload models.CreateBookRequest) (err error) {
	//logic
	var thickness string
	if payload.TotalPage > 200 {
		thickness = "thick"
	} else {
		thickness = "thin"
	}

	sql := `INSERT INTO books 
        (title, description, release_year, price, total_page, thickness, category_id) 
        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(sql, payload.Title, payload.Description, payload.ReleaseYear, payload.Price, payload.TotalPage, thickness, payload.CategoryID)
	if err != nil {
		return err
	}
	return
}

func GetAllBooks(db *sql.DB) (result []models.Book, err error) {
	sql := "SELECT id, title, description, release_year, price, total_page, thickness, category_id FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book models.Book

		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CategoryID)
		if err != nil {
			return
		}

		result = append(result, book)
	}

	return
}

func GetBookByID(db *sql.DB, id int) (result models.Book, err error) {
	sqlQuery := "SELECT id, title, description, release_year, price, total_page, thickness, category_id FROM books WHERE id = $1"

	// Eksekusi query dengan parameter id
	err = db.QueryRow(sqlQuery, id).Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.ReleaseYear,
		&result.Price,
		&result.TotalPage,
		&result.Thickness,
		&result.CategoryID,
	)

	// Tangani jika tidak ada data ditemukan
	if err == sql.ErrNoRows {
		return result, nil // Kembalikan struct kosong tanpa error
	} else if err != nil {
		return result, err // Kembalikan error lain jika terjadi
	}

	return result, nil
}

func GetBooksByCategory(db *sql.DB, categoryID int) (result []models.Book, err error) {
	sqlQuery := `SELECT id, title, description, release_year, price, total_page, thickness, category_id
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
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CategoryID,
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

func DeleteBook(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM books WHERE id = $1"
	_, err = db.Exec(sql, id)
	return
}

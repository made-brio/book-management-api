package repository

import (
	"book-management-api/models"
	"database/sql"
)

func GetAllBooks(db *sql.DB) (result []models.Book, err error) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book models.Book

		err = rows.Scan(&book.ID,
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
			&book.ModifiedBy)
		if err != nil {
			return
		}

		result = append(result, book)
	}

	return
}

func CreateBook(db *sql.DB, book models.Book) (err error) {
	sql := `INSERT INTO books 
        (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by, modified_by) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err = db.Exec(sql, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedBy, book.ModifiedBy)
	if err != nil {
		return err
	}
	return
}

func GetBookByID(db *sql.DB, id int) (result models.Book, err error) {
	sqlQuery := "SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE id = $1"

	// Eksekusi query dengan parameter id
	err = db.QueryRow(sqlQuery, id).Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.ImageURL,
		&result.ReleaseYear,
		&result.Price,
		&result.TotalPage,
		&result.Thickness,
		&result.CategoryID,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.ModifiedAt,
		&result.ModifiedBy,
	)

	// Tangani jika tidak ada data ditemukan
	if err == sql.ErrNoRows {
		return result, nil // Kembalikan struct kosong tanpa error
	} else if err != nil {
		return result, err // Kembalikan error lain jika terjadi
	}

	return result, nil
}

func DeleteBook(db *sql.DB, book models.Book) (err error) {
	sql := "DELETE FROM books WHERE id = $1"
	_, err = db.Exec(sql, book.ID)
	return
}

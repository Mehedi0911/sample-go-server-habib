package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgURL      string  `json:"img_url" db:"image_url"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
	UpdatedAt   string  `json:"updated_at" db:"updated_at"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(p *Product) error
	Delete(id int) error
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}

}

func (r *productRepo) Create(product Product) (*Product, error) {
	query := `INSERT INTO products (
			title,
			description,
			price,
			image_url
	) VALUES (
			$1,
			$2,
			$3,
			$4
    ) RETURNING id`
	row := r.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImgURL)
	err := row.Scan(&product.ID)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to execute query: %v", err)
	}
	return &product, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	var product Product

	query := `
	SELECT 
	id, 
	title, 
	description, 
	price, 
	image_url 
	FROM products 
	WHERE id = $1`

	err := r.db.Get(&product, query, id)

	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Product with id %d not found", id)
		}
		return nil, fmt.Errorf("Failed to execute query: %v", err)
	}
	return &product, nil
}

func (r *productRepo) List() ([]*Product, error) {
	query := `SELECT * FROM products`
	var productList []*Product
	err := r.db.Select(&productList, query)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to execute query: %v", err)
	}
	return productList, nil
}

func (r *productRepo) Update(p *Product) error {
	query := `
		UPDATE products SET
			title = $1,
			description = $2,
			price = $3,
			image_url = $4
		WHERE id = $5
	`

	res, err := r.db.Exec(query,
		p.Title,
		p.Description,
		p.Price,
		p.ImgURL,
		p.ID,
	)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows check failed: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}

func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Failed to execute query: %v", err)
	}
	return nil
}

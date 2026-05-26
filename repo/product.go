package repo

import "fmt"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(id int) error
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(product Product) (*Product, error) {
	product.ID = len(r.productList) + 1
	r.productList = append(r.productList, &product)
	return &product, nil
}
func (r *productRepo) Get(id int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Update(p Product) (*Product, error) {
	for i, product := range r.productList {
		if product.ID == p.ID {
			r.productList[i] = &p
			return r.productList[i], nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

func (r *productRepo) Delete(id int) error {
	tempPrd := []*Product{}
	for _, product := range r.productList {
		if product.ID != id {
			tempPrd = append(tempPrd, product)
		}
	}
	r.productList = tempPrd
	return nil
}

func generateInitialProducts(r *productRepo) {
	pd1 := Product{
		ID:          1,
		Title:       "Product 1",
		Description: "Description of Product Orange",
		Price:       10.99,
		ImgURL:      "https://example.com/product1.jpg",
	}
	pd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Fresh red apple",
		Price:       5.99,
		ImgURL:      "https://example.com/apple.jpg",
	}

	pd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Sweet ripe bananas",
		Price:       3.49,
		ImgURL:      "https://example.com/banana.jpg",
	}

	pd4 := Product{
		ID:          4,
		Title:       "Mango",
		Description: "Juicy tropical mango",
		Price:       7.99,
		ImgURL:      "https://example.com/mango.jpg",
	}

	pd5 := Product{
		ID:          5,
		Title:       "Orange",
		Description: "Citrus fresh orange",
		Price:       4.25,
		ImgURL:      "https://example.com/orange.jpg",
	}

	pd6 := Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "Sweet and tangy pineapple",
		Price:       6.75,
		ImgURL:      "https://example.com/pineapple.jpg",
	}

	pd7 := Product{
		ID:          7,
		Title:       "Strawberry",
		Description: "Fresh strawberries",
		Price:       8.50,
		ImgURL:      "https://example.com/strawberry.jpg",
	}
	r.productList = append(r.productList, &pd1, &pd2, &pd3, &pd4, &pd5, &pd6, &pd7)

}

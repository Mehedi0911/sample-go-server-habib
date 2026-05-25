package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"img_url"`
}

func GetList() []Product {
	return productList
}

func GetProductByID(id int) *Product {
	for _, product := range productList {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func CreateProduct(product Product) *Product {
	product.ID = len(productList) + 1
	productList = append(productList, product)
	return &product
}

func UpdateProduct(id int, updatedProduct Product) *Product {
	for i, product := range productList {
		if product.ID == id {
			updatedProduct.ID = id
			productList[i] = updatedProduct
			return &productList[i]
		}
	}
	return nil
}

func DeleteProduct(id int) bool {
	tempPrd := []Product{}
	for _, product := range productList {
		if product.ID != id {
			tempPrd = append(tempPrd, product)
		}
	}
	productList = tempPrd
	return true
}

func init() {
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
	productList = append(productList, pd1, pd2, pd3, pd4, pd5, pd6, pd7)

}

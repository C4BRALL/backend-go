package entity

type SellerRepositoryInterface interface {
	Create(seller *Seller) error
	FindByEmail(email string) (*Seller, error)
	FindById(id string) (*Seller, error)
	FindByDocument(document string) (*Seller, error)
	FindAll(page, limit int, sort string) ([]Seller, error)
	Delete(document string) error
	Update(seller *Seller) error
}

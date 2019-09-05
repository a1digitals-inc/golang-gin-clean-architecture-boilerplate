package repositories

type ProductRepository interface {
	Save(msg string) (string, error)
	FindAll(map[string]string) (interface{}, error)
	FindById(msg string) (string, error)
	Update(msg string) (string, error)
	Delete(msg string) (string, error)
}

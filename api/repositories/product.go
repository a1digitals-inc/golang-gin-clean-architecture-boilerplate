package repositories

type ProductRepository interface {
	Save(msg string) (string, error)
	FindAll(map[string]interface{}) (interface{}, error)
	FindById(msg string) (string, error)
	Update(msg string) (string, error)
	Delete(msg string) (string, error)
}

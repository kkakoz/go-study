package gomock_study


type Repository interface {
	Create(key string, value string) error
	Retrieve(key string) ([]byte, error)
	Update(key string, value string) error
	Delete(key string) error
	GetInt(key string) int
}

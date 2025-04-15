package product

type ProductUseCase struct {
	repository Repository
}

func NewProductUseCase(repository Repository) *ProductUseCase {
	return &ProductUseCase{repository}
}
func (uc *ProductUseCase) GetProductByID(id int) (*Product, error) {
	return uc.repository.GetProductByID(id)
}

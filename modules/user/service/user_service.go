package service

import (
	"golang-clean-arch/modules/user/dto"
)

type UserService interface {
	GetUsers() (dto.UserRes, error)
}

type UserServiceImpl struct {
	//UserRepository repositories.Repositories
	//transaction       *transaction.TransactionImpl
}

func NewUserService() UserServiceImpl {
	return UserServiceImpl{}
}

func (s UserServiceImpl) GetUsers() (dto.UserRes, error) {
	//productList, err := s.ProductRepository.GetProductsList(ctx)
	//if err != nil {
	//	log.Err(err).Msg("Error fetch productList from DB")
	//}
	//productsResp := dto.CreateProductsListResponse(productList)

	userRes := dto.UserRes{
		FirstName: "Khoa",
		LastName:  "Nguyen",
		Email:     "khoand@gmail.com",
	}
	return userRes, nil
}

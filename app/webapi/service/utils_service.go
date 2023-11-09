package service

import (
	"district/model"
	repository "district/repository/handler"
)

type UtilsService struct {
	repositoryPool *repository.RepositoryPool
}

func NewUtilsService(repositoryPool *repository.RepositoryPool) *UtilsService {
	return &UtilsService{repositoryPool: repositoryPool}
}

func (us *UtilsService) PopulateDatabase() error {

	// Users setup

	if err := us.repositoryPool.UserRepository.CreateUser(&model.User{
		Identification: 1,
		Email:          "admin@district.com",
		Username:       "admin",
		Password:       "district2023",
		Address:        "District Land",
		Balance:        10000,
		IsAdmin:        true,
	}); err != nil {
		return err
	}

	if err := us.repositoryPool.UserRepository.CreateUser(&model.User{
		Identification: 2,
		Email:          "danny@gmail.com",
		Username:       "danny",
		Password:       "batman01",
		Address:        "Colombia",
		Balance:        500,
		IsAdmin:        false,
	}); err != nil {
		return err
	}

	if err := us.repositoryPool.UserRepository.CreateUser(&model.User{
		Identification: 3,
		Email:          "attacker@gmail.com",
		Username:       "attacker",
		Password:       "tryhackme",
		Address:        "United States",
		Balance:        50,
		IsAdmin:        false,
	}); err != nil {
		return err
	}

	// Products setup

	if _, err := us.repositoryPool.ProductRepository.CreateProduct(&model.Product{
		Name:        "Apple",
		Description: "its a tasty apple",
		Stock:       10,
		Price:       50,
	}); err != nil {
		return err
	}

	if _, err := us.repositoryPool.ProductRepository.CreateProduct(&model.Product{
		Name:        "Pear",
		Description: "its a cheap pear",
		Stock:       2,
		Price:       40,
	}); err != nil {
		return err
	}

	if _, err := us.repositoryPool.ProductRepository.CreateProduct(&model.Product{
		Name:        "Bicycle",
		Description: "easy to assemble, with tools and instructions included",
		Stock:       5,
		Price:       500,
	}); err != nil {
		return err
	}

	if _, err := us.repositoryPool.ProductRepository.CreateProduct(&model.Product{
		Name:        "Xbox Series X",
		Description: "the best console in this generation",
		Stock:       3,
		Price:       700,
	}); err != nil {
		return err
	}

	if _, err := us.repositoryPool.ProductRepository.CreateProduct(&model.Product{
		Name:        "Play Station 5",
		Description: "marvel at amazing graphics and experience new features",
		Stock:       5,
		Price:       650,
	}); err != nil {
		return err
	}

	// Reviews setup

	if err := us.repositoryPool.ReviewRepository.CreateReview(&model.Review{
		ProductID: 1,
		UserID:    1,
		UserEmail: "admin@district.com",
		Content:   "i really liked this apple",
	}); err != nil {
		return err
	}

	if err := us.repositoryPool.ReviewRepository.CreateReview(&model.Review{
		ProductID: 1,
		UserID:    2,
		UserEmail: "danny@gmail.com",
		Content:   "this apple was not that good",
	}); err != nil {
		return err
	}

	if err := us.repositoryPool.ReviewRepository.CreateReview(&model.Review{
		ProductID: 2,
		UserID:    2,
		UserEmail: "danny@gmail.com",
		Content:   "i feel that this pear could have been better",
	}); err != nil {
		return err
	}

	if err := us.repositoryPool.ReviewRepository.CreateReview(&model.Review{
		ProductID: 3,
		UserID:    1,
		UserEmail: "admin@district.com",
		Content:   "i have had this bicycle for 2 years and its perfect",
	}); err != nil {
		return err
	}

	if err := us.repositoryPool.ReviewRepository.CreateReview(&model.Review{
		ProductID: 4,
		UserID:    1,
		UserEmail: "admin@district.com",
		Content:   "my Xbox Series X do not work as expected",
	}); err != nil {
		return err
	}

	if err := us.repositoryPool.ReviewRepository.CreateReview(&model.Review{
		ProductID: 5,
		UserID:    2,
		UserEmail: "danny@gmail.com",
		Content:   "Play Station is definitely the best console for gaming",
	}); err != nil {
		return err
	}

	return nil
}

func (us *UtilsService) ClearDatabase() error {
	if err := us.repositoryPool.UserRepository.DeleteUsersTable(); err != nil {
		return err
	}
	if err := us.repositoryPool.UserRepository.CreateUsersTable(); err != nil {
		return err
	}

	if err := us.repositoryPool.ProductRepository.DeleteProductsTable(); err != nil {
		return err
	}
	if err := us.repositoryPool.ProductRepository.CreateProductsTable(); err != nil {
		return err
	}

	if err := us.repositoryPool.ReviewRepository.DeleteReviewsTable(); err != nil {
		return err
	}
	if err := us.repositoryPool.ReviewRepository.CreateReviewsTable(); err != nil {
		return err
	}

	return nil
}

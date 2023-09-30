package services

import (
	"golang-mongodb-repository/dtos"
	"golang-mongodb-repository/models"
	"golang-mongodb-repository/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService interface {
	InsertProduct(product models.Product) (*dtos.ProductDTO, error)
	GetAllProducts() ([]models.Product, error)
	DeleteProduct(id primitive.ObjectID) (bool, error)
}

type DefaultProductService struct {
	Repo repository.ProductRepository
}

func (service DefaultProductService) InsertProduct(product models.Product) (*dtos.ProductDTO, error) {
	var res dtos.ProductDTO
	if len(product.Title) <= 2 {
		res.Status = false
		return &res, nil
	}

	result, err := service.Repo.Insert(product)

	if err != nil || result == false {
		res.Status = false
		return &res, err
	}

	res = dtos.ProductDTO{Status: result}
	return &res, nil
}

func (service DefaultProductService) GetAllProducts() ([]models.Product, error) {
	result, err := service.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service DefaultProductService) DeleteProduct(id primitive.ObjectID) (bool, error) {
	result, err := service.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func NewProductService(Repo repository.ProductRepository) ProductService {
	return &DefaultProductService{Repo: Repo}
}

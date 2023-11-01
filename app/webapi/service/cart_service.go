package service

import (
	request "district/controller/request"
	"district/model"
	repository "district/repository/handler"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type CartService struct {
	repositoryPool *repository.RepositoryPool
}

func NewCartService(repositoryPool *repository.RepositoryPool) *CartService {
	return &CartService{repositoryPool: repositoryPool}
}

func (cs *CartService) PlaceOrder(authToken, cartToken string) (*float64, error) {
	decodedAuthToken, err := base64.StdEncoding.DecodeString(authToken)
	if err != nil {
		return nil, err
	}
	authTokenValues := strings.Split(string(decodedAuthToken), ":")

	identification, err := strconv.Atoi(authTokenValues[0])
	if err != nil {
		return nil, err
	}

	user, err := cs.repositoryPool.UserRepository.GetUserByIdentification(identification)
	if err != nil {
		return nil, err
	}

	cart, err := cs.decodeCartCookie(cartToken)
	if err != nil {
		return nil, err
	}

	if len(cart) == 0 {
		return nil, fmt.Errorf("your cart does not have products.")
	}

	totalPrice, err := cs.getCartTotal(&cart)
	if err != nil {
		return nil, err
	}

	if user.Balance < *totalPrice {
		return nil, fmt.Errorf("there is not enough balance on your account.")
	}

	user.Balance -= *totalPrice
	cs.repositoryPool.UserRepository.UpdateUser(user)

	return &user.Balance, nil
}

func (cs *CartService) GetCartInformation(cartToken string) (*float64, *model.Cart, error) {
	cart, err := cs.decodeCartCookie(cartToken)
	if err != nil {
		return nil, nil, err
	}

	total, err := cs.getCartTotal(&cart)
	if err != nil {
		return nil, nil, err
	}

	return total, &cart, nil
}

func (cs *CartService) AddItemToCart(cartToken string, request request.AddCartItemRequest) (*string, error) {
	cart, err := cs.decodeCartCookie(cartToken)
	if err != nil {
		return nil, err
	}

	itemExists := false
	for i, existingItem := range cart {
		if existingItem.ProductID == *request.ProductID {
			if existingItem.Price != *request.Price {
				cart[i].Price = *request.Price
			}
			cart[i].Quantity += *request.Quantity
			if cart[i].Quantity <= 0 {
				cart = append(cart[:i], cart[i+1:]...)
			}
			itemExists = true
			break
		}
	}

	if !itemExists {
		item := model.CartItem{
			ProductID: *request.ProductID,
			Quantity:  *request.Quantity,
			Price:     *request.Price,
		}

		cart = append(cart, item)
	}

	encodedCart, err := json.Marshal(cart)
	if err != nil {
		return nil, err
	}

	encodedToken := base64.StdEncoding.EncodeToString(encodedCart)

	return &encodedToken, nil
}

func (cs *CartService) UpdateProductQuantity(cartToken string, request request.UpdateCartItemRequest) (*string, error) {
	cart, err := cs.decodeCartCookie(cartToken)
	if err != nil {
		return nil, err
	}

	itemExists := false
	for i, item := range cart {
		if item.ProductID == *request.ProductID {
			if *request.Quantity <= 0 {
				cart = append(cart[:i], cart[i+1:]...)
			} else {
				cart[i].Quantity = *request.Quantity
			}
			itemExists = true
			break
		}
	}

	if !itemExists {
		return nil, fmt.Errorf("product with ID %d not found in cart.", request.ProductID)
	}

	encodedCart, err := json.Marshal(cart)
	if err != nil {
		return nil, err
	}

	encodedToken := base64.StdEncoding.EncodeToString(encodedCart)

	return &encodedToken, nil
}

func (cs *CartService) RemoveItemFromCart(cartToken string, productId int) (*string, error) {
	cart, err := cs.decodeCartCookie(cartToken)
	if err != nil {
		return nil, err
	}

	itemExists := false
	for i, item := range cart {
		if item.ProductID == productId {
			cart = append(cart[:i], cart[i+1:]...)
			itemExists = true
			break
		}
	}

	if !itemExists {
		return nil, fmt.Errorf("product with ID %d not found in cart.", productId)
	}

	encodedCart, err := json.Marshal(cart)
	if err != nil {
		return nil, err
	}

	encodedToken := base64.StdEncoding.EncodeToString(encodedCart)

	return &encodedToken, nil
}

func (cs *CartService) getCartTotal(cart *model.Cart) (*float64, error) {
	var total float64

	for _, item := range *cart {
		total += item.Price * float64(item.Quantity)
	}

	return &total, nil
}

func (s *CartService) decodeCartCookie(cartToken string) (model.Cart, error) {
	var cart model.Cart

	decodedToken, err := base64.StdEncoding.DecodeString(cartToken)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(decodedToken, &cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

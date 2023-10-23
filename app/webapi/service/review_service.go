package service

import (
	"district/model"
	dto "district/model/dto"
	"district/repository"
	"encoding/base64"
	"strconv"
	"strings"
)

type ReviewService struct {
	reviewRepository *repository.ReviewRepository
	userRepository   *repository.UserRepository
}

func NewReviewService(reviewRepository *repository.ReviewRepository, userRepository *repository.UserRepository) *ReviewService {
	return &ReviewService{
		reviewRepository: reviewRepository,
		userRepository:   userRepository,
	}
}

func (r *ReviewService) CreateReview(token string, review *model.Review) error {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}

	tokenValues := strings.Split(string(decodedToken), ":")
	identification, err := strconv.Atoi(tokenValues[0])
	if err != nil {
		return err
	}

	user, err := r.userRepository.GetUserByIdentification(identification)
	if err != nil {
		return err
	}

	review.UserEmail = user.Email

	return r.reviewRepository.CreateReview(review)
}

func (r *ReviewService) GetReviewById(id int) (*dto.ReviewDTO, error) {
	review, err := r.reviewRepository.GetReviewById(id)
	if err != nil {
		return nil, err
	}

	reviewDto := dto.ConvertToReviewDTO(review)

	return reviewDto, nil
}

func (r *ReviewService) UpdateReview(review *model.Review) error {
	_, err := r.reviewRepository.GetReviewById(review.ID)
	if err != nil {
		return err
	}

	return r.reviewRepository.UpdateReview(review)
}

package service

import (
	"district/model"
	dto "district/model/dto"
	"district/repository"
)

type ReviewService struct {
	reviewRepository *repository.ReviewRepository
}

func NewReviewService(reviewRepository *repository.ReviewRepository) *ReviewService {
	return &ReviewService{reviewRepository: reviewRepository}
}

func (r *ReviewService) CreateReview(review *model.Review) error {
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

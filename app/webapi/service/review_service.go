package service

import (
	"district/model"
	dto "district/model/dto"
	repository "district/repository/handler"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type ReviewService struct {
	repositoryPool *repository.RepositoryPool
}

func NewReviewService(repositoryPool *repository.RepositoryPool) *ReviewService {
	return &ReviewService{repositoryPool: repositoryPool}
}

func (rs *ReviewService) CreateReview(token string, review *model.Review) error {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}

	tokenValues := strings.Split(string(decodedToken), ":")
	identification, err := strconv.Atoi(tokenValues[0])
	if err != nil {
		return err
	}

	user, err := rs.repositoryPool.UserRepository.GetUserByIdentification(identification)
	if err != nil {
		return err
	}

	review.UserID = user.Identification
	review.UserEmail = user.Email

	return rs.repositoryPool.ReviewRepository.CreateReview(review)
}

func (rs *ReviewService) GetReviewById(id int) (*dto.ReviewDTO, error) {
	review, err := rs.repositoryPool.ReviewRepository.GetReviewById(id)
	if err != nil {
		return nil, err
	}

	reviewDto := dto.ConvertToReviewDTO(review)

	return reviewDto, nil
}

func (rs *ReviewService) UpdateReview(review *model.Review) error {
	_, err := rs.repositoryPool.ReviewRepository.GetReviewById(review.ID)
	if err != nil {
		return err
	}

	return rs.repositoryPool.ReviewRepository.UpdateReview(review)
}

func (rs *ReviewService) DeleteReviewById(token string, id int) error {
	review, err := rs.repositoryPool.ReviewRepository.GetReviewById(id)
	if err != nil {
		return err
	}

	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}

	tokenValues := strings.Split(string(decodedToken), ":")
	identification, err := strconv.Atoi(tokenValues[0])
	if err != nil {
		return fmt.Errorf("your session token is not valid.")
	}
	is_admin, err := strconv.ParseBool(tokenValues[3])
	if err != nil {
		return fmt.Errorf("your session token is not valid.")
	}

	if identification != review.UserID && !is_admin {
		return fmt.Errorf("you have no permissions to delete this review.")
	}

	return rs.repositoryPool.ReviewRepository.DeleteReview(id)
}

func (rs *ReviewService) DeleteReviewsTable() error {
	if err := rs.repositoryPool.ReviewRepository.DeleteReviewsTable(); err != nil {
		return err
	}

	return nil
}

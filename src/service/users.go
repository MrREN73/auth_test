package service

import (
	"auth/src/common"
	"auth/src/models"
	"log"

	"errors"

	"github.com/jinzhu/gorm"
)

func (s *Service) GetUser(userID common.BaseID) (*models.User, error) {
	if common.IsEmptyID(userID) {
		return nil, nil
	}

	result := models.User{}

	if err := s.db.Take(&result, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Service) CreateUser(user models.User) (*models.User, error) {
	if user.Secret == "" {
		return nil, errors.New("parameter secret is empty")
	}

	user.CommonInfo.Hash()

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return s.GetUser(user.ID)
}

func (s *Service) ListUser(r models.RequestUser) ([]models.User, error) {
	q := addWhere(s.db.New(), r.CommonInfo)

	users := []models.User{}
	if err := r.PageRequest.Append(q).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func addWhere(q *gorm.DB, filter models.CommonInfo) *gorm.DB {
	if filter.Email != "" {
		q = q.Where("email ILIKE ?", "%"+filter.Email+"%")
	}

	if filter.FirstName != "" {
		q = q.Where("first_name ILIKE ?", "%"+filter.FirstName+"%")
	}

	if filter.LastName != "" {
		q = q.Where("last_name ILIKE ?", "%"+filter.LastName+"%")
	}

	if filter.MiddleName != "" {
		q = q.Where("middle_name ILIKE ?", "%"+filter.MiddleName+"%")
	}

	if !filter.Role.IsEmpty() {
		q = q.Where("role ILIKE ?", "%"+filter.Role.String()+"%")
	}

	return q
}

func (s *Service) UpdateUser(user models.User) (*models.User, error) {
	return nil, nil
}

func (s *Service) DeleteUser(userID common.BaseID) error {
	user := models.User{ID: userID}

	if err := s.db.Delete(&user).Error; err != nil {
		return err
	}

	log.Printf("User deleted: ID:%d", userID)

	return nil
}

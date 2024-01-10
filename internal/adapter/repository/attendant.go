package repository

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

type Attendant struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type AttendantRepository struct {
	db *gorm.DB
}

func NewAttendantRepository(db *gorm.DB) *AttendantRepository {
	return &AttendantRepository{
		db,
	}
}

func (c *AttendantRepository) CreateAttendant(name string) (*entity.Attendant, error) {
	attendant := &Attendant{
		Name: name,
	}

	var err error
	if err = c.db.Create(attendant).Error; err != nil {
		return nil, err
	}

	var result Attendant
	c.db.Where("name = ?", name).First(&result)

	return &entity.Attendant{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (c *AttendantRepository) GetAttendantById(id uint) (*entity.Attendant, error) {
	var result Attendant
	if err := c.db.First(&result, id).Error; err != nil {
		return nil, fmt.Errorf("error to find attendant with id %d - %v", id, err)
	}

	return &entity.Attendant{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (c *AttendantRepository) GetAttendants() ([]*entity.Attendant, error) {
	var results []*Attendant
	if err := c.db.Find(&results).Error; err != nil {
		return nil, err
	}

	var attendants []*entity.Attendant
	for _, result := range results {
		attendants = append(attendants, &entity.Attendant{
			ID:        result.ID,
			Name:      result.Name,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		})
	}

	return attendants, nil
}

func (c *AttendantRepository) UpdateAttendant(attendant *entity.Attendant) (*entity.Attendant, error) {
	var result Attendant
	if err := c.db.First(&result, attendant.ID).Error; err != nil {
		return nil, err
	}

	result.Name = attendant.Name

	if err := c.db.Save(&result).Error; err != nil {
		return nil, err
	}

	return c.GetAttendantById(attendant.ID)
}

func (c *AttendantRepository) DeleteAttendant(id uint) error {
	if err := c.db.Delete(&Attendant{}, id).Error; err != nil {
		return err
	}

	return nil
}

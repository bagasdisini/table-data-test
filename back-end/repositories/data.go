package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type DataRepository interface {
	ShowData() ([]models.Data, error)
	GetDataByID(ID int) (models.Data, error)
	CreateData(Data models.Data) (models.Data, error)
	UpdateData(Data models.Data, ID int) (models.Data, error)
	DeleteData(Data models.Data, ID int) (models.Data, error)
	ShowStatus() ([]models.Status, error)
	GetStatusByID(ID int) (models.Status, error)
}

func RepositoryData(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowData() ([]models.Data, error) {
	var Data []models.Data
	err := r.db.Preload("Status").Find(&Data).Error

	return Data, err
}

func (r *repository) GetDataByID(ID int) (models.Data, error) {
	var Data models.Data
	err := r.db.Preload("Status").First(&Data, ID).Error

	return Data, err
}

func (r *repository) CreateData(Data models.Data) (models.Data, error) {
	err := r.db.Create(&Data).Error

	return Data, err
}

func (r *repository) UpdateData(Data models.Data, ID int) (models.Data, error) {
	err := r.db.Model(&Data).Where("id=?", ID).Updates(&Data).Error

	return Data, err
}

func (r *repository) DeleteData(Data models.Data, ID int) (models.Data, error) {
	err := r.db.Delete(&Data).Error

	return Data, err
}

func (r *repository) ShowStatus() ([]models.Status, error) {
	var Status []models.Status
	err := r.db.Find(&Status).Error

	return Status, err
}

func (r *repository) GetStatusByID(ID int) (models.Status, error) {
	var Status models.Status
	err := r.db.First(&Status, ID).Error

	return Status, err
}

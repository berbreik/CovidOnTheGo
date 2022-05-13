package http

import "assignment/models"

type HospitalHandler interface {
	AddService(patient *models.Patient) error
	GetService(id int) (*models.Patient, error)
	UpdateService(id int, patient *models.Patient) error
	RemoveService(id int) error
	GetAllService() ([]*models.Patient, error)
}

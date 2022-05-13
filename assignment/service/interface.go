package service

import "assignment/models"

type HospitalService interface {
	AdmitPatient(patient *models.Patient) error
	GetPatient(id int) (*models.Patient, error)
	UpdatePatient(id int, patient *models.Patient) error
	RemovePatient(id int) error
	GetAllPatients() ([]*models.Patient, error)
}

package db

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/newtoallofthis123/patients/types"
	"github.com/newtoallofthis123/ranhash"
)

func requestToPatient(pr types.PatientRequest, patientId string) types.Patient {
	return types.Patient{
		PatientID:   patientId,
		FirstName:   pr.FirstName,
		LastName:    pr.LastName,
		Age:         pr.Age,
		Gender:      pr.Gender,
		Address:     pr.Address,
		Identity:    pr.Identity,
		Phone:       pr.Phone,
		Email:       pr.Email,
		Description: pr.Description,
		Recurring:   pr.Recurring,
	}
}

func (d *Store) NewPatient(pr types.PatientRequest) (types.Patient, error) {
	patientId := ranhash.GenerateRandomString(12)

	_, err := d.pq.Insert("patients").Columns("id", "first_name", "last_name", "age", "gender", "address", "identity", "phone", "email", "description", "recurring").Values(patientId, pr.FirstName, pr.LastName, pr.Age, pr.Gender, pr.Address, pr.Identity, pr.Phone, pr.Email, pr.Description, pr.Recurring).RunWith(d.db).Exec()

	return requestToPatient(pr, patientId), err
}

func (d *Store) GetPatient(patientId string) (types.Patient, error) {
	row := d.pq.Select("*").From("patients").Where(squirrel.Eq{"id": patientId}).RunWith(d.db).QueryRow()

	var pr types.Patient
	var createdAt time.Time

	err := row.Scan(&pr.PatientID, &pr.FirstName, &pr.LastName, &pr.Age, &pr.Gender, &pr.Address, &pr.Identity, &pr.Phone, &pr.Email, &pr.Description, &pr.Recurring, &createdAt)
	if err != nil {
		return types.Patient{}, err
	}

	return pr, nil
}

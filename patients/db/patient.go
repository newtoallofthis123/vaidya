package db

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/newtoallofthis123/patients/types"
	"github.com/newtoallofthis123/ranhash"
)

func (d *Store) NewPatient(pr types.PatientRequest) (string, error) {
	patientId := ranhash.GenerateRandomString(12)

	_, err := d.pq.Insert("patients").Columns("id", "name", "age", "gender", "address", "identity", "phone", "conditions", "problems", "description", "created_at").
		Values(patientId, pr.Name, pr.Age, pr.Gender, pr.Address, pr.Identity, pr.Phone, pr.Conditions, pr.Problems, pr.Description, time.Now()).
		RunWith(d.db).Exec()

	return patientId, err
}

func (d *Store) GetPatient(patientId string) (types.Patient, error) {
	row := d.pq.Select("*").From("patients").Where(squirrel.Eq{"id": patientId}).RunWith(d.db).QueryRow()

	var pr types.Patient
	var createdAt time.Time

	err := row.Scan(
		&pr.PatientID,
		&pr.Name,
		&pr.Age,
		&pr.Gender,
		&pr.Address,
		&pr.Identity,
		&pr.Phone,
		&pr.Conditions,
		&pr.Problems,
		&pr.Description,
		&createdAt,
	)
	if err != nil {
		return types.Patient{}, err
	}

	return pr, nil
}

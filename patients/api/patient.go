package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/newtoallofthis123/patients/types"
)

func (s *ApiServer) handlePatientCreate(c *gin.Context) {
	var patientRequest types.PatientRequest

	if err := c.ShouldBindBodyWithJSON(&patientRequest); err != nil {
		s.logger.Error(fmt.Sprintln("Error binding request to types.patientRequest:", err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := s.store.NewPatient(patientRequest)
	if err != nil {
		s.logger.Error(fmt.Sprintln("Error creating a new patient:", err))
	}

	c.JSON(201, gin.H{"id": id})
}

func (s *ApiServer) handlePatientGet(c *gin.Context) {
	id := c.Param("id")

	patient, err := s.store.GetPatient(id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error getting patient with id %s: %s", id, err))
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(200, patient)
}

func (s *ApiServer) handlePatientUpdate(c *gin.Context) {
	id := c.Param("id")

	var patientRequest types.Patient

	if err := c.ShouldBindBodyWithJSON(&patientRequest); err != nil {
		s.logger.Error(fmt.Sprintln("Error binding request to types.patientRequest:", err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := s.store.EditPatient(patientRequest)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error updating patient with id %s: %s", id, err))
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Patient updated successfully"})
}

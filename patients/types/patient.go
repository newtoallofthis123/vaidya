package types

type PatientRequest struct {
	Name        string `json:"name,omitempty"`
	Age         uint32 `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Address     string `json:"address,omitempty"`
	Identity    string `json:"identity,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Conditions  string `json:"conditions,omitempty"`
	Problems    string `json:"problems,omitempty"`
	Description string `json:"description,omitempty"`
}

type Patient struct {
	PatientID   string `json:"patient_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Age         uint32 `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Address     string `json:"address,omitempty"`
	Identity    string `json:"identity,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Conditions  string `json:"conditions,omitempty"`
	Problems    string `json:"problems,omitempty"`
	Description string `json:"description,omitempty"`
	Medicines   string `json:"medicines,omitempty"`
	Diagnosis   string `json:"diagnosis,omitempty"`
	NextSession string `json:"next_session,omitempty"`
}

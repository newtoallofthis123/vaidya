"use client";

import { EMRForm, PatientFormData } from "@/components/EMRForm";

function constructAddress(
  street: string,
  city: string,
  state: string,
  country: string,
) {
  const ret = `${street}, ${city}, ${state}, ${country}`;
  if (ret === "undefined, undefined, undefined, undefined") {
    return "";
  }

  return ret;
}

export default function FormPage() {
  const handleSubmit = (data: PatientFormData) => {
    console.log("Form submitted with data:", data);
  };

  const jsonInfo = window.localStorage.getItem("info") ?? "";
  if (jsonInfo === "") {
    window.location.href = "/";
  }
  const info = JSON.parse(jsonInfo);
  const message = window.localStorage.getItem("message") ?? "";
  const jsonSymptoms = window.localStorage.getItem("symptoms") ?? "";
  let symptoms = [];
  if (jsonSymptoms != "") {
    symptoms = JSON.parse(jsonSymptoms);
  }

  const problems = [];
  for (let i = 0; i < symptoms.length; i++) {
    if (symptoms[i]["type"] === "symptoms") {
      problems.push(symptoms[i]["name"]);
    }
  }

  const samplePatientData: PatientFormData = {
    patient_id: info["pid"] ?? "",
    first_name: info["firstname"] ?? "",
    last_name: info["lastname"] ?? "",
    age: info["age"] ?? "",
    gender: info["gender"] ?? "",
    address: constructAddress(
      info["street"],
      info["city"],
      info["state"],
      info["country"],
    ),
    identity: "",
    phone: info["phonenumber"] ?? "",
    email: info["email"] ?? "",
    description: message,
    recurring: false,
    problems: problems,
    medicines: [],
    conditions: [],
    diagnosis: "",
    next_session: "2023-07-15",
  };
  return (
    <div>
      <EMRForm onSubmit={handleSubmit} initialData={samplePatientData} />
    </div>
  );
}

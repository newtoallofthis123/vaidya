"use client";

import { EMRForm, PatientFormData } from "@/components/EMRForm";

function constructAddress(
  street: string,
  city: string,
  state: string,
  country: string,
) {
  const ret = [];
  if (street != "undefined") {
    ret.push(street);
  }
  if (city != "undefined") {
    ret.push(city);
  }
  if (state != "undefined") {
    ret.push(state);
  }
  if (country != "undefined") {
    ret.push(country);
  }

  console.log(ret);
  return ret.join(", ");
}

function constructData(info: unknown): PatientFormData {
  return {
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
    description: info["message"],
    recurring: false,
    problems: info["problems"],
    medicines: info["medicines"],
    conditions: info["conditions"],
    diagnosis: "",
    next_session: "2023-07-15",
  };
}

export default function FormPage() {
  const handleSubmit = (data: PatientFormData) => {
    console.log("Form submitted with data:", data);
  };

  let jsonInfo = "";
  jsonInfo = window.localStorage.getItem("info") ?? "";
  if (jsonInfo === "") {
    if (typeof window !== "undefined") {
      window.location.href = "/";
    }
  }
  const info = JSON.parse(jsonInfo);
  const message = window.localStorage.getItem("description") ?? "";
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

  info["message"] = message;
  const patientData = constructData(info);

  return (
    <div>
      <EMRForm onSubmit={handleSubmit} initialData={patientData} />
    </div>
  );
}

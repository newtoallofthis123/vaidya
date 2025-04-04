"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import { EMRForm, PatientFormData } from "@/components/EMRForm";
import Link from "next/link";
import { Bot } from "lucide-react";

const PatientPage = () => {
  const params = useParams();
  const [patientData, setPatientData] = useState<PatientFormData>();
  const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;

  useEffect(() => {
    const fetchPatient = async () => {
      try {
        const response = await fetch(
          `http://${BACKEND_URL}/patients/${params.id}`,
        );
        if (!response.ok) {
          throw new Error("Failed to fetch patient data");
        }
        const data = await response.json();
        if (data["problems"]) {
          data["problems"] = JSON.parse(data["problems"]);
        }
        if (data["conditions"]) {
          data["conditions"] = JSON.parse(data["conditions"]);
        }
        if (data["medicines"]) {
          data["medicines"] = JSON.parse(data["medicines"]);
        }
        setPatientData(data);
      } catch (error) {
        console.error("Error fetching patient:", error);
      }
    };

    fetchPatient();
  }, [params.id, BACKEND_URL]);

  if (!patientData) {
    return <div>Loading...</div>;
  }

  return (
    <div className="py-4">
      <Link
        href={`/chat/${params.id}`}
        className="fixed bottom-6 right-6 p-4 rounded-full shadow-lg hover:bg-primary/50 transition-colors"
        aria-label="Open Chat"
      >
        <Bot />
      </Link>
      <EMRForm initialData={patientData} />
    </div>
  );
};

export default PatientPage;

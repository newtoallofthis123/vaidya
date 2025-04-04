import { EMRForm } from "@/components/EMRForm";
import Chat from "./chat";

const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;

async function getPatient(id: string) {
  try {
    const response = await fetch(`http://${BACKEND_URL}/patients/${id}`);

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

    return data;
  } catch (error) {
    console.error("Error fetching patient:", error);
    throw error;
  }
}

const PatientPage = async ({ params }) => {
  const { id } = await params;
  const patientData = await getPatient(id);
  const googleApikey = process.env.GOOGLE_GENERATIVE_AI_API_KEY;
  if (!googleApikey) {
    console.error("Google Generative AI API key not found");
    return null;
  }

  return (
    <div className="py-4 gap-4 flex">
      <div className="w-1/2">
        <EMRForm initialData={patientData} />
      </div>
      <div className="w-1/2">
        <Chat apiKey={googleApikey} patient={JSON.stringify(patientData)} />
      </div>
    </div>
  );
};

export default PatientPage;

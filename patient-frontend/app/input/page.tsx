"use client";

import { useState, useRef, useEffect } from "react";
import { Button } from "@/components/ui/button";
import ListeningAnimation from "@/components/custom/talking_circle";
import WavingEmoji from "@/components/custom/wavingemoji";
import ThinkingAnimation from "@/components/custom/thinking_circle";
import { MicIcon, StopCircleIcon } from "lucide-react";
import { useSearchParams } from "next/navigation";
import { InitialForm, InitialPatientFormData } from "@/components/InitialForm";

const speak = (text) => {
  if ("speechSynthesis" in window) {
    const utterance = new SpeechSynthesisUtterance(text);
    window.speechSynthesis.speak(utterance);
    console.log("Inited");
  } else {
    console.log("Speech synthesis not supported");
  }
};

const Conversation = () => {
  // read query params
  const params = useSearchParams();

  const [isRecording, setIsRecording] = useState(false);
  const [message, setMessage] = useState("");
  const [next, setNext] = useState("");
  const [form, setForm] = useState<InitialPatientFormData>();
  const [page, setPage] = useState(
    params.get("page") ? parseInt(params.get("page")) : 0,
  );
  const mediaRecorderRef = useRef<MediaRecorder | null>(null);
  const audioChunksRef = useRef([]);
  const sockRef = useRef<WebSocket | null>(null);
  const [thinking, setThinking] = useState(false);
  const [thoughts, setThoughts] = useState<string>("");

  const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;

  function convertToFormData(data): InitialPatientFormData {
    const problems = [];
    if (data["problems"]) {
      for (const problem of data["problems"]) {
        if (typeof problem === "string") {
          problems.push(problem);
        } else {
          problems.push(problem["name"]);
        }
      }
    }

    return {
      name: data["name"],
      age: data["age"],
      gender: data["gender"],
      address: data["address"],
      identity: data["identity"],
      phone: data["phone"],
      description: data["description"],
      problems: problems,
      conditions: data["conditions"],
    };
  }

  function saveToLocalStorage(data: InitialPatientFormData) {
    window.localStorage.setItem("pdata", JSON.stringify(data));
  }

  useEffect(() => {
    sockRef.current = new WebSocket(`ws://${BACKEND_URL}/talk`);
    const initialData = JSON.parse(window.localStorage.getItem("pdata"));
    if (initialData) {
      setForm(convertToFormData(initialData));
      setPage(2);
    } else {
      console.log("No save point found");
    }
    speak("Hello World");
  }, [sockRef, BACKEND_URL]);

  sockRef.current?.addEventListener("message", (event) => {
    console.log("Received message from server:", event.data);
    const parsed = JSON.parse(event.data);
    setThinking(false);
    setPage(2);
    setThoughts(parsed["thoughts"]);
    const nextQuestion = parsed["next_question"]
      .replace("/\n/g", "")
      .replace("/\t/g", "");
    setNext(nextQuestion);

    const parsedData = convertToFormData(parsed["info"]);
    setForm(parsedData);

    speak(next + thoughts);

    if (!parsed["next_question"] || parsed["success"] === "ok") {
      submit(parsedData);
    }
    saveToLocalStorage(parsedData);
  });

  sockRef.current?.addEventListener("open", () => {
    console.log("Connected to server");
  });

  sockRef.current?.addEventListener("close", () => {
    console.log("Disconnected from server");
  });

  const startRecording = async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      mediaRecorderRef.current = new MediaRecorder(stream, {
        mimeType: "audio/webm",
        audioBitsPerSecond: 16000,
      });
      console.log("Recording started...");

      mediaRecorderRef.current.ondataavailable = (event: BlobEvent) => {
        audioChunksRef.current.push(event.data as never);
      };

      mediaRecorderRef.current.onstop = async () => {
        const audioBlob = new Blob(audioChunksRef.current, {
          type: "audio/webm",
        });
        audioChunksRef.current = [];

        console.log("Sending audio to server...");
        setThinking(true);

        await sendAudioToServer(audioBlob).catch(console.error);
      };

      mediaRecorderRef.current.start();
      setIsRecording(true);
    } catch (error) {
      console.error("Error accessing microphone:", error);
    }
  };

  const stopRecording = () => {
    if (mediaRecorderRef.current) {
      console.log("Recording stopped...");
      mediaRecorderRef.current.stop();
    }
  };

  const sendAudioToServer = async (audioBlob: Blob) => {
    try {
      const formData = new FormData();
      formData.append("content", audioBlob, "recording.wav");

      const response = await fetch(`http://${BACKEND_URL}/transcribe`, {
        method: "POST",
        body: formData,
      });
      setIsRecording(false);

      if (response.ok) {
        const data = await response.json();
        console.log(data);
        setMessage(data.msg);
        console.log(data.msg);
        sockRef.current?.send(data.msg + " " + JSON.stringify(form));
      } else {
        setMessage("Failed to upload audio.");
      }
    } catch (error) {
      console.error("Error sending audio to server:", error);
      setMessage("An error occurred while uploading the audio.");
    }
  };

  async function submit(data: InitialPatientFormData) {
    console.log(data);
    const finalData: unknown = data;
    finalData["problems"] = JSON.stringify(finalData["problems"]);
    finalData["conditions"] = JSON.stringify(finalData["conditions"]);
    const res = await fetch(`http://${BACKEND_URL}/patients/create`, {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
      },
    });

    const result = await res.json();
    console.log("Patient created:", result);
    window.localStorage.removeItem("pdata");
    window.location.href = "/patients/" + result.id;
  }

  return (
    <div>
      {page === 0 && (
        <div>
          <div className="flex flex-col items-center justify-center w-full h-screen">
            <WavingEmoji emoji="👨🏻‍⚕️" />
            <h1 className="text-4xl font-bold mt-4">
              What{"'"}s troubling you today?
            </h1>
            <Button onClick={() => setPage(1)} className="text-2xl p-6 mt-4">
              Start Conversation
            </Button>
          </div>
        </div>
      )}
      {page === 1 && (
        <div>
          <div className="flex flex-col items-center justify-center w-full h-screen">
            {thinking ? (
              <ThinkingAnimation bg="#000000" fg="#ffffff">
                T
              </ThinkingAnimation>
            ) : (
              <div>
                {isRecording ? (
                  <div onClick={stopRecording}>
                    <ListeningAnimation
                      running={true}
                      bg="#000000"
                      fg="#ffffff"
                    >
                      <StopCircleIcon size={40} />
                    </ListeningAnimation>
                  </div>
                ) : (
                  <div onClick={startRecording}>
                    <ListeningAnimation
                      running={false}
                      bg="#000000"
                      fg="#ffffff"
                    >
                      <MicIcon size={40} />
                    </ListeningAnimation>
                  </div>
                )}
              </div>
            )}
          </div>
        </div>
      )}
      {page === 2 && (
        <div className="flex flex-row justify-center w-full h-screen">
          <div className="w-1/2 flex flex-col h-full">
            <div className="border-black p-4 h-1/2">
              <h1 className="text-3xl font-bold">Transcribed</h1>
              <p className="text-xl pt-3">{message}</p>
            </div>
            <div className="border-black p-4 h-1/2">
              <h1 className="text-3xl font-bold">Response</h1>
              <p className="text-xl pt-3">{thoughts}</p>
              <p className="text-xl pt-3">{next}</p>
              <Button
                onClick={() => {
                  setPage(1);
                  startRecording();
                }}
                className="text-2xl p-6 mt-4"
              >
                Respond
              </Button>
            </div>
          </div>
          <div className="w-1/2 border-black p-4">
            <InitialForm onSubmit={submit} initialData={form} />
          </div>
        </div>
      )}
    </div>
  );
};

export default Conversation;

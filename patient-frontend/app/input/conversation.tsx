"use client";

import React, { useState, useRef, useEffect } from "react";
import { motion } from "motion/react";
import { Button } from "@/components/ui/button";
import ListeningAnimation from "@/components/custom/talking_circle";
import { BACKEND_API_URL, BACKEND_HOST_URL } from "@/lib/consts";

const Conversation = () => {
  const [isRecording, setIsRecording] = useState(false);
  const [message, setMessage] = useState("");
  const [response, setResponse] = useState("");
  const mediaRecorderRef = useRef<MediaRecorder | null>(null);
  const audioChunksRef = useRef([]);
  const sockRef = useRef<WebSocket | null>(null);
  const [thinking, setThinking] = useState(false);

  useEffect(() => {
    sockRef.current = new WebSocket(`ws://${BACKEND_HOST_URL}/talk`);
  }, [sockRef]);

  sockRef.current?.addEventListener("message", (event) => {
    console.log("Received message from server:", event.data);
    const parsed = JSON.parse(event.data);
    setThinking(false);
    setResponse(parsed);
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

      const response = await fetch(`${BACKEND_API_URL}/transcribe`, {
        method: "POST",
        body: formData,
      });
      setIsRecording(false);

      if (response.ok) {
        const data = await response.json();
        console.log(data);
        setMessage(data.msg);
        console.log(data.msg);
        setThinking(true);
        sockRef.current?.send(data.msg);
      } else {
        setMessage("Failed to upload audio.");
      }
    } catch (error) {
      console.error("Error sending audio to server:", error);
      setMessage("An error occurred while uploading the audio.");
    }
  };

  function submit() {
    window.localStorage.setItem("response", JSON.stringify(response));

    window.location.href = "/form";
  }

  return (
    <div>
      {!isRecording ? (
        !message ? (
          <div>
            <motion.h1
              initial={{ scale: 0 }}
              animate={{
                scale: 1,
                transition: {
                  opacity: { ease: "linear" },
                },
              }}
              transition={{ bounce: 0.2, duration: 1.5 }}
              className="text-3xl font-bold text-center"
            >
              Please state your general information.
            </motion.h1>
            <motion.p
              initial={{ scale: 0 }}
              animate={{ scale: 1 }}
              transition={{ delay: 1 }}
              className="text-lg pt-2 text-center"
            >
              Include your name, address, and contact information.
            </motion.p>
            <motion.div
              initial={{ scale: 0 }}
              animate={{ scale: 1 }}
              transition={{ delay: 2 }}
              className="text-center mt-4"
            >
              <motion.div
                whileHover={{ scale: 1.1 }}
                whileTap={{ scale: 0.9 }}
                animate={{
                  y: [0, -4, 0, 0],
                }}
                transition={{
                  delay: 2,
                  duration: 1,
                  ease: "easeInOut",
                  repeat: Infinity,
                }}
              >
                <Button className="text-3xl p-8" onClick={startRecording}>
                  Start 🎤
                </Button>
              </motion.div>
            </motion.div>
          </div>
        ) : (
          <div className="text-center mt-4">
            <p className="text-xl">Captured Message, sending to the server.</p>
            <p>{message}</p>
          </div>
        )
      ) : (
        <div>
          <button className="" onClick={stopRecording}>
            <ListeningAnimation />
          </button>
        </div>
      )}
      {!thinking && !isRecording && response && (
        <div>
          <p>
            {response["next_question"] &&
              response["next_question"]
                .replace("/\\n/g", "")
                .replace("/\\t/g", "")}
          </p>
          <p>
            {response["success"] ||
              (response["next_question"] === "ok" && (
                <Button className="text-xl p-2" onClick={submit}>
                  Proceed
                </Button>
              ))}
          </p>
          <Button className="text-xl p-2" onClick={startRecording}>
            Respond 🎤
          </Button>
        </div>
      )}
    </div>
  );
};

export default Conversation;

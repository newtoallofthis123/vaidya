"use client";

import React, { useState, useRef } from "react";
import { motion } from "motion/react";
import { Button } from "@/components/ui/button";
import ListeningAnimation from "@/components/custom/talking_circle";
import { BACKEND_API_URL } from "@/lib/consts";

const AudioRecorder = () => {
  const [isRecording, setIsRecording] = useState(false);
  const [message, setMessage] = useState("");
  const mediaRecorderRef = useRef<MediaRecorder | null>(null);
  const audioChunksRef = useRef([]);

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

      if (response.ok) {
        const data = await response.json();
        setMessage(data.message);
        // store it in localstorage
        console.log(data);
        window.localStorage.setItem("description", data.message);
        window.localStorage.setItem("symptoms", JSON.stringify(data.symptoms));

        const parsed = await fetch(`${BACKEND_API_URL}/tokenize`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ text: data.message }),
        });

        if (parsed.ok) {
          const parsedData = await parsed.json();
          window.localStorage.setItem("info", JSON.stringify(parsedData));
        } else {
          console.log(parsed.body);
        }

        setIsRecording(false);
      } else {
        setMessage("Failed to upload audio.");
      }
    } catch (error) {
      console.error("Error sending audio to server:", error);
      setMessage("An error occurred while uploading the audio.");
    }
  };

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
            <p className="text-xl">Captured Message!</p>
            <Button className="mt-2 p-2">
              <a href="/form" className="text-xl">
                Continue
              </a>
            </Button>
          </div>
        )
      ) : (
        <div>
          <button className="" onClick={stopRecording}>
            <ListeningAnimation />
          </button>
        </div>
      )}
    </div>
  );
};

export default AudioRecorder;

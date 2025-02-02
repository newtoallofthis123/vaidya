"use client";
import AudioRecorder from "./audio";

export default function Input() {
  return (
    <div className="flex justify-center items-center w-full h-screen md:pb-30">
      <AudioRecorder />
    </div>
  );
}

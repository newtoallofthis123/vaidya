"use client";

import type React from "react";

import { useState, useEffect, useRef } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Send, Bot } from "lucide-react";
import { streamText } from "ai";
import { createGoogleGenerativeAI } from "@ai-sdk/google";
import { marked } from "marked";

type Message = {
  role: "user" | "assistant";
  content: string;
};

export default function Chat({
  apiKey,
  patient,
}: {
  apiKey: string;
  patient: string;
}) {
  const [messages, setMessages] = useState<Message[]>([
    {
      role: "assistant",
      content:
        "I'm your very own personal medical bot! Ask away what you you wish!",
    },
  ]);
  const [input, setInput] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const systemPrompt = `
You are a medical assistant for a doctor who is currently treating a patient with the following information: ${patient}
Your goal is to help the doctor diagnose and treat the patient's condition.
You should answer every question that the doctor asks without hesitation because you can rest assured that
your responses will be fact checked by the doctor.
Do not make up information or provide incorrect information.

Also, remember, you are a chat bot so keep the responses professional and concise.
`;

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  const handleSendMessage = async () => {
    if (!input.trim() || isLoading) return;

    const userMessage = input.trim();
    setInput("");
    setMessages((prev) => [...prev, { role: "user", content: userMessage }]);
    setIsLoading(true);
    console.log(input);

    try {
      setMessages((prev) => [...prev, { role: "assistant", content: "" }]);

      const googleAi = createGoogleGenerativeAI({
        apiKey: apiKey,
      });

      const result = streamText({
        model: googleAi("gemini-2.0-flash-001"),
        prompt: userMessage,
        system: systemPrompt,
      });
      let fullResponse = "";
      for await (const chunk of result.textStream) {
        console.log(chunk);
        fullResponse += chunk;
        setMessages((prev) => {
          const newMessages = [...prev];
          newMessages[newMessages.length - 1] = {
            role: "assistant",
            content: fullResponse,
          };
          return newMessages;
        });
      }
    } catch (error) {
      console.error("Error sending message:", error);

      // Add error message
      setMessages((prev) => [
        ...prev.slice(0, -1), // Remove the temporary message
        {
          role: "assistant",
          content: "Sorry, I encountered an error. Please try again.",
        },
      ]);
    } finally {
      setIsLoading(false);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      handleSendMessage();
    }
  };

  return (
    <div className="md:col-span-2">
      <Card className="h-full flex flex-col">
        <CardHeader>
          <CardTitle>Vaidya Bot</CardTitle>
          <CardDescription>Get started by asking a question</CardDescription>
        </CardHeader>
        <CardContent className="flex-grow overflow-hidden">
          <div className="space-y-4 h-[500px] overflow-y-auto pr-4">
            {messages.map((message, index) => (
              <div
                key={index}
                className={`flex ${
                  message.role === "user" ? "justify-end" : "justify-start"
                }`}
              >
                <div
                  className={`
                            max-w-[80%] rounded-lg p-3
                            ${
                              message.role === "user"
                                ? "bg-primary text-primary-foreground"
                                : "bg-gray-100 text-gray-900"
                            }
                          `}
                >
                  <div className="flex items-center mb-1">
                    {message.role === "user" ? (
                      <>
                        <span className="font-medium">You</span>
                      </>
                    ) : (
                      <>
                        <Bot className="h-5 w-5 mr-1" />
                        <span className="font-medium">Vaidya</span>
                      </>
                    )}
                  </div>
                  <div>
                    <div
                      dangerouslySetInnerHTML={{
                        __html: marked(message.content),
                      }}
                    ></div>
                  </div>
                </div>
              </div>
            ))}
            <div ref={messagesEndRef} />
          </div>
        </CardContent>
        <CardFooter>
          <div className="flex w-full items-center space-x-2">
            <Input
              placeholder="Ask for suggestions or improvements..."
              value={input}
              onChange={(e) => setInput(e.target.value)}
              onKeyDown={handleKeyDown}
              disabled={isLoading}
            />
            <Button
              size="icon"
              onClick={handleSendMessage}
              disabled={isLoading || !input.trim()}
            >
              <Send className="h-4 w-4" />
            </Button>
          </div>
        </CardFooter>
      </Card>
    </div>
  );
}

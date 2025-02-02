"use client";

import WavingHand from "@/components/custom/wavinghand";
import { Button } from "@/components/ui/button";
import { motion } from "motion/react";

export default function Home() {
  return (
    <main>
      <div className="flex justify-center items-center w-full h-screen md:pb-30">
        <div className="rounded-lg border-gray-300 backdrop-blur-lg">
          <form>
            <motion.h1
              initial={{ scale: 0, y: 20 }}
              animate={{
                scale: 1,
                transition: {
                  opacity: { ease: "linear" },
                },
                y: -25,
              }}
              transition={{ bounce: 0.2, duration: 1.5 }}
              className="text-3xl font-bold text-center"
            >
              Welcome <WavingHand />
            </motion.h1>
            <motion.p
              initial={{ scale: 0 }}
              animate={{ scale: 1 }}
              transition={{ delay: 1 }}
              className="text-xl"
            >
              Start your preliminary examination with a few simple questions
            </motion.p>
            <motion.div
              initial={{ scale: 0 }}
              animate={{ scale: 1 }}
              transition={{ delay: 2 }}
              className="text-center mt-10"
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
                <Button className="text-4xl p-8">
                  <a href="/input"> Let{"'"}s Go </a>
                </Button>
              </motion.div>
            </motion.div>
          </form>
        </div>
      </div>
    </main>
  );
}

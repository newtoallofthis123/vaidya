import { motion } from "motion/react";

export default function TalkingAnimation() {
  return (
    <motion.div
      initial={{ scale: 1 }}
      animate={{ scale: [1, 1.1, 0.9, 1] }}
      transition={{
        duration: 1,
        repeat: Infinity,
        repeatType: "loop",
        ease: "easeInOut",
      }}
      style={ball}
    />
  );
}

const ball = {
  width: 100,
  height: 100,
  backgroundColor: "#1700ee",
  borderRadius: "50%",
};

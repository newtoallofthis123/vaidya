import { motion } from "framer-motion";

export default function WavingHand() {
  return (
    <motion.div
      style={{
        display: "inline-block",
        fontSize: "3rem",
      }}
      animate={{
        rotate: [0, 20, -10, 20, -10, 0],
      }}
      transition={{
        duration: 1.5,
        ease: "easeInOut",
        repeat: Infinity,
      }}
    >
      👋🏻
    </motion.div>
  );
}

import { motion } from "motion/react";

export default function ThinkingAnimation({
  bg = "#1700ee",
  fg = "white",
  children,
}: {
  bg: string | null;
  fg: string | null;
  children: React.ReactNode;
}) {
  const ball = {
    width: 100,
    height: 100,
    backgroundColor: bg,
    borderRadius: "50%",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    color: fg,
    fontSize: "1.5em",
  };
  return (
    <motion.div
      initial={{ y: 0 }}
      animate={{
        y: [0, -10, 0, 10, 0],
        rotate: [0, 2, 0, -2, 0], // Subtle rotation
      }}
      transition={{
        duration: 2,
        repeat: Infinity,
        repeatType: "loop",
        ease: "easeInOut",
      }}
      style={ball}
    >
      {children}
    </motion.div>
  );
}

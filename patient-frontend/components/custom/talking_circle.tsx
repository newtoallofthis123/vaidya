import { motion } from "motion/react";

export default function TalkingAnimation({
  bg = "#1700ee",
  fg = "white",
  running = true,
  children,
}: {
  bg: string | null;
  fg: string | null;
  running: boolean;
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
    cursor: "pointer",
  };
  return (
    <motion.div
      initial={{ scale: 1 }}
      animate={{ scale: running ? [1, 1.1, 0.9, 1] : 1 }}
      transition={{
        duration: 1,
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

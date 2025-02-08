"use client";

import EnterAnimation from "@/components/custom/talking_circle";
import { Button } from "@/components/ui/button";
import React from "react";

function Conversation() {
  const [userTurn, setUserTurn] = React.useState(true);

  return (
    <div>
      {userTurn ? (
        <div>
          User Turn
          <Button onClick={() => setUserTurn(false)}>Next</Button>
        </div>
      ) : (
        <div>
          <p>AI Turn</p>
          <EnterAnimation />
        </div>
      )}
    </div>
  );
}

export default Conversation;

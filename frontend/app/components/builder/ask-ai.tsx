import { PaperPlaneTilt, Sparkle } from "@phosphor-icons/react";
import React from "react";
import { Input } from "../ui/input";
import { Button } from "../ui/button";

export default function AskB0() {
  return (
    <div className="h-8 border border-input bg-background shadow-lg flex">
      <div className="h-8 border-r border-input flex items-center px-2">
        <Sparkle size={20} className="h-4 w-4" />
      </div>
      <Input
        className="h-8 border-none focus-visible:ring-0 w-72 px-2"
        placeholder="Ask b0 anything..."
      />
      <Button
        className="!h-8 border-l border-t-0 border-r-0 border-b-0 px-2 bg-transparent"
        variant={"outline"}
      >
        <PaperPlaneTilt className="h-4 w-4" />
      </Button>
    </div>
  );
}

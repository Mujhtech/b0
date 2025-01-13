import React from "react";
import { usePlayground } from "./playground/provider";
import { cn } from "~/lib/utils";

export default function BuilderMenu() {
  const {} = usePlayground();
  return (
    <div className="absolute left-4 top-4 z-10">
      <div
        className={cn(
          "bg-background shadow-lg border border-input w-auto select-none max-h-[calc(100%-60px)]"
        )}
      >
        <div className={cn("flex items-center justify-between px-2 py-1")}>
          <h1 className="font-black text-md">b0</h1>
        </div>
      </div>
    </div>
  );
}

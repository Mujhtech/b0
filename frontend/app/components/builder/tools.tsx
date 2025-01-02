import { ArrowsIn } from "@phosphor-icons/react";
import React from "react";

export default function BuilderTools() {
  return (
    <div className="absolute right-4 top-4">
      <div className="bg-background shadow-lg border border-input w-[24rem] select-none max-h-[calc(100%-60px)]">
        <div className="border-b border-input flex items-center justify-between px-2 py-1">
          <h3 className="text-md font-mono">Tools</h3>
          <button className="bg-background h-5 w-5 border border-input inline-flex items-center justify-center shadow-sm">
            <ArrowsIn size={28} className="h-4 w-4" />
          </button>
        </div>
        <div className="p-2"></div>
      </div>
    </div>
  );
}

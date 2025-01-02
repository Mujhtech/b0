import { ArrowsIn, Minus, Plus, ArrowsOut } from "@phosphor-icons/react";
import React from "react";
import { Input } from "../ui/input";

export default function ZoomInAndOut() {
  return (
    <div className="flex items-center gap-2">
      <div className="flex border border-input h-8 bg-background shadow-lg">
        <button className="px-2 border-r border-input">
          <Plus size={28} className="h-4 w-4" />
        </button>
        <Input
          className="h-8 border-none focus-visible:ring-0 w-10 max-w-10 px-2 items-center justify-center text-center"
          defaultValue={20}
        />
        <button className="px-2 border-l border-input">
          <Minus size={28} className="h-4 w-4" />
        </button>
      </div>
      <button className="px-2 bg-background h-8 w-8 border border-input shadow-lg">
        <ArrowsIn size={28} className="h-4 w-4" />
      </button>
    </div>
  );
}

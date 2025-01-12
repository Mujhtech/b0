import React from "react";
import { Dialog, DialogContent, DialogTrigger } from "~/components/ui/dialog";

export default function Connector() {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <div className="border border-input w-[250px] bg-background shadow-sm flex flex-col hover:drop-shadow-xl cursor-pointer">
          <div className="border-b border-input flex justify-between">
            <h3 className="text-xs font-mono text-muted-foreground p-2">
              HTTP Request
            </h3>
          </div>
          <div className="p-2"></div>
        </div>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]"></DialogContent>
    </Dialog>
  );
}

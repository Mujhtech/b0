import React from "react";
import { usePlayground } from "./playground/provider";
import { cn } from "~/lib/utils";
import Paragraph from "../ui/paragraph";
import { LogPreviewDialog } from "./log-preview-dialog";

export default function BuilderMenu() {
  const {} = usePlayground();
  const [openLog, setOpenLog] = React.useState(false);

  const handleOpenLog = () => {
    setOpenLog(!openLog);
  };
  return (
    <>
      <div className="absolute left-4 top-4 z-10">
        <div
          className={cn(
            "bg-background shadow-lg border flex border-input w-auto select-none max-h-[calc(100%-60px)]"
          )}
        >
          <div
            className={cn(
              "flex items-center justify-between px-2 py-1 border-r border-input"
            )}
          >
            <h1 className="font-black text-md">b0</h1>
          </div>
          <div className="mx-2">
            <div className="h-full flex items-center justify-center gap-2">
              <Paragraph onClick={handleOpenLog}>Logs</Paragraph>
              <Paragraph>Logs</Paragraph>
            </div>
          </div>
        </div>
      </div>
      <LogPreviewDialog open={openLog} setOpen={setOpenLog} />
    </>
  );
}

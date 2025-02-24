import React from "react";
import { usePlayground } from "./playground/provider";
import { cn } from "~/lib/utils";
import Paragraph from "../ui/paragraph";
import { usePlaygroundBuilder } from "./provider";
import { useNavigate } from "@remix-run/react";
import { useOptionalProject } from "~/hooks/use-project";

export default function BuilderMenu() {
  const {} = usePlayground();
  const { setOpenLogPreviewDialog, openLogPreviewDialog } =
    usePlaygroundBuilder();

  const navigate = useNavigate();

  const project = useOptionalProject();

  const handleOpenLog = () => {
    setOpenLogPreviewDialog(!openLogPreviewDialog);
  };

  const handleGotToSettings = () => {
    navigate(`/${project?.slug}/settings`);
  };
  return (
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
          <div className="h-full flex items-center justify-center gap-2.5">
            <Paragraph
              className="cursor-pointer font-bold"
              onClick={handleOpenLog}
            >
              Logs
            </Paragraph>
            <Paragraph
              className="cursor-pointer font-bold"
              onClick={handleGotToSettings}
            >
              Settings
            </Paragraph>
          </div>
        </div>
      </div>
    </div>
  );
}

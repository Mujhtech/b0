import { ArrowUpRight } from "@phosphor-icons/react";
import { Link } from "@remix-run/react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "~/components/ui/dialog";
import { CodeBlock } from "../codeblock";
import { usePlaygroundBuilder } from "./provider";

export function LogPreviewDialog() {
  const { logs, openLogPreviewDialog, setOpenLogPreviewDialog } =
    usePlaygroundBuilder();
  return (
    <Dialog open={openLogPreviewDialog} onOpenChange={setOpenLogPreviewDialog}>
      <DialogContent className="sm:max-w-[90%] h-[90%] flex flex-col p-0">
        <DialogHeader className="hidden">
          <DialogTitle>Log</DialogTitle>
          <DialogDescription className="hidden">
            View and manage your projects
          </DialogDescription>
        </DialogHeader>
        <CodeBlock
          fileName="logs.json"
          showChrome={true}
          showCopyButton={false}
          code={logs.join("\n")}
          language="json"
          className="h-full border-none"
          maxLines={1000}
        />
      </DialogContent>
    </Dialog>
  );
}

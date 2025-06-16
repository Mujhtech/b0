import {
  AlignCenterHorizontal,
  ArrowsIn,
  ArrowsOut,
  Code,
  DiscordLogo,
  GithubLogo,
  OpenAiLogo,
  Repeat,
  SignOut,
  SlackLogo,
  StripeLogo,
  TelegramLogo,
  TreeStructure,
  X,
} from "@phosphor-icons/react";
import React from "react";
import { usePlayground } from "./playground/provider";
import { cn } from "~/lib/utils";
import { useSortable } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";
import { Icons } from "../icons";

export const TOOLS = {
  "tool-if": {
    name: "IF",
    icon: <TreeStructure className="mr-1" />,
  },
  "tool-loop": {
    name: "Loop",
    icon: <Repeat className="mr-1" />,
  },
  "tool-switch": {
    name: "Switch",
    icon: <AlignCenterHorizontal className="mr-1" />,
  },
  "tool-variable": {
    name: "Variable",
    icon: <X className="mr-1" />,
  },
  "tool-output": {
    name: "Output",
    icon: <SignOut className="mr-1" />,
  },
  "tool-codeblock": {
    name: "Codeblock",
    icon: <Code className="mr-1" />,
  },
  "tool-openai": {
    name: "OpenAI",
    icon: <OpenAiLogo className="mr-1" />,
  },
  "tool-stripe": {
    name: "Stripe",
    icon: <StripeLogo className="mr-1" />,
  },
  "tool-github": {
    name: "GitHub",
    icon: <GithubLogo className="mr-1" />,
  },
  "tool-slack": {
    name: "Slack",
    icon: <SlackLogo className="mr-1" />,
  },
  "tool-telegram": {
    name: "Telegram",
    icon: <TelegramLogo className="mr-1" />,
  },
  "tool-discord": {
    name: "Discord",
    icon: <DiscordLogo className="mr-1" />,
  },
  "tool-supabase": {
    name: "Supabase",
    icon: <Icons.Supabase className="mr-1 h-3 w-3 grayscale" />,
  },
  "tool-resend": {
    name: "Resend",
    icon: <Icons.Resend className="mr-1 h-3 w-3 grayscale" />,
  },
};

export default function BuilderTools() {
  const { handleExpandToolPanel, expandToolPanel } = usePlayground();

  return (
    <div className="absolute right-4 top-4 z-10">
      <div
        className={cn(
          "bg-background shadow-lg border border-input w-[24rem] select-none max-h-[calc(100%-60px)]",
          expandToolPanel == true && ""
        )}
      >
        <div
          className={cn(
            "flex items-center justify-between px-2 py-1",
            expandToolPanel == true && "border-b border-input "
          )}
        >
          <h3 className="text-sm font-mono text-muted-foreground">Tools</h3>
          <button
            onClick={handleExpandToolPanel}
            className="bg-background h-5 w-5 border border-input inline-flex items-center justify-center shadow-sm"
          >
            {expandToolPanel ? (
              <ArrowsIn size={28} className="h-4 w-4" />
            ) : (
              <ArrowsOut size={28} className="h-4 w-4" />
            )}
          </button>
        </div>
        {expandToolPanel && (
          <div className="p-2 flex flex-col gap-2">
            <div className="flex flex-col">
              {/* <h3 className="text-muted-foreground text-xs font-medium font-mono mb-1">
                Flow
              </h3> */}
              <div className="grid grid-cols-2 gap-2">
                {Object.entries(TOOLS).map(([id, tool]) => (
                  <ToolCard key={id} id={id}>
                    {tool.icon} {tool.name}
                  </ToolCard>
                ))}
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export const ToolCard = ({
  children,
  id,
}: {
  children: React.ReactNode;
  id: string;
}) => {
  const { attributes, listeners, setNodeRef, transform, transition } =
    useSortable({ id });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  return (
    <div
      ref={setNodeRef}
      style={style}
      {...attributes}
      {...listeners}
      className="cursor-grab border border-input px-2 py-1 flex items-center text-sm bg-background drop-shadow-md"
    >
      {children}
    </div>
  );
};

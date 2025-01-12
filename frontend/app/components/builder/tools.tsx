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

export default function BuilderTools() {
  const { handleExpandToolPanel, expandToolPanel } = usePlayground();

  return (
    <div className="absolute right-4 top-4">
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
              <h3 className="text-muted-foreground text-xs font-medium font-mono mb-1">
                Flow
              </h3>
              <div className="grid grid-cols-2 gap-2">
                <ToolCard>
                  <TreeStructure className="mr-1" /> IF
                </ToolCard>
                <ToolCard>
                  <Repeat className="mr-1" /> Loop
                </ToolCard>
                <ToolCard>
                  <AlignCenterHorizontal className="mr-1" /> Switch
                </ToolCard>
                <ToolCard>
                  <X className="mr-1" /> Variable
                </ToolCard>
                <ToolCard>
                  <SignOut className="mr-1" /> Output
                </ToolCard>
                <ToolCard>
                  <Code className="mr-1" /> Codeblock
                </ToolCard>
              </div>
            </div>
            <div className="flex flex-col">
              <h3 className="text-muted-foreground text-xs font-medium font-mono mb-1">
                Integration
              </h3>
              <div className="grid grid-cols-2 gap-2">
                <ToolCard>
                  <OpenAiLogo className="mr-1" /> OpenAI
                </ToolCard>

                <ToolCard>
                  <StripeLogo className="mr-1" /> Stripe
                </ToolCard>
                <ToolCard>
                  <GithubLogo className="mr-1" /> GitHub
                </ToolCard>

                <ToolCard>
                  <SlackLogo className="mr-1" /> Slack
                </ToolCard>
                <ToolCard>
                  <TelegramLogo className="mr-1" /> Telegram
                </ToolCard>
                <ToolCard>
                  <DiscordLogo className="mr-1" /> Discord
                </ToolCard>
                <ToolCard> Supabase</ToolCard>
                <ToolCard>Resend</ToolCard>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

const ToolCard = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="cursor-grab border border-input px-2 py-1 flex items-center text-sm bg-background drop-shadow-md">
      {children}
    </div>
  );
};

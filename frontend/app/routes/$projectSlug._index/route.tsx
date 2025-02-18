import AskB0 from "~/components/builder/ask-ai";
import DeployAndTestBtn from "~/components/builder/deploy";
import BuilderTools from "~/components/builder/tools";
import ZoomInAndOut from "~/components/builder/zoom";
import UserMenu from "~/components/menus/user-menu";
import { PlaygroundProvider } from "~/components/builder/playground/provider";
import Playground from "~/components/builder/playground/playground";
import BuilderMenu from "~/components/builder/builder-menu";
import {
  PlaygroundBuilderProvider,
  usePlaygroundBuilder,
} from "~/components/builder/provider";
import { useSSE } from "~/hooks/use-ses";
import { useAuthToken, usePlatformUrl, useUser } from "~/hooks/use-user";
import { useProject } from "~/hooks/use-project";
import { useCallback, useMemo, useState } from "react";
import { Spinner } from "@phosphor-icons/react";
import { useNavigate } from "@remix-run/react";
import { LogPreviewDialog } from "~/components/builder/log-preview-dialog";

export default function Page() {
  return (
    <PlaygroundBuilderProvider>
      <PlaygroundProvider>
        <PageContent />
      </PlaygroundProvider>
    </PlaygroundBuilderProvider>
  );
}

const PageContent = () => {
  const project = useProject();
  const { isThinking, contextMessage, error, isSaving } =
    usePlaygroundBuilder();
  const user = useUser();

  const messageToDisplay = useMemo(() => {
    if (contextMessage != undefined) {
      return contextMessage;
    }

    if (error != undefined) {
      return error;
    }

    if (isThinking) {
      return "Thinking...";
    }

    if (isSaving) {
      return "Saving...";
    }

    return undefined;
  }, [contextMessage, error, isThinking, isSaving]);

  const showLoaderIcon = useMemo(() => {
    return isThinking || isSaving;
  }, [isSaving, isThinking]);

  return (
    <main className="h-full w-full relative canvas-bg">
      <Playground />

      <div className="absolute bottom-4 w-full z-10">
        {messageToDisplay != undefined && (
          <div className="flex items-center justify-center gap-2 mb-1">
            {showLoaderIcon && (
              <Spinner
                size={18}
                className="text-muted-foreground animate-spin"
              />
            )}
            <p className="font-mono text-[10px] text-muted-foreground text-center">
              {messageToDisplay}
            </p>
          </div>
        )}
        <div className="flex justify-between">
          <div className=" ml-4">
            <div className="flex items-center gap-2">
              <UserMenu user={user} showHomepage={true} />
              <ZoomInAndOut />
            </div>
          </div>
          <div className="flex flex-col gap-0.5">
            <AskB0 isThinking={isThinking} projectModel={project.model} />
            <div className="flex items-center justify-center">
              <p className="font-mono text-[10px] text-muted-foreground text-center">
                b0 can make mistakes. Please double-check it.
              </p>
            </div>
          </div>
          <div className="mr-4">
            <DeployAndTestBtn isThinking={isThinking} />
          </div>
        </div>
      </div>
      <BuilderMenu />
      <BuilderTools />
      <LogPreviewDialog />
    </main>
  );
};

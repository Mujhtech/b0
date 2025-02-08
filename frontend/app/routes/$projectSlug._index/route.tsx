import AskB0 from "~/components/builder/ask-ai";
import DeployAndTestBtn from "~/components/builder/deploy";
import BuilderTools from "~/components/builder/tools";
import ZoomInAndOut from "~/components/builder/zoom";
import UserMenu from "~/components/menus/user-menu";
import { PlaygroundProvider } from "~/components/builder/playground/provider";
import Playground from "~/components/builder/playground/playground";
import BuilderMenu from "~/components/builder/builder-menu";
import { PlaygroundBuilderProvider } from "~/components/builder/provider";
import { useSSE } from "~/hooks/use-ses";
import { useAuthToken, usePlatformUrl } from "~/hooks/use-user";
import { useProject } from "~/hooks/use-project";
import { useCallback, useMemo, useState } from "react";
import { Spinner } from "@phosphor-icons/react";

export default function Page() {
  const platformUrl = usePlatformUrl();
  const project = useProject();
  const accessToken = useAuthToken();
  const DEFAULT_EVENTS = [
    "task_started",
    "task_updated",
    "task_completed",
    "task_failed",
  ];
  const [isThinking, setIsThinking] = useState(false);

  useSSE({
    baseUrl: `${platformUrl}/projects/${project.id}/sse`,
    shouldRun: true,
    projectId: project.id,
    accessToken: accessToken ?? "",
    events: useMemo(() => DEFAULT_EVENTS, []),
    onEvent: useCallback((type: string, data: any) => {
      const { message, workflows, error } = JSON.parse(data);

      console.log("TYPE", type);
      console.log("MESSAGE", message);
      if (error != undefined) {
        console.log("ERROR", error);
      }

      if (type === "task_started") {
        setIsThinking(true);
      }

      if (type === "task_completed" || type === "task_failed") {
        setIsThinking(false);
      }

      if (type === "task_updated" && workflows != undefined) {
        console.log("WORKFLOWS", workflows);
      }
    }, []),
  });

  return (
    <PlaygroundBuilderProvider>
      <PlaygroundProvider>
        <main className="h-full w-full relative canvas-bg">
          <Playground />

          <div className="absolute bottom-4 w-full z-10">
            {isThinking && (
              <div className="flex items-center justify-center gap-2 mb-1">
                <Spinner
                  size={18}
                  className="text-muted-foreground animate-spin"
                />
                <p className="font-mono text-[10px] text-muted-foreground text-center">
                  b0 is thinking...
                </p>
              </div>
            )}
            <div className="flex justify-between">
              <div className=" ml-4">
                <div className="flex items-center gap-2">
                  <UserMenu />
                  <ZoomInAndOut />
                </div>
              </div>
              <div className="flex flex-col gap-0.5">
                <AskB0 isThinking={isThinking} />
                <div className="flex items-center justify-center">
                  <p className="font-mono text-[10px] text-muted-foreground text-center">
                    b0 can make mistakes. Please double-check it.
                  </p>
                </div>
              </div>
              <div className="mr-4">
                <DeployAndTestBtn />
              </div>
            </div>
          </div>
          <BuilderMenu />
          <BuilderTools />
        </main>
      </PlaygroundProvider>
    </PlaygroundBuilderProvider>
  );
}

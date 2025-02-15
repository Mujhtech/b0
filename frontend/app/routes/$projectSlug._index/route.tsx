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
import { useAuthToken, usePlatformUrl, useUser } from "~/hooks/use-user";
import { useProject } from "~/hooks/use-project";
import { useCallback, useMemo, useState } from "react";
import { Spinner } from "@phosphor-icons/react";
import { useNavigate } from "@remix-run/react";

export default function Page() {
  const platformUrl = usePlatformUrl();
  const project = useProject();
  const accessToken = useAuthToken();
  const DEFAULT_EVENTS = [
    "task_started",
    "task_updated",
    "task_failed",
    "task_completed",
  ];
  const [isThinking, setIsThinking] = useState(false);
  const [error, setError] = useState<string | undefined>(undefined);
  const [contextMessage, setContextMessage] = useState<string | undefined>(
    undefined
  );
  const navigate = useNavigate();
  const user = useUser();

  useSSE({
    baseUrl: `${platformUrl}/projects/${project.id}/sse`,
    shouldRun: true,
    projectId: project.id,
    accessToken: accessToken ?? "",
    events: useMemo(() => DEFAULT_EVENTS, []),
    onEvent: useCallback(
      (type: string, data: any) => {
        const { message, workflows, code, error, should_reload_window } =
          JSON.parse(data);

        if (code != undefined) {
          console.log("CODE GENERATION", code);
        }

        if (type === "task_started") {
          setError(undefined);
          setIsThinking(true);
        }

        if (error != undefined) {
          setContextMessage(undefined);
          setError(error);
        }

        if (error == undefined && type === "task_updated") {
          setContextMessage(message);
        }

        if (type === "task_completed" || type === "task_failed") {
          setIsThinking(false);

          if (type === "task_completed") {
            setContextMessage(undefined);
          }
        }

        if (should_reload_window && should_reload_window == true) {
          setTimeout(() => {
            navigate(`/${project.slug}`);
          }, 1500);
        }
      },
      [
        contextMessage,
        isThinking,
        error,
        setContextMessage,
        setIsThinking,
        setError,
      ]
    ),
  });

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

    return undefined;
  }, [contextMessage, error, isThinking]);

  return (
    <PlaygroundBuilderProvider>
      <PlaygroundProvider>
        <main className="h-full w-full relative canvas-bg">
          <Playground />

          <div className="absolute bottom-4 w-full z-10">
            {messageToDisplay != undefined && (
              <div className="flex items-center justify-center gap-2 mb-1">
                {isThinking && (
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
                  <UserMenu user={user} />
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

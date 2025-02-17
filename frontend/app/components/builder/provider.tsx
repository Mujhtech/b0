import { useNavigate } from "@remix-run/react";
import React, {
  createContext,
  ReactNode,
  useCallback,
  useContext,
  useMemo,
  useState,
} from "react";
import { useOptionalProject, useProject } from "~/hooks/use-project";
import { useSSE } from "~/hooks/use-ses";
import { useAuthToken, usePlatformUrl } from "~/hooks/use-user";
import { ServerResponse, ServerResponseSchema } from "~/models/default";
import { EndpointWorkflow } from "~/models/endpoint";
import { Project } from "~/models/project";
import { clientApi } from "~/services/api.client";

interface PlaygroundBuilderProviderProps {
  project?: Project;
  handleProjectAction: (action: string) => Promise<ServerResponse>;
  handleUpdateEndpointWorkflow: (
    id: string,
    workflows: EndpointWorkflow[]
  ) => Promise<ServerResponse>;
  defaultAction: string;
  setDefaultAction: React.Dispatch<React.SetStateAction<string>>;
  isThinking: boolean;
  error: string | undefined;
  contextMessage: string | undefined;
  logs: string[];
  isSaving: boolean;
  setIsSaving: React.Dispatch<React.SetStateAction<boolean>>;
}

const PlaygroundBuilderProviderContext = createContext<
  PlaygroundBuilderProviderProps | undefined
>(undefined);

export const PlaygroundBuilderProvider = ({
  children,
}: {
  children: ReactNode;
}) => {
  const accessToken = useAuthToken();
  const backendBaseUrl = usePlatformUrl();
  const project = useProject();
  const [defaultAction, setDefaultAction] = useState<string>("deploy");
  const DEFAULT_EVENTS = [
    "task_started",
    "task_updated",
    "task_failed",
    "task_completed",
  ];
  const [isThinking, setIsThinking] = useState(false);
  const [isSaving, setIsSaving] = useState(false);
  const [error, setError] = useState<string | undefined>(undefined);
  const [contextMessage, setContextMessage] = useState<string | undefined>(
    undefined
  );
  const [logs, setLogs] = useState<string[]>([]);
  const navigate = useNavigate();

  async function handleProjectAction(action: string) {
    return await clientApi.post<ServerResponse>({
      path: `/projects/${project.id}/action`,
      body: {
        action,
      },
      backendBaseUrl: backendBaseUrl!,
      accessToken: accessToken!,
      schema: ServerResponseSchema,
    });
  }

  async function handleUpdateEndpointWorkflow(
    id: string,
    workflows: EndpointWorkflow[]
  ) {
    return await clientApi.put<ServerResponse>({
      path: `/endpoints/${id}/workflows`,
      body: {
        workflows,
      },
      backendBaseUrl: backendBaseUrl!,
      accessToken: accessToken!,
      schema: ServerResponseSchema,
    });
  }

  useSSE({
    baseUrl: `${backendBaseUrl}/projects/${project.id}/sse`,
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

  useSSE({
    baseUrl: `${backendBaseUrl}/projects/${project.id}/log`,
    shouldRun: true,
    projectId: project.id,
    accessToken: accessToken ?? "",
    events: useMemo(() => ["log_started", "log_updated"], []),
    onEvent: useCallback(
      (type: string, data: any) => {
        const { message, log, error } = JSON.parse(data);
        console.log("LOG", message, log, error);

        if (log != undefined) {
          setLogs((prevLogs) => [...prevLogs, log]);
        }
      },
      [setLogs]
    ),
  });

  const value = useMemo<PlaygroundBuilderProviderProps>(
    () => ({
      project,
      handleProjectAction,
      handleUpdateEndpointWorkflow,
      defaultAction,
      setDefaultAction,
      isThinking,
      error,
      contextMessage,
      logs,
      isSaving,
      setIsSaving,
    }),
    [
      project,
      handleProjectAction,
      defaultAction,
      setDefaultAction,
      isThinking,
      error,
      contextMessage,
      logs,
      setLogs,
      isSaving,
      setIsSaving,
    ]
  );

  return (
    <PlaygroundBuilderProviderContext.Provider value={value}>
      {children}
    </PlaygroundBuilderProviderContext.Provider>
  );
};

export const usePlaygroundBuilder = () => {
  const context = useContext(PlaygroundBuilderProviderContext);
  if (!context) {
    throw new Error(
      "usePlaygroundBuilder must be used within an PlaygroundBuilderProviderContext"
    );
  }
  return context;
};

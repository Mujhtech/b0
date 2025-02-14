import React, {
  createContext,
  ReactNode,
  useContext,
  useMemo,
  useState,
} from "react";
import { useOptionalProject, useProject } from "~/hooks/use-project";
import { useAuthToken, usePlatformUrl } from "~/hooks/use-user";
import { ServerResponse, ServerResponseSchema } from "~/models/default";
import { Project } from "~/models/project";
import { clientApi } from "~/services/api.client";

interface PlaygroundBuilderProviderProps {
  project?: Project;
  handleProjectAction: (action: string) => Promise<ServerResponse>;
  defaultAction: string;
  setDefaultAction: React.Dispatch<React.SetStateAction<string>>;
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

  const value = useMemo<PlaygroundBuilderProviderProps>(
    () => ({ project, handleProjectAction, defaultAction, setDefaultAction }),
    [project, handleProjectAction, defaultAction, setDefaultAction]
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

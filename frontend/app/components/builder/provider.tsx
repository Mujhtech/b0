import { createContext, ReactNode, useContext, useMemo } from "react";
import { useOptionalProject } from "~/hooks/use-project";
import { Project } from "~/models/project";

interface PlaygroundBuilderProviderProps {
  project?: Project;
}

const PlaygroundBuilderProviderContext = createContext<
  PlaygroundBuilderProviderProps | undefined
>(undefined);

export const PlaygroundBuilderProvider = ({
  children,
}: {
  children: ReactNode;
}) => {
  const project = useOptionalProject();
  const value = useMemo<PlaygroundBuilderProviderProps>(
    () => ({ project }),
    [project]
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

import { UIMatch } from "@remix-run/react";
import type { Project } from "~/models/project";
import { useTypedMatchesData } from "./use-typed-match";
import { loader } from "~/routes/$projectSlug._index/route";
import { useChanged } from "./use-changed";
import { UseDataFunctionReturn } from "remix-typedjson";

export type MatchedProject = UseDataFunctionReturn<typeof loader>["project"];

export function useOptionalProject(matches?: UIMatch[]): Project | undefined {
  const routeMatch = useTypedMatchesData<typeof loader>({
    id: "root",
    matches,
  });

  return routeMatch?.project ?? undefined;
}

export function useProject(matches?: UIMatch[]): Project {
  const maybeUser = useOptionalProject(matches);
  if (!maybeUser) {
    throw new Error("No project found in loader.");
  }
  return maybeUser;
}

export const useProjectChanged = (
  action: (project: MatchedProject | undefined) => void
) => {
  useChanged(useOptionalProject, action);
};

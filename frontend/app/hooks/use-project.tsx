import { UIMatch } from "@remix-run/react";
import type { Project, Projects } from "~/models/project";
import { useTypedMatchesData } from "./use-typed-match";
import { loader } from "~/routes/$projectSlug/route";
import { useChanged } from "./use-changed";
import { UseDataFunctionReturn } from "remix-typedjson";
import { Endpoint, Endpoints } from "~/models/endpoint";

export type MatchedProject = UseDataFunctionReturn<typeof loader>["project"];

export function useOptionalProjects(matches?: UIMatch[]): Projects | undefined {
  const routeMatch = useTypedMatchesData<typeof loader>({
    id: "routes/$projectSlug",
    matches,
  });

  return routeMatch?.projects ?? undefined;
}

export function useProjects(matches?: UIMatch[]): Projects {
  const maybe = useOptionalProjects(matches);
  if (!maybe) {
    throw new Error("No projects found in loader.");
  }
  return maybe;
}

export function useOptionalProject(matches?: UIMatch[]): Project | undefined {
  const routeMatch = useTypedMatchesData<typeof loader>({
    id: "routes/$projectSlug",
    matches,
  });

  return routeMatch?.project ?? undefined;
}

export function useProject(matches?: UIMatch[]): Project {
  const maybe = useOptionalProject(matches);
  if (!maybe) {
    throw new Error("No project found in loader.");
  }
  return maybe;
}

export function useOptionalEndpoints(
  matches?: UIMatch[]
): Endpoints | undefined {
  const routeMatch = useTypedMatchesData<typeof loader>({
    id: "routes/$projectSlug",
    matches,
  });

  return routeMatch?.endpoints;
}

export function useEndpoints(matches?: UIMatch[]): Endpoints {
  const maybe = useOptionalEndpoints(matches);
  if (!maybe) {
    throw new Error("No endpoints found in loader.");
  }
  return maybe;
}

export function useOptionalEndpoint(matches?: UIMatch[]): Endpoint | undefined {
  const routeMatch = useTypedMatchesData<typeof loader>({
    id: "routes/$projectSlug",
    matches,
  });

  return routeMatch?.endpoint;
}

export function useEndpoint(matches?: UIMatch[]): Endpoint {
  const maybe = useOptionalEndpoint(matches);
  if (!maybe) {
    throw new Error("No endpoint found in loader.");
  }
  return maybe;
}

export const useProjectChanged = (
  action: (project: MatchedProject | undefined) => void
) => {
  useChanged(useOptionalProject, action);
};

import { UIMatch } from "@remix-run/react";
import type { User } from "~/models/user";
import { loader } from "~/root";
import { useTypedMatchesData } from "./use-typed-match";

export function useOptionalUser(matches?: UIMatch[]): User | undefined {
  const routeMatch = useTypedMatchesData<typeof loader>({
    id: "root",
    matches,
  });

  return routeMatch?.user ?? undefined;
}

export function useUser(matches?: UIMatch[]): User {
  const maybeUser = useOptionalUser(matches);
  if (!maybeUser) {
    throw new Error(
      "No user found in root loader, but user is required by useUser. If user is optional, try useOptionalUser instead."
    );
  }
  return maybeUser;
}

export function useAuthToken(matches?: UIMatch[]): string | undefined | null {
  const routeMatch = useTypedMatchesData<typeof loader>({
    id: "root",
    matches,
  });

  return routeMatch?.accessToken;
}

export function useBackendUrl(matches?: UIMatch[]): string | undefined {
  const routeMatch = useTypedMatchesData<typeof loader>({
    id: "root",
    matches,
  });

  return routeMatch?.backendUrl;
}

export function usePlatformUrl(matches?: UIMatch[]): string | undefined {
  const routeMatch = useTypedMatchesData<typeof loader>({
    id: "root",
    matches,
  });

  return routeMatch?.platformUrl;
}

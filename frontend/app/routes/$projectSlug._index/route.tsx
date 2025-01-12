import React, { useCallback, useEffect, useRef, useState } from "react";
import AskB0 from "~/components/builder/ask-ai";
import DeployAndTestBtn from "~/components/builder/deploy";
import BuilderTools from "~/components/builder/tools";
import ZoomInAndOut from "~/components/builder/zoom";
import UserMenu from "~/components/menus/user-menu";
import invariant from "tiny-invariant";
import { typedjson, useTypedLoaderData } from "remix-typedjson";
import { z } from "zod";
import { ActionFunction, LoaderFunctionArgs, redirect } from "@remix-run/node";
import { requireUser } from "~/services/user.server";
import { Project } from "~/models/project";
import { getProject } from "~/services/project.server";
import { RouteErrorDisplay } from "~/components/error-boundary";
import { Position } from "~/models/flow";
import { PlaygroundProvider } from "~/components/builder/playground/provider";
import Playground from "~/components/builder/playground/playground";

export const ProjectSlugParamSchema = z.object({
  projectSlug: z.string(),
});

export const loader = async ({ request, params }: LoaderFunctionArgs) => {
  const { projectSlug } = ProjectSlugParamSchema.parse(params);

  invariant(projectSlug, "No project found in request.");

  const user = await requireUser(request);

  let project: Project | null = null;

  try {
    project = await getProject(request, projectSlug);
  } catch (err) {
    // redirect or throw not found
    throw new Response(null, {
      status: 404,
      statusText: "Not Found",
    });
  }

  // if (!appId) {
  //   return redirect(appsPath());
  // }

  return typedjson({
    user,
    projectSlug,
    project,
  });
};

export default function Page() {
  return (
    <PlaygroundProvider>
      <main className="h-full w-full relative canvas-bg">
        <Playground />

        <div className="absolute bottom-4 w-full">
          <div className="flex justify-between">
            <div className="flex items-center gap-2 ml-4">
              <UserMenu />
              <ZoomInAndOut />
            </div>
            <div>
              <AskB0 />
            </div>
            <div className="mr-4">
              <DeployAndTestBtn />
            </div>
          </div>
        </div>
        <BuilderTools />
      </main>
    </PlaygroundProvider>
  );
}

export function ErrorBoundary() {
  return <RouteErrorDisplay className="canvas-bg" />;
}

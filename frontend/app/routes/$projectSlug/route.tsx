import invariant from "tiny-invariant";
import { typedjson, useTypedLoaderData } from "remix-typedjson";
import { z } from "zod";
import { ActionFunction, LoaderFunctionArgs, redirect } from "@remix-run/node";
import { requireUser } from "~/services/user.server";
import { Project, Projects } from "~/models/project";
import { getProject, getProjects } from "~/services/project.server";
import { RouteErrorDisplay } from "~/components/error-boundary";
import { Outlet } from "@remix-run/react";
import { Endpoints } from "~/models/endpoint";
import { getEndpoints } from "~/services/endpoint.server";

export const ProjectSlugParamSchema = z.object({
  projectSlug: z.string(),
});

export const ProjectSearchSchema = z.object({
  endpoint: z.string().optional(),
});

export const loader = async ({ request, params }: LoaderFunctionArgs) => {
  const { projectSlug } = ProjectSlugParamSchema.parse(params);

  invariant(projectSlug, "No project found in request.");

  const user = await requireUser(request);

  let project: Project | null = null;
  let endpoints: Endpoints = [];
  let projects: Projects = [];

  const url = new URL(request.url);
  const s = Object.fromEntries(url.searchParams.entries());
  const searchParams = ProjectSearchSchema.parse(s);

  try {
    projects = await getProjects(request);

    project = await getProject(request, projectSlug);

    endpoints = await getEndpoints(request, project.id);
  } catch (err) {
    // redirect or throw not found
    throw new Response(null, {
      status: 404,
      statusText: "Not Found",
    });
  }

  return typedjson({
    user,
    projectSlug,
    project,
    projects,
    endpoints,
    endpoint: endpoints.find((e) => e.id === searchParams.endpoint),
  });
};

export default function Page() {
  return <Outlet />;
}

export function ErrorBoundary() {
  return <RouteErrorDisplay className="canvas-bg" />;
}

import { ServerResponseSchema } from "~/models/default";
import { api } from "./api.server";
import {
  CreateProjectForm,
  CreateProjectResponseSchema,
  GetProjects,
  GetProjectsSchema,
  GetProjectSchema,
  Project,
  GetProject,
} from "~/models/project";

export async function getProject(request: Request, id: string) {
  const res = await api.get<GetProject>({
    request,
    path: "/projects/" + id,
    schema: GetProjectSchema,
  });

  return res.data;
}

export async function getProjects(request: Request) {
  const res = await api.get<GetProjects>({
    request,
    path: "/projects",
    schema: GetProjectsSchema,
  });

  return res.data;
}

export async function createProject(request: Request, data: CreateProjectForm) {
  return await api.post({
    request,
    path: "/projects",
    body: data,
    schema: CreateProjectResponseSchema,
  });
}

export async function updateProject(
  request: Request,
  id: string,
  data: CreateProjectForm
) {
  return await api.put({
    request,
    path: `/projects/${id}`,
    body: data,
    schema: ServerResponseSchema,
  });
}

export async function deleteProject(
  request: Request,
  id: string,
  data: {
    name: string;
  }
) {
  return await api.delete({
    request,
    path: `/projects/${id}`,
    body: data,
    schema: ServerResponseSchema,
  });
}

import { z } from "zod";
import { ServerResponse, ServerResponseSchema } from "./default";

export const ProjectSchema = z.object({
  id: z.string(),
  name: z.string(),
  owner_id: z.string(),
  description: z.string().optional(),
  slug: z.string(),
  created_at: z.string(),
  updated_at: z.string(),
});

export type Project = z.infer<typeof ProjectSchema>;

export const GetProjectSchema = ServerResponseSchema.extend({
  data: ProjectSchema,
});

export type GetProject = z.infer<typeof GetProjectSchema>;

export const ProjectsSchema = z.array(ProjectSchema);

export type Projects = z.infer<typeof ProjectsSchema>;

export const GetProjectsSchema = ServerResponseSchema.extend({
  data: ProjectsSchema,
});

export type GetProjects = z.infer<typeof GetProjectsSchema>;

export const CreateProjectFormSchema = z.object({
  prompt: z.string(),
  model: z
    .enum(["gpt-4o", "claude-sonnet-3.5"])
    .default("claude-sonnet-3.5")
    .optional(),
});

export type CreateProjectForm = z.infer<typeof CreateProjectFormSchema>;

export const CreateProjectResponseSchema = ServerResponseSchema.extend({
  data: ProjectSchema,
});

export type CreateProjectResponse = z.infer<typeof CreateProjectResponseSchema>;

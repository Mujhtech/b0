import { z } from "zod";
import { ServerResponseSchema } from "./default";

export const EndpointSchema = z.object({
  id: z.string(),
  name: z.string(),
  description: z.string().optional(),
  owner_id: z.string(),
  project_id: z.string(),
  path: z.string(),
  method: z.enum(["GET", "POST", "PUT", "DELETE", "PATCH"]),
  is_public: z.boolean(),
  status: z.enum(["active", "inactive", "draft"]),
  created_at: z.string(),
  updated_at: z.string(),
});

export type Endpoint = z.infer<typeof EndpointSchema>;

export const GetEndpointSchema = ServerResponseSchema.extend({
  data: EndpointSchema,
});

export type GetEndpoint = z.infer<typeof GetEndpointSchema>;

export const EndpointsSchema = z.array(EndpointSchema);

export type Endpoints = z.infer<typeof EndpointsSchema>;

export const GetEndpointsSchema = ServerResponseSchema.extend({
  data: EndpointsSchema,
});

export type GetEndpoints = z.infer<typeof GetEndpointsSchema>;

export const CreateOrUpdateEndpointFormSchema = z.object({
  name: z.string(),
  description: z.string().optional(),
  method: z.enum(["GET", "POST", "PUT", "DELETE", "PATCH"]),
  path: z.string(),
  is_public: z.boolean().default(false).optional(),
});

export type CreateOrUpdateEndpointForm = z.infer<
  typeof CreateOrUpdateEndpointFormSchema
>;

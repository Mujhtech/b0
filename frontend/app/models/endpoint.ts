import { z } from "zod";
import { ServerResponseSchema } from "./default";

export const EndpointResponseStatusSchema = z.enum([
  "200",
  "201",
  "400",
  "401",
  "403",
  "404",
  "500",
]);

// extract the status code from the schema
export const EndpointResponseStatuses = EndpointResponseStatusSchema.Enum;

export const EndpointWorkflowCaseSchema = z.object({
  body: z.unknown().optional(),
  value: z.string().optional(),
});

export type EndpointWorkflowCase = z.infer<typeof EndpointWorkflowCaseSchema>;

export const EndpointWorkflowSchema = z.object({
  name: z.string().optional(),
  action_id: z.string().optional(),
  instruction: z.string().optional(),
  condition: z.string().optional(),
  method: z.string().optional(),
  url: z.string().optional(),
  type: z.enum([
    "request",
    "response",
    "if",
    "variable",
    "switch",
    "codeblock",
    "resend",
    "stripe",
    "openai",
    "github",
    "telegram",
    "discord",
    "slack",
    "supabase",
  ]),
  value: z.unknown().optional(),
  cases: z.array(EndpointWorkflowCaseSchema),
  then: z.unknown().optional(),
  else: z.unknown().optional(),
  body: z.unknown().optional(),
  model: z.string().optional(),
  provider: z.string().optional(),
  status: EndpointResponseStatusSchema.optional(),
});

export type EndpointWorkflow = z.infer<typeof EndpointWorkflowSchema>;

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
  workflows: z.array(EndpointWorkflowSchema),
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

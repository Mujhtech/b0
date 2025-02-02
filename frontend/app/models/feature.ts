import { z } from "zod";
import { ServerResponseSchema } from "./default";

export const ModelSchema = z.object({
  name: z.string(),
  model: z.string(),
  is_enabled: z.boolean().default(true).optional(),
  is_experimental: z.boolean().default(false).optional(),
  is_default: z.boolean().default(false).optional(),
});

export const FeatureSchema = z.object({
  name: z.string(),
  description: z.string(),
  is_github_auth_enabled: z.boolean(),
  is_google_auth_enabled: z.boolean(),
  is_aws_configured: z.boolean(),
  version: z.string(),
  available_models: z.array(ModelSchema),
});

export type Feature = z.infer<typeof FeatureSchema>;

export const GetFeatureSchema = ServerResponseSchema.extend({
  data: FeatureSchema,
});

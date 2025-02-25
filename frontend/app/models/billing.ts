import { z } from "zod";
import { ServerResponseSchema } from "./default";

export const UsageSchema = z.object({
  total_input_tokens: z.number().optional(),
  total_output_tokens: z.number().optional(),
  total_usage: z.number(),
});

export type Usage = z.infer<typeof UsageSchema>;

export const GetUsageSchema = ServerResponseSchema.extend({
  data: UsageSchema,
});

export type GetUsage = z.infer<typeof GetUsageSchema>;

export const UpgradePlanSchema = ServerResponseSchema.extend({
  data: z.object({
    portal_link: z.string(),
  }),
});

export type UpgradePlan = z.infer<typeof UpgradePlanSchema>;

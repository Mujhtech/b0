import { z } from "zod";
import { ServerResponse, ServerResponseSchema } from "./default";

export const SecretSchema = z.object({
  name: z.string(),
  value: z.string(),
  note: z.string().optional(),
  protected: z.boolean().optional(),
});

export type Secret = z.infer<typeof SecretSchema>;

export const SecretsSchema = z.array(SecretSchema);

export type Secrets = z.infer<typeof SecretsSchema>;

export const GetSecretsSchema = ServerResponseSchema.extend({
  data: SecretsSchema,
});

export type GetSecrets = z.infer<typeof GetSecretsSchema>;

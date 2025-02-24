import { GetSecrets, GetSecretsSchema, Secrets } from "~/models/secret";
import { api } from "./api.server";
import { ServerResponseSchema } from "~/models/default";

export async function getSecrets(request: Request, id: string) {
  const res = await api.get<GetSecrets>({
    request,
    path: "/projects/" + id + "/secrets",
    schema: GetSecretsSchema,
  });

  return res.data;
}

export async function createOrUpdateSceret(
  request: Request,
  id: string,
  data: Secrets
) {
  return await api.put({
    request,
    path: `/projects/${id}/secrets`,
    body: data,
    schema: ServerResponseSchema,
  });
}

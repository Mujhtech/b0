import { GetEndpoints, GetEndpointsSchema } from "~/models/endpoint";
import { api } from "./api.server";

export async function getEndpoints(request: Request, id?: string) {
  const res = await api.get<GetEndpoints>({
    request,
    path: "/endpoints?project_id=" + id,
    schema: GetEndpointsSchema,
  });

  return res.data;
}

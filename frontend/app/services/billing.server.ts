import { api } from "./api.server";
import { GetUsage, GetUsageSchema } from "~/models/billing";

export async function getUsage(request: Request) {
  const res = await api.get<GetUsage>({
    request,
    path: "/billing/usage",
    schema: GetUsageSchema,
  });

  return res.data;
}

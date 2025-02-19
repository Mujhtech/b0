import { api } from "./api.server";
import {
  GetUsage,
  GetUsageSchema,
  UpgradePlan,
  UpgradePlanSchema,
} from "~/models/billing";

export async function getUsage(request: Request) {
  const res = await api.get<GetUsage>({
    request,
    path: "/billing/usage",
    schema: GetUsageSchema,
  });

  return res.data;
}

export async function upgradePlan(request: Request, plan: string) {
  const res = await api.post<UpgradePlan>({
    request,
    path: "/billing/upgrade",
    schema: UpgradePlanSchema,
    body: {
      plan,
    },
  });

  return res.data.portal_link;
}

import { Feature, FeatureSchema, GetFeatureSchema } from "~/models/feature";
import { api } from "./api.server";

export async function getFeatures(request: Request) {
  const res = await api.get({
    request,
    path: "/features",
    schema: GetFeatureSchema,
  });

  return res.data;
}

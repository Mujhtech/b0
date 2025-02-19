import { parseWithZod } from "@conform-to/zod";
import { ActionFunction, MetaFunction, redirect } from "@remix-run/node";
import React from "react";
import { z } from "zod";
import {
  BasicPlanCard,
  ScalePlanCard,
  ProPlanCard,
} from "~/components/billing/plan";
import { Card, CardContent, CardHeader } from "~/components/ui/card";
import { redirectBackWithErrorMessage } from "~/models/message.server";
import { upgradePlan } from "~/services/billing.server";

export const action: ActionFunction = async ({ request, params }) => {
  const formData = await request.formData();

  const submission = parseWithZod(formData, {
    schema: z.object({
      plan: z.string(),
    }),
  });

  if (submission.status !== "success") {
    return redirectBackWithErrorMessage(request, "Invalid form submission");
  }

  try {
    // handle creation of project
    const portal_link = await upgradePlan(request, submission.value.plan);

    let headers = new Headers({});

    return redirect(portal_link, {
      headers,
    });
  } catch (e: any) {
    return redirectBackWithErrorMessage(request, e.error ?? "Unknown error");
  }
};

export const meta: MetaFunction = () => {
  return [
    {
      title: `Billing`,
    },
  ];
};

export default function Page() {
  return (
    <div className="flex flex-col w-full max-w-5xl mx-auto">
      <Card className="mb-8">
        <CardHeader className="flex flex-col p-4">
          <h1 className="font-semibold">Billing</h1>
        </CardHeader>
        <CardContent className="grid grid-cols-1 md:grid-cols-3 gap-4 px-4 pb-4">
          <BasicPlanCard />
          <ProPlanCard />
          <ScalePlanCard />
        </CardContent>
      </Card>
    </div>
  );
}

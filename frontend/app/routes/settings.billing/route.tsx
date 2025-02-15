import { MetaFunction } from "@remix-run/node";
import React from "react";
import {
  BasicPlanCard,
  ScalePlanCard,
  ProPlanCard,
} from "~/components/billing/plan";
import { Card, CardContent, CardHeader } from "~/components/ui/card";

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

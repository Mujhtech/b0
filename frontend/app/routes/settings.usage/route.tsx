import { LoaderFunctionArgs, MetaFunction } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { useMemo } from "react";
import { typedjson, useTypedLoaderData } from "remix-typedjson";
import { Card, CardContent, CardHeader } from "~/components/ui/card";
import Paragraph from "~/components/ui/paragraph";
import { Progress } from "~/components/ui/progress";
import { useUser } from "~/hooks/use-user";
import { Usage } from "~/models/billing";
import { getUsage } from "~/services/billing.server";

export const loader = async ({ request, params }: LoaderFunctionArgs) => {
  let usage: Usage | null = null;

  try {
    usage = await getUsage(request);
  } catch (err) {}

  return typedjson({
    usage,
  });
};

export const meta: MetaFunction = () => {
  return [
    {
      title: `Usage`,
    },
  ];
};

export default function Page() {
  const { usage } = useTypedLoaderData<typeof loader>();

  const user = useUser();

  const usageLimit = useMemo(() => {
    switch (user.subscription_plan) {
      case "free":
        return 20;
      case "starter":
        return 50;
      case "pro":
        return 100;
      case "scale":
        return 500;
      default:
        return 0;
    }
  }, [user]);

  const percentage = useMemo(() => {
    return (usage!.total_usage / usageLimit) * 100;
  }, [usageLimit, usage]);

  const usageLeft = useMemo(() => {
    if (usageLimit === 0) {
      return "Unlimited";
    }

    return usageLimit - usage!.total_usage;
  }, []);

  return (
    <div className="flex flex-col w-full max-w-5xl mx-auto">
      <Card className="mb-8">
        <CardHeader className="flex flex-col p-4">
          <h1 className="font-semibold">Usage</h1>
        </CardHeader>
        <CardContent className="px-4 pb-4">
          <div className="flex flex-col gap-2">
            <div className="flex justify-between">
              <Paragraph className="capitalize">
                {user.subscription_plan}
              </Paragraph>
              <Paragraph>
                {usage!.total_usage}/{usageLimit > 0 ? usageLimit : "Unlimited"}
              </Paragraph>
            </div>
            <Progress value={percentage} />
            <Paragraph>{usageLeft} remaining limit</Paragraph>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

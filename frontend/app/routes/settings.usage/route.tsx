import { MetaFunction } from "@remix-run/node";
import { useMemo } from "react";
import { Card, CardContent, CardHeader } from "~/components/ui/card";
import Paragraph from "~/components/ui/paragraph";
import { Progress } from "~/components/ui/progress";

export const meta: MetaFunction = () => {
  return [
    {
      title: `Usage`,
    },
  ];
};

export default function Page() {
  const percentage = useMemo(() => {
    return (150 / 200) * 100;
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
              <Paragraph>Free</Paragraph>
              <Paragraph>1/200</Paragraph>
            </div>
            <Progress value={percentage} />
            <Paragraph>199 remaining limit</Paragraph>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

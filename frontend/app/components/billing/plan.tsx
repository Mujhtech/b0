import React from "react";
import { Card, CardContent, CardHeader } from "../ui/card";
import { Form } from "@remix-run/react";
import { Button } from "../ui/button";
import { CheckIcon, X } from "lucide-react";
import { cn } from "~/lib/utils";
import { Separator } from "../ui/separator";
import Paragraph from "../ui/paragraph";
import { useUser } from "~/hooks/use-user";

export function BasicPlanCard() {
  return (
    <PlanCard
      title="free"
      price={0}
      storage={1}
      feature={
        <ul className="flex flex-col gap-3">
          <Item checked>Access to b0</Item>
          <Item checked>Up 3 public projects</Item>
          <Item checked>Deploy on the go</Item>
          <Item checked>Export project</Item>
          <Item checked>Integration</Item>
          <Item checked={false}>
            Access to Claude Sonnet 3.5 & OpenAI GPT-4o
          </Item>
          <Item checked={false}>Pay-as-you-go for additional usage</Item>
          <Item checked={false}>Custom Domain</Item>
          <Item checked={false}>Log retention</Item>
          <Item checked>Community Support</Item>
        </ul>
      }
    />
  );
}

export function ProPlanCard() {
  return (
    <PlanCard
      recommended
      title="pro"
      price={20}
      storage={5}
      feature={
        <ul className="flex flex-col gap-3">
          <Item checked>Access to b0</Item>
          <Item checked>2.5x monthly limits</Item>
          <Item checked>Unlimited public and private project</Item>
          <Item checked>Deploy on the go</Item>
          <Item checked>Export project</Item>
          <Item checked>Access to Claude Sonnet 3.5 & OpenAI GPT-4o</Item>
          <Item checked>Integration</Item>
          <Item checked>Pay-as-you-go for additional usage</Item>
          <Item checked>Custom Domain</Item>
          <Item checked>Log retention</Item>
          <Item checked>Support</Item>
        </ul>
      }
    />
  );
}

export function ScalePlanCard() {
  return (
    <PlanCard
      title="scale"
      price={100}
      feature={
        <ul className="flex flex-col gap-3">
          <Item checked>Access to b0</Item>
          <Item checked>Higher monthly limits</Item>
          <Item checked>Unlimited public and private project</Item>
          <Item checked>Deploy on the go</Item>
          <Item checked>Export project</Item>
          <Item checked>Integration</Item>
          <Item checked>Access to Claude Sonnet 3.5 & OpenAI GPT-4o</Item>
          <Item checked>Pay-as-you-go for additional usage</Item>
          <Item checked>Custom Domain</Item>
          <Item checked>Log retention</Item>
          <Item checked>Dedicated Support</Item>
        </ul>
      }
    />
  );
}

export default function PlanCard({
  recommended,
  title,
  feature,
  price,
  storage,
}: {
  recommended?: boolean;
  title: string;
  feature: React.ReactNode;
  price?: number;
  storage?: number;
}) {
  const user = useUser();
  return (
    <Card className={cn("w-full p-2", recommended && "border-white border-2")}>
      <Form method="post">
        <CardHeader className="!gap-3">
          <h3 className="font-medium text-muted-foreground capitalize">
            {title}
          </h3>
          <h1 className="text-4xl font-semibold">
            {price != undefined ? `$${price}` : "Custom"}
            <span className="text-sm font-medium text-muted-foreground">
              /month
            </span>
          </h1>
        </CardHeader>
        <CardContent className="!pb-16">
          <Paragraph className="mb-4 !text-sm">
            {storage ? `Up to ${storage}GB storage` : "Custom storage"}
          </Paragraph>
          <Separator />
          <Button
            className="my-5 !h-8 w-full"
            disabled={user.subscription_plan === title}
          >
            Upgrade
          </Button>
          {feature}
        </CardContent>
      </Form>
    </Card>
  );
}

const Item = ({
  checked,
  children,
}: {
  checked: boolean;
  children: React.ReactNode;
}) => {
  return (
    <li className="flex items-center gap-2">
      {checked ? (
        <CheckIcon className="h-4 w-4 text-primary" />
      ) : (
        <X className="h-4 w-4 text-muted-foreground" />
      )}
      <div className={cn("text-sm", checked ? "" : "text-muted-foreground")}>
        {children}
      </div>
    </li>
  );
};

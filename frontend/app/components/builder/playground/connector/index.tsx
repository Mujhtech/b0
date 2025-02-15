import React from "react";
import HttpResponseConnector from "./http-response";
import IfConnector from "./if";
import VariableConnector from "./variable";
import SwitchConnector from "./switch";
import CodeblockConnector from "./codeblock";
import OpenAIConnector from "./openai";
import TelegramConnector from "./telegram";
import SlackConnector from "./slack";
import DiscordConnector from "./discord";
import GithubConnector from "./github";
import SupabaseConnector from "./supabase";
import ResendConnector from "./resend";
import StripeConnector from "./stripe";
import { EndpointWorkflow } from "~/models/endpoint";

export default function Connector({
  workflow,
  index,
}: {
  workflow: EndpointWorkflow;
  index: number;
}) {
  switch (workflow.type) {
    case "response":
      return <HttpResponseConnector key={index} workflow={workflow} />;
    case "if":
      return <IfConnector key={index} workflow={workflow} />;
    case "variable":
      return <VariableConnector key={index} workflow={workflow} />;
    case "switch":
      return <SwitchConnector key={index} workflow={workflow} />;
    case "codeblock":
      return <CodeblockConnector key={index} workflow={workflow} />;
    case "openai":
      return <OpenAIConnector key={index} workflow={workflow} />;
    case "telegram":
      return <TelegramConnector key={index} workflow={workflow} />;
    case "slack":
      return <SlackConnector key={index} workflow={workflow} />;
    case "discord":
      return <DiscordConnector key={index} workflow={workflow} />;
    case "github":
      return <GithubConnector key={index} workflow={workflow} />;
    case "supabase":
      return <SupabaseConnector key={index} workflow={workflow} />;
    case "resend":
      return <ResendConnector key={index} workflow={workflow} />;
    case "stripe":
      return <StripeConnector key={index} workflow={workflow} />;
    case "request":
    default:
      return null;
  }
}

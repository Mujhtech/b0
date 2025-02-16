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

export interface ConnectorProps {
  workflow: EndpointWorkflow;
  index: number;
  onRemove?: () => void;
  onUpdate?: (workflow: EndpointWorkflow) => void;
}

export default function Connector(props: ConnectorProps) {
  switch (props.workflow.type) {
    case "response":
      return <HttpResponseConnector key={props.index} {...props} />;
    case "if":
      return <IfConnector key={props.index} {...props} />;
    case "variable":
      return <VariableConnector key={props.index} {...props} />;
    case "switch":
      return <SwitchConnector key={props.index} {...props} />;
    case "codeblock":
      return <CodeblockConnector key={props.index} {...props} />;
    case "openai":
      return <OpenAIConnector key={props.index} {...props} />;
    case "telegram":
      return <TelegramConnector key={props.index} {...props} />;
    case "slack":
      return <SlackConnector key={props.index} {...props} />;
    case "discord":
      return <DiscordConnector key={props.index} {...props} />;
    case "github":
      return <GithubConnector key={props.index} {...props} />;
    case "supabase":
      return <SupabaseConnector key={props.index} {...props} />;
    case "resend":
      return <ResendConnector key={props.index} {...props} />;
    case "stripe":
      return <StripeConnector key={props.index} {...props} />;
    case "request":
    default:
      return null;
  }
}

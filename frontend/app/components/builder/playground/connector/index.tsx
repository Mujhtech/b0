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
import { useSortable } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";

export interface ConnectorProps {
  workflow: EndpointWorkflow;
  index: number;
  onRemove?: () => void;
  onUpdate?: (workflow: EndpointWorkflow) => void;
  draggable?: {
    isActive?: boolean;
    isDragging?: boolean;
    style?: {
      transform: string | undefined;
      transition: string | undefined;
    };
    listeners?: any;
    attributes?: any;
    setNodeRef?: any;
  };
}

export default function Connector(props: ConnectorProps) {
  const { attributes, listeners, setNodeRef, transform, transition } =
    useSortable({ id: props.workflow.action_id ?? props.index.toString() });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  const draggable = {
    isActive: props.draggable?.isActive,
    isDragging: props.draggable?.isDragging,
    style,
    setNodeRef,
    attributes,
    listeners,
  };
  switch (props.workflow.type) {
    case "response":
      return <HttpResponseConnector key={props.index} {...props} />;
    case "if":
      return <IfConnector key={props.index} {...props} draggable={draggable} />;
    case "variable":
      return (
        <VariableConnector key={props.index} {...props} draggable={draggable} />
      );
    case "switch":
      return (
        <SwitchConnector key={props.index} {...props} draggable={draggable} />
      );
    case "codeblock":
      return (
        <CodeblockConnector
          key={props.index}
          {...props}
          draggable={draggable}
        />
      );
    case "openai":
      return (
        <OpenAIConnector key={props.index} {...props} draggable={draggable} />
      );
    case "telegram":
      return (
        <TelegramConnector key={props.index} {...props} draggable={draggable} />
      );
    case "slack":
      return (
        <SlackConnector key={props.index} {...props} draggable={draggable} />
      );
    case "discord":
      return (
        <DiscordConnector key={props.index} {...props} draggable={draggable} />
      );
    case "github":
      return (
        <GithubConnector key={props.index} {...props} draggable={draggable} />
      );
    case "supabase":
      return (
        <SupabaseConnector key={props.index} {...props} draggable={draggable} />
      );
    case "resend":
      return (
        <ResendConnector key={props.index} {...props} draggable={draggable} />
      );
    case "stripe":
      return (
        <StripeConnector key={props.index} {...props} draggable={draggable} />
      );
    case "request":
    default:
      return null;
  }
}

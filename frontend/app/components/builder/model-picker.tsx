import { Sparkle } from "@phosphor-icons/react";
import * as React from "react";

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "~/components/ui/select";
import { useFeature } from "~/hooks/use-feature";
import { useUser } from "~/hooks/use-user";

export function AIModelPicker({
  onSelect,
  value,
}: {
  onSelect: (model: string) => void;
  value?: string;
}) {
  const { available_models } = useFeature();
  const user = useUser();

  return (
    <Select onValueChange={onSelect} value={value}>
      <SelectTrigger className="focus:outline-none focus-visible:ring-0 h-8 w-auto  border-r border-t-0 border-l-0 border-b-0 focus:ring-0">
        <Sparkle size={20} className="h-4 w-4 mr-1" />
        <SelectValue placeholder="Select a model" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {available_models.map((model, index) => (
            <SelectItem
              key={index}
              value={model.model}
              disabled={
                model.is_enabled === false ||
                (user.subscription_plan === "free" && model.is_premium === true)
              }
            >
              {model.name}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}

export function AIModelPicker2({
  onSelect,
  value,
}: {
  onSelect: (model: string) => void;
  value?: string;
}) {
  const { available_models } = useFeature();
  const user = useUser();

  return (
    <Select onValueChange={onSelect} value={value}>
      <SelectTrigger className="focus:outline-none focus-visible:ring-0 border border-input h-6 w-auto p-1 inline-flex items-center justify-center focus:ring-0">
        <Sparkle size={20} className="h-4 w-4 mr-1" />
        <SelectValue placeholder="Select a model" />
      </SelectTrigger>
      <SelectContent side="top">
        <SelectGroup>
          {available_models.map((model, index) => (
            <SelectItem
              key={index}
              value={model.model}
              disabled={
                model.is_enabled === false ||
                (user.subscription_plan === "free" && model.is_premium === true)
              }
            >
              {model.name}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}

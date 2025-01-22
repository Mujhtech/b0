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

const models = [
  {
    name: "Claude Sonnet 3.5",
    value: "claude-sonnet-3.5",
    isExperimental: true,
    disabled: true,
  },
  {
    name: "GPT 4o",
    value: "gpt-4o",
    isExperimental: true,
    disabled: true,
  },
  {
    name: "GPT 3.5",
    value: "gpt-3.5",
    isExperimental: true,
    disabled: true,
  },
  {
    name: "Gemini 1.5 flash",
    value: "gemini-1.5-flash",
    isExperimental: true,
    disabled: false,
  },
];

export function AIModelPicker() {
  return (
    <Select>
      <SelectTrigger className="focus:outline-none focus-visible:ring-0 h-8 w-auto  border-r border-t-0 border-l-0 border-b-0 focus:ring-0">
        <Sparkle size={20} className="h-4 w-4 mr-1" />
        <SelectValue placeholder="Select a model" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {models.map((model, index) => (
            <SelectItem key={index} value={model.value}>
              {model.name}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}

export function AIModelPicker2() {
  return (
    <Select>
      <SelectTrigger className="focus:outline-none focus-visible:ring-0 border border-input h-6 w-auto p-1 inline-flex items-center justify-center focus:ring-0">
        <Sparkle size={20} className="h-4 w-4 mr-1" />
        <SelectValue placeholder="Select a model" />
      </SelectTrigger>
      <SelectContent side="top">
        <SelectGroup>
          {models.map((model, index) => (
            <SelectItem key={index} value={model.value}>
              {model.name}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}

import { Sparkle } from "@phosphor-icons/react";
import * as React from "react";

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
} from "~/components/ui/select";

export function AIModelPicker() {
  return (
    <Select>
      <SelectTrigger className="focus:outline-none focus-visible:ring-0 h-8 w-auto  border-r border-t-0 border-l-0 border-b-0">
        <Sparkle size={20} className="h-4 w-4 mr-1" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectItem value="claude-sonnet-3.5">Claude Sonnet 3.5</SelectItem>
          <SelectItem value="gpt-4o">GPT 4o</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}

export function AIModelPicker2() {
  return (
    <Select>
      <SelectTrigger className="focus:outline-none focus-visible:ring-0 border border-input h-6 w-auto p-1 inline-flex items-center justify-center">
        <Sparkle size={20} className="h-4 w-4 mr-1" />
      </SelectTrigger>
      <SelectContent side="top">
        <SelectGroup>
          <SelectItem value="claude-sonnet-3.5">Claude Sonnet 3.5</SelectItem>
          <SelectItem value="gpt-4o">GPT 4o</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}

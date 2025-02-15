import { Sparkle } from "@phosphor-icons/react";

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "~/components/ui/select";
import { useFeature } from "~/hooks/use-feature";

export function LanguagePicker({
  onSelect,
  value,
}: {
  onSelect: (id: string) => void;
  value?: string;
}) {
  const { available_languages } = useFeature();

  return (
    <Select onValueChange={onSelect} value={value}>
      <SelectTrigger className="focus:outline-none focus-visible:ring-0 border border-input h-6 w-auto p-1 inline-flex items-center justify-center focus:ring-0">
        <Sparkle size={20} className="h-4 w-4 mr-1" />
        <SelectValue placeholder="Select a language" />
      </SelectTrigger>
      <SelectContent side="top" align="end">
        <SelectGroup>
          {available_languages.map((language, index) => (
            <SelectItem key={index} value={language.id}>
              {language.language} ({language.framework})
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}

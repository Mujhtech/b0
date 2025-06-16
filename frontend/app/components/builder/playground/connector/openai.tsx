import React, { useCallback, useState } from "react";
import { EndpointWorkflow } from "~/models/endpoint";
import { ConnectorProps } from ".";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "~/components/ui/sheet";
import { Label } from "~/components/ui/label";
import { Textarea } from "~/components/ui/textarea";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "~/components/ui/form";
import { z } from "zod";
import { Button } from "~/components/ui/button";
import { CaretDown, Check, X } from "@phosphor-icons/react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs";
import { Input } from "~/components/ui/input";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "~/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "~/components/ui/popover";
import { cn } from "~/lib/utils";

const providers = ["openai", "deepseek", "google"];
const models = [
  {
    id: "gpt-3",
    name: "GPT-3",
    provider: "openai",
  },
  {
    id: "gpt-4",
    name: "GPT-4",
    provider: "openai",
  },
  {
    id: "gpt-5",
    name: "GPT-5",
    provider: "openai",
  },
  {
    id: "gemini-1.5-flash",
    name: "Gemini 1.5 Flash",
    provider: "google",
  },
  {
    id: "gemini-2.0-flash",
    name: "Gemini 2.0 Flash",
    provider: "google",
  },
  {
    id: "deepseek-chat",
    name: "Deepseek Chat",
    provider: "deepseek",
  },
  {
    id: "deepseek-r1",
    name: "Deepseek R1",
    provider: "deepseek",
  },
];

const formSchema = z.object({
  provider: z.string().optional(),
  model: z.string().optional(),
});

export default function OpenAIConnector({
  workflow,
  draggable,
  onRemove,
  onUpdate,
}: ConnectorProps) {
  const [isOpen, setIsOpen] = useState(false);

  const handleClose = useCallback(() => {
    setIsOpen(false);
  }, [isOpen]);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      provider: workflow?.provider,
      model: workflow?.model,
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      onUpdate?.({
        ...workflow,
        provider: values.provider,
        model: values.model,
      });
      setIsOpen(false);
    } catch (error) {
      console.error(error);
    }
  }

  const handleDelete = () => {
    onRemove?.();
    setIsOpen(false);
  };

  return (
    <Sheet open={isOpen} onOpenChange={setIsOpen}>
      <SheetTrigger>
        <div
          className="flex flex-col group h-min"
          ref={draggable?.setNodeRef}
          style={draggable?.style}
          {...draggable?.attributes}
          {...draggable?.listeners}
        >
          <div className="border border-input w-[250px] bg-background shadow-sm flex self-center flex-col hover:drop-shadow-xl cursor-pointer">
            <div className="border-b border-input flex items-center justify-center">
              <h3 className="text-xs font-mono text-muted-foreground p-2">
                OpenAI
              </h3>
            </div>
            <div className="p-2">
              {workflow && (
                <div>
                  <p className="text-xs font-mono text-muted-foreground">
                    {workflow.instruction}
                  </p>
                </div>
              )}
            </div>
          </div>
          <div className="relative flex self-center min-h-[22px] border-solid border-x border-input cursor-pointer"></div>
        </div>
      </SheetTrigger>
      <SheetContent
        className="bg-transparent border-none p-4"
        hideCloseBtn={true}
      >
        <SheetDescription className="hidden"></SheetDescription>
        <div className="relative bg-background w-full h-full border border-input">
          <SheetHeader className="flex flex-row items-center justify-between border-b border-input space-y-0 p-2">
            <SheetTitle className="text-xs font-mono text-muted-foreground">
              OpenAI
            </SheetTitle>
            <Button
              variant={"outline"}
              size={"icon"}
              onClick={handleClose}
              className="h-5 w-5 [&_svg]:size-4"
            >
              <X size={16} className="h-4 w-4" />
            </Button>
          </SheetHeader>
          <Tabs defaultValue="general" className="w-full">
            <TabsList className="w-full bg-transparent justify-start border-b border-input">
              <TabsTrigger value="general" className="font-mono text-xs">
                General
              </TabsTrigger>
            </TabsList>
            <TabsContent value="general">
              <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)}>
                  <div className="flex flex-col p-2">
                    <div className="flex flex-col gap-2">
                      <FormField
                        control={form.control}
                        name="provider"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Provider</FormLabel>
                            <FormControl>
                              <SelecteProviderDialog
                                value={field.value}
                                onChange={(value) => {
                                  field.onChange(value);
                                  form.resetField("model");
                                }}
                              />
                            </FormControl>

                            <FormMessage />
                          </FormItem>
                        )}
                      />
                      <FormField
                        control={form.control}
                        name="model"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Model</FormLabel>
                            <FormControl>
                              <SelectModelDialog
                                value={field.value}
                                onChange={field.onChange}
                                models={models
                                  .filter(
                                    (item) =>
                                      item.provider ===
                                      form.getValues("provider")
                                  )
                                  .map((item) => item.id)}
                              />
                            </FormControl>

                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>
                  </div>
                  <div className="bg-background border-t border-input absolute w-full bottom-0 p-2">
                    <div className="flex items-center justify-end gap-2">
                      <Button
                        type="button"
                        onClick={handleDelete}
                        variant="outline"
                        className="h-8 shadow-lg border-destructive text-red-500 hover:border-destructive hover:text-red-500"
                      >
                        Delete
                      </Button>
                      <Button className="h-8 shadow-lg">Save</Button>
                    </div>
                  </div>
                </form>
              </Form>
            </TabsContent>
          </Tabs>
        </div>
      </SheetContent>
    </Sheet>
  );
}

function SelectModelDialog({
  value,
  onChange,
  models,
}: {
  value?: string;
  onChange: (value: string) => void;
  models: string[];
}) {
  const [open, setOpen] = useState(false);

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between"
        >
          {value ? models.find((item) => item === value) : "Select model..."}
          <CaretDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0">
        <Command>
          <CommandInput placeholder="Search model..." />
          <CommandList>
            <CommandEmpty>No model found.</CommandEmpty>
            <CommandGroup>
              {models.map((model) => (
                <CommandItem
                  key={model}
                  value={model}
                  onSelect={(currentValue) => {
                    onChange(currentValue === value ? "" : currentValue);
                    setOpen(false);
                  }}
                >
                  <Check
                    className={cn(
                      "mr-2 h-4 w-4",
                      value === model ? "opacity-100" : "opacity-0"
                    )}
                  />
                  {model}
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}

function SelecteProviderDialog({
  value,
  onChange,
}: {
  value?: string;
  onChange: (value: string) => void;
}) {
  const [open, setOpen] = useState(false);

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between"
        >
          {value
            ? providers.find((item) => item === value)
            : "Select provider..."}
          <CaretDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0">
        <Command>
          <CommandInput placeholder="Search method..." />
          <CommandList>
            <CommandEmpty>No provider found.</CommandEmpty>
            <CommandGroup>
              {providers.map((provider) => (
                <CommandItem
                  key={provider}
                  value={provider}
                  onSelect={(currentValue) => {
                    onChange(currentValue === value ? "" : currentValue);
                    setOpen(false);
                  }}
                >
                  <Check
                    className={cn(
                      "mr-2 h-4 w-4 capitalize",
                      value === provider ? "opacity-100" : "opacity-0"
                    )}
                  />
                  {provider}
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}

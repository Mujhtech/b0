import { CaretDown, Check, X } from "@phosphor-icons/react";
import React, { useCallback, useState } from "react";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";

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
import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs";
import {
  EndpointResponseStatusSchema,
  EndpointWorkflow,
} from "~/models/endpoint";
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
import { CodeEditor } from "~/components/code-editor/view";
import { ConnectorProps } from ".";

const statuses = ["200", "201", "400", "401", "403", "404", "500"];

const formSchema = z.object({
  status: EndpointResponseStatusSchema,
  description: z.string().optional(),
  body: z.string().optional(),
});

export default function HttpResponseConnector({
  workflow,
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
      description: workflow?.instruction,
      status: workflow?.status,
      body: JSON.stringify(workflow?.body, null, 2),
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      onUpdate?.({
        ...workflow,
        status: values.status,
        instruction: values.description,
        body: values.body ? JSON.parse(values.body) : undefined,
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
        <div className="flex flex-col group h-min">
          <div className="border border-input w-[250px] bg-background shadow-sm flex self-center flex-col hover:drop-shadow-xl cursor-pointer">
            <div className="border-b border-input flex items-center justify-center">
              <h3 className="text-xs font-mono text-muted-foreground p-2">
                HTTP Response
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
          {/* <div className="relative flex self-center min-h-[22px] border-solid border-x border-input cursor-pointer"></div> */}
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
              HTTP Response
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
              {/* <TabsTrigger value="body" className="font-mono text-xs">
                Body
              </TabsTrigger>
              <TabsTrigger value="other" className="font-mono text-xs">
                Other
              </TabsTrigger> */}
            </TabsList>
            <TabsContent value="general">
              <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)}>
                  <div className="flex flex-col p-2">
                    <div className="flex flex-col gap-2">
                      <FormField
                        control={form.control}
                        name="description"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Description (Optional)</FormLabel>
                            <FormControl>
                              <Textarea
                                placeholder="Description (Optional)"
                                className="resize-none"
                                {...field}
                              />
                            </FormControl>

                            <FormMessage />
                          </FormItem>
                        )}
                      />

                      <FormField
                        control={form.control}
                        name="status"
                        render={({ field }) => (
                          <FormItem>
                            <Label>Status</Label>
                            <SelectStatusDialog
                              value={field.value}
                              onChange={(val) => {
                                field.onChange(val);
                              }}
                            />
                          </FormItem>
                        )}
                      />

                      <FormField
                        control={form.control}
                        name="body"
                        render={({ field }) => (
                          <FormItem>
                            <Label>Body</Label>
                            <CodeEditor
                              defaultValue={field.value}
                              readOnly={false}
                              basicSetup
                              showClearButton={false}
                              showCopyButton={false}
                              onChange={(v) => {
                                field.onChange(v);
                              }}
                              height="100%"
                              autoFocus={false}
                              placeholder=""
                              className={cn(
                                "h-full overflow-auto border border-input"
                              )}
                            />
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
            <TabsContent value="body">
              <div className="p-2">Coming soon</div>
            </TabsContent>
            <TabsContent value="other">
              <div className="p-2">Coming soon</div>
            </TabsContent>
          </Tabs>
        </div>
      </SheetContent>
    </Sheet>
  );
}

export function SelectStatusDialog({
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
          {value ? statuses.find((item) => item === value) : "Select status..."}
          <CaretDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0">
        <Command>
          <CommandInput placeholder="Search method..." />
          <CommandList>
            <CommandEmpty>No status found.</CommandEmpty>
            <CommandGroup>
              {statuses.map((method) => (
                <CommandItem
                  key={method}
                  value={method}
                  onSelect={(currentValue) => {
                    onChange(currentValue === value ? "" : currentValue);
                    setOpen(false);
                  }}
                >
                  <Check
                    className={cn(
                      "mr-2 h-4 w-4",
                      value === method ? "opacity-100" : "opacity-0"
                    )}
                  />
                  {method}
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}

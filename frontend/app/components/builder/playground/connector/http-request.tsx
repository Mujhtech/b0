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
import { clientApi } from "~/services/api.client";
import {
  CreateOrUpdateEndpointForm,
  CreateOrUpdateEndpointFormSchema,
  GetEndpoint,
  GetEndpointSchema,
} from "~/models/endpoint";
import { useNavigate } from "@remix-run/react";
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
import { useAuthToken, usePlatformUrl } from "~/hooks/use-user";
import { useOptionalEndpoint, useProject } from "~/hooks/use-project";
import { Switch } from "~/components/ui/switch";
import { Badge } from "~/components/ui/badge";

const methods = ["GET", "POST", "PUT", "DELETE", "PATCH"];

export default function HttpRequestConnector() {
  const endpoint = useOptionalEndpoint();
  const [isOpen, setIsOpen] = useState(false);
  const navigate = useNavigate();

  const handleClose = useCallback(() => {
    setIsOpen(false);
  }, [isOpen]);

  const accessToken = useAuthToken();
  const backendBaseUrl = usePlatformUrl();
  const project = useProject();

  const form = useForm<CreateOrUpdateEndpointForm>({
    resolver: zodResolver(CreateOrUpdateEndpointFormSchema),
    defaultValues: {
      name: endpoint?.name ?? "",
      description: endpoint?.description ?? "",
      method: endpoint?.method,
      path: endpoint?.path ?? "",
      is_public: endpoint?.is_public ?? false,
    },
  });

  async function onSubmit(values: CreateOrUpdateEndpointForm) {
    try {
      if (!endpoint) {
        const res = await clientApi.post<GetEndpoint>({
          path: "/endpoints",
          body: { ...values, project_id: project.id },
          backendBaseUrl: backendBaseUrl!,
          accessToken: accessToken!,
          schema: GetEndpointSchema,
        });
        navigate("?endpoint" + res.data.id);
      } else {
        await clientApi.put<GetEndpoint>({
          path: `/endpoints/${endpoint?.id}`,
          body: { ...values, project_id: project.id },
          backendBaseUrl: backendBaseUrl!,
          accessToken: accessToken!,
          schema: GetEndpointSchema,
        });
        navigate("?endpoint" + endpoint?.id);
      }
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <Sheet open={isOpen} onOpenChange={setIsOpen}>
      <SheetTrigger>
        <div className="flex flex-col group h-min">
          <div className="border border-input w-[250px] bg-background shadow-sm flex self-center flex-col hover:drop-shadow-xl cursor-pointer">
            <div className="border-b border-input flex items-center justify-center">
              {endpoint && (
                <Badge className="font-mono text-[10px] mr-1">
                  {endpoint.method}
                </Badge>
              )}
              <h3 className="text-xs font-mono text-muted-foreground p-2">
                HTTP Request
              </h3>
            </div>
            <div className="p-2">
              {endpoint && (
                <div>
                  <p className="text-xs font-mono text-muted-foreground">
                    {endpoint.name}
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
        <div className="relative bg-background w-full h-full border border-input">
          <SheetHeader className="flex flex-row items-center justify-between border-b border-input space-y-0 p-2">
            <SheetTitle className="text-xs font-mono text-muted-foreground">
              HTTP Request
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
              <TabsTrigger value="body" className="font-mono text-xs">
                Body
              </TabsTrigger>
              <TabsTrigger value="other" className="font-mono text-xs">
                Other
              </TabsTrigger>
            </TabsList>
            <TabsContent value="general">
              <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)}>
                  <div className="flex flex-col p-2">
                    <div className="flex flex-col gap-2">
                      <FormField
                        control={form.control}
                        name="name"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Name</FormLabel>
                            <FormControl>
                              <Input placeholder="Name" {...field} />
                            </FormControl>

                            <FormMessage />
                          </FormItem>
                        )}
                      />

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
                        name="path"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Url Path</FormLabel>
                            <FormControl>
                              <Input placeholder="/index" {...field} />
                            </FormControl>

                            <FormMessage />
                          </FormItem>
                        )}
                      />

                      <FormField
                        control={form.control}
                        name="method"
                        render={({ field }) => (
                          <FormItem>
                            <Label>Url Method</Label>
                            <SelectMethodDialog
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
                        name="is_public"
                        render={({ field }) => (
                          <FormItem className="flex flex-row items-center justify-between">
                            <div className="">
                              <FormLabel>Go public</FormLabel>
                            </div>
                            <FormControl>
                              <Switch
                                checked={field.value}
                                onCheckedChange={field.onChange}
                              />
                            </FormControl>
                          </FormItem>
                        )}
                      />
                    </div>
                  </div>
                  <div className="bg-background border-t border-input absolute w-full bottom-0 p-2">
                    <div className="flex items-center justify-end gap-2">
                      <Button
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

export function SelectMethodDialog({
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
          {value ? methods.find((item) => item === value) : "Select method..."}
          <CaretDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0">
        <Command>
          <CommandInput placeholder="Search method..." />
          <CommandList>
            <CommandEmpty>No method found.</CommandEmpty>
            <CommandGroup>
              {methods.map((method) => (
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

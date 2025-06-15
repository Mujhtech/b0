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
import { X } from "@phosphor-icons/react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs";
import { Input } from "~/components/ui/input";

const formSchema = z.object({
  name: z.string().optional(),
  value: z.string().optional(),
});

export default function VariableConnector({
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
      value: workflow?.value as string,
      name: workflow?.name as string,
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
      onUpdate?.({
        ...workflow,
        value: values.value,
        name: values.name,
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
                Variable
              </h3>
            </div>
            <div className="p-2">
              {workflow && (
                <div className="flex flex-col justify-center items-center">
                  <p className="text-xs font-mono text-muted-foreground">
                    {workflow.value
                      ? (workflow.value as string)
                      : workflow.instruction}
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
              Variable
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
                        name="value"
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Value</FormLabel>
                            <FormControl>
                              <Textarea
                                placeholder="Value"
                                className="resize-none"
                                {...field}
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

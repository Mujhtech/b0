import { MetaFunction } from "@remix-run/node";
import { Outlet } from "@remix-run/react";
import React from "react";
import Paragraph from "~/components/ui/paragraph";
import { useUser } from "~/hooks/use-user";
import { Card, CardContent, CardHeader } from "~/components/ui/card";
import DeleteProjectDialog from "~/components/projects/delete-project-dialog";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button } from "~/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "~/components/ui/form";
import { Input } from "~/components/ui/input";
import { useProject } from "~/hooks/use-project";
import { Textarea } from "~/components/ui/textarea";
import { Plus } from "@phosphor-icons/react";

export const meta: MetaFunction = () => {
  return [
    {
      title: `Environment Variables`,
    },
  ];
};

const env = z.object({
  key: z.string(),
  value: z.string(),
  note: z.string().optional(),
  protected: z.boolean().optional(),
});

const formSchema = z.object({
  envs: z.array(env),
});

export default function Page() {
  const project = useProject();
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      envs: [{ key: "", value: "" }],
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
    } catch (error) {
      console.error(error);
    }
  }

  const addEnv = () => {
    form.setValue("envs", [...form.getValues("envs"), { key: "", value: "" }]);
  };

  return (
    <div className="flex flex-col w-full max-w-4xl mx-auto">
      <Card className="mb-8">
        <CardHeader className="flex flex-col p-4">
          <h1 className="font-semibold">Environment Variables</h1>
        </CardHeader>
        <CardContent className="">
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)}>
              <div className="flex flex-col p-4 gap-4">
                {form.watch("envs").map((_, index) => {
                  return (
                    <div className="flex flex-col gap-2">
                      <div className="grid grid-cols-2 gap-2">
                        <FormField
                          control={form.control}
                          name={`envs.${index}.key`}
                          render={({ field }) => (
                            <FormItem>
                              <FormLabel>Key</FormLabel>
                              <FormControl>
                                <Input placeholder="B0_PORT" {...field} />
                              </FormControl>

                              <FormMessage />
                            </FormItem>
                          )}
                        />
                        <FormField
                          control={form.control}
                          name={`envs.${index}.value`}
                          render={({ field }) => (
                            <FormItem>
                              <FormLabel>Value</FormLabel>
                              <FormControl>
                                <Input placeholder="5000" {...field} />
                              </FormControl>

                              <FormMessage />
                            </FormItem>
                          )}
                        />
                      </div>
                      <FormField
                        control={form.control}
                        name={`envs.${index}.note`}
                        render={({ field }) => (
                          <FormItem>
                            <FormLabel>Note</FormLabel>
                            <FormControl>
                              <Textarea
                                placeholder="Note"
                                {...field}
                                className="resize-none h-[50px]"
                              />
                            </FormControl>

                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>
                  );
                })}
              </div>
              <div className="bg-background border-t border-input w-full p-4">
                <div className="flex items-center justify-between gap-2">
                  <div>
                    <Button
                      onClick={addEnv}
                      type="button"
                      variant={"outline"}
                      className="h-8 shadow-none"
                    >
                      <Plus className="!h-4 !w-4" />
                      Add
                    </Button>
                  </div>
                  <Button disabled={true} className="h-8 shadow-none">
                    Save
                  </Button>
                </div>
              </div>
            </form>
          </Form>
        </CardContent>
      </Card>
    </div>
  );
}

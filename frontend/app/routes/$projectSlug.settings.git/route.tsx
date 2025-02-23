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
import { UserAvatar } from "~/components/menus/user-menu";
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

export const meta: MetaFunction = () => {
  return [
    {
      title: `Git`,
    },
  ];
};

const formSchema = z.object({
  name: z.string().optional(),
  description: z.string().optional(),
});

export default function Page() {
  const project = useProject();
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: project.name,
      description: project.description,
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    try {
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <div className="flex flex-col w-full max-w-4xl mx-auto">
      <Card className="mb-8">
        <CardHeader className="flex flex-col p-4">
          <h1 className="font-semibold">Git Repository</h1>
        </CardHeader>
        <CardContent className="">
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)}>
              <div className="flex flex-col p-4">
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
                    name="name"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Description</FormLabel>
                        <FormControl>
                          <Textarea
                            placeholder="Description"
                            {...field}
                            className="resize-none h-[100px]"
                          />
                        </FormControl>

                        <FormMessage />
                      </FormItem>
                    )}
                  />
                </div>
              </div>
              <div className="bg-background border-t border-input w-full p-4">
                <div className="flex items-center justify-end gap-2">
                  <Button disabled={true} className="h-8 shadow-lg">
                    Save
                  </Button>
                </div>
              </div>
            </form>
          </Form>
        </CardContent>
      </Card>
      <Card className="border-red-500 mb-8">
        <CardHeader className="flex flex-col px-4 pt-4">
          <h1 className="font-semibold">Delete Project</h1>
          <Paragraph>Are you sure you want to delete this project?</Paragraph>
        </CardHeader>
        <CardContent className="flex justify-between px-4 py-2 mt-4 bg-red-500 border-t border-red-500 bg-opacity-20">
          <div></div>
          <DeleteProjectDialog />
        </CardContent>
      </Card>
    </div>
  );
}

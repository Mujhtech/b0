import { MetaFunction } from "@remix-run/node";
import { Card, CardContent, CardHeader } from "~/components/ui/card";

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
      title: `Domains`,
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
        <CardHeader className="flex flex-row justify-between items-center p-4">
          <h1 className="font-semibold">Domains</h1>
          <Button className="">Add Domain</Button>
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
    </div>
  );
}

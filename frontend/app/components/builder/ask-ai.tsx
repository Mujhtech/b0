import { PaperPlaneTilt } from "@phosphor-icons/react";
import React from "react";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { AIModelPicker } from "./model-picker";
import { useFeature } from "~/hooks/use-feature";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { clientApi } from "~/services/api.client";
import { useAuthToken, usePlatformUrl } from "~/hooks/use-user";
import { useProject } from "~/hooks/use-project";
import { Form, FormControl, FormField, FormItem } from "../ui/form";
import { ServerResponse, ServerResponseSchema } from "~/models/default";

const chatSchema = z.object({
  text: z.string().min(1),
  model: z.string().min(1),
});

type Chat = z.infer<typeof chatSchema>;

export default function AskB0({
  isThinking,
  projectModel,
}: {
  isThinking: boolean;
  projectModel?: string;
}) {
  const { available_models } = useFeature();
  const defaultModel = available_models.find((model) => model.is_default);
  const [isLoading, setIsLoading] = React.useState(false);

  const accessToken = useAuthToken();
  const backendBaseUrl = usePlatformUrl();
  const project = useProject();

  const form = useForm<Chat>({
    resolver: zodResolver(chatSchema),
    defaultValues: {
      model: projectModel ?? defaultModel?.model!,
      text: "",
    },
  });

  async function onSubmit(values: Chat) {
    try {
      if (!accessToken || !backendBaseUrl || isThinking == true || isLoading) {
        return;
      }
      setIsLoading(true);
      await clientApi.post<ServerResponse>({
        path: "/chat/" + project.id,
        body: { ...values },
        backendBaseUrl: backendBaseUrl!,
        accessToken: accessToken!,
        schema: ServerResponseSchema,
      });
    } catch (error) {
      console.error(error);
    } finally {
      form.reset();
      setIsLoading(false);
    }
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <div className="h-8 border border-input bg-background shadow-lg flex">
          <FormField
            control={form.control}
            name="model"
            render={({ field }) => (
              <FormItem>
                <FormControl>
                  <AIModelPicker
                    onSelect={(val) => field.onChange(val)}
                    value={field.value}
                  />
                </FormControl>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="text"
            render={({ field }) => (
              <FormItem>
                <FormControl>
                  <Input
                    className="h-8 border-none focus-visible:ring-0 w-72 px-2"
                    placeholder="Ask b0 anything..."
                    value={field.value}
                    onChange={field.onChange}
                    onBlur={field.onBlur}
                  />
                </FormControl>

                {/* <FormMessage /> */}
              </FormItem>
            )}
          />

          <Button
            className="!h-8 border-l border-t-0 border-r-0 border-b-0 px-2 bg-transparent"
            variant={"outline"}
            disabled={isThinking == true || isLoading}
          >
            <PaperPlaneTilt className="h-4 w-4" />
          </Button>
        </div>
      </form>
    </Form>
  );
}

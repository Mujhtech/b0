import { ArrowUpRight, PaperPlaneTilt } from "@phosphor-icons/react";
import {
  redirect,
  type ActionFunction,
  type MetaFunction,
} from "@remix-run/node";
import { parseWithZod } from "@conform-to/zod";
import { AIModelPicker2 } from "~/components/builder/model-picker";
import { Button } from "~/components/ui/button";
import { Textarea } from "~/components/ui/textarea";
import { requireUser } from "~/services/user.server";
import { CreateProjectFormSchema } from "~/models/project";
import { createProject } from "~/services/project.server";
import { useFetcher } from "@remix-run/react";
import { useForm, useInputControl } from "@conform-to/react";
import React from "react";
import { useFeature } from "~/hooks/use-feature";
import { LanguagePicker } from "~/components/builder/language-picker";
import { useOptionalUser } from "~/hooks/use-user";
import UserMenu from "~/components/menus/user-menu";
import { redirectBackWithErrorMessage } from "~/models/message.server";

export const action: ActionFunction = async ({ request, params }) => {
  const user = await requireUser(request);

  const formData = await request.formData();

  const submission = parseWithZod(formData, {
    schema: CreateProjectFormSchema,
  });

  if (submission.status !== "success") {
    return redirectBackWithErrorMessage(request, "Invalid form submission");
  }

  try {
    // handle creation of project
    const project = await createProject(request, submission.value);

    let headers = new Headers({});

    return redirect(`/${project.data.slug}`, {
      headers,
    });
  } catch (e: any) {
    return redirectBackWithErrorMessage(request, e.error ?? "Unknown error");
  }
};

export const meta: MetaFunction = () => {
  return [
    { title: "b0 - Your AI backend builder" },
    { name: "description", content: "Build backend in 1 minute with b0" },
  ];
};

export default function Index() {
  const fetcher = useFetcher();
  const { available_models } = useFeature();
  const defaultModel = available_models.find(
    (model) => model.is_default && model.is_enabled
  );
  const [model, setModel] = React.useState<string | undefined>(
    defaultModel?.model
  );
  const [language, setLanguage] = React.useState<string | undefined>("1");
  const user = useOptionalUser();

  const [form, fields] = useForm({
    id: "create-project-form",
    shouldValidate: "onBlur",
    shouldRevalidate: "onSubmit",
    onValidate({ formData }) {
      return parseWithZod(formData, {
        schema: CreateProjectFormSchema,
      });
    },
  });

  return (
    <main className="relative h-full w-full flex flex-col flex-1">
      {/* <AnimatedGradient
        speed={0.05}
        blur="medium"
        colors={["#0E0E0E", "#0E0E0E"]}
      /> */}
      <div className="z-10 flex flex-col w-full max-w-[50rem] items-center mx-auto pt-56">
        <h1 className="font-sans text-3xl font-black">
          b0, Your AI backend builder
        </h1>
        <div className="mt-4 flex flex-col gap-4 w-full">
          <fetcher.Form
            onSubmit={form.onSubmit}
            method="post"
            className="w-full border border-input bg-background shadow-lg"
          >
            <Textarea
              name={fields.prompt.name}
              value={fields.prompt.value}
              key={fields.prompt.key}
              placeholder="What are you building today?"
              className=" resize-none border-none focus-visible:ring-0 min-h-[100px] shadow-none"
            ></Textarea>
            <input
              type="hidden"
              name={fields.model.name}
              key={fields.model.key}
              value={model}
            />
            <input
              type="hidden"
              name={fields.framework_id.name}
              key={fields.framework_id.key}
              value={language}
            />
            <div className="flex items-center px-3 pb-2">
              <div className="flex gap-2">
                <AIModelPicker2
                  value={model}
                  onSelect={(mode) => {
                    setModel(mode);
                  }}
                />
                <LanguagePicker value={language} onSelect={setLanguage} />
              </div>
              <div className="ml-auto flex items-center gap-2">
                <button
                  type="submit"
                  className="border border-input h-6 w-6 p-1 inline-flex items-center justify-center"
                >
                  <PaperPlaneTilt className="h-5 w-5" />
                </button>
              </div>
            </div>
          </fetcher.Form>
          <div className="flex items-center justify-center">
            <p className="font-mono text-[10px] text-muted-foreground text-center">
              b0 can make mistakes. Please double-check it.
            </p>
          </div>
        </div>
        <div className="mt-12 flex flex-wrap justify-center gap-3 w-full">
          {[
            { title: "Hello world", prompt: "Build a simple hello world app" },
            {
              title: "Stripe webhook",
              prompt:
                "Build a web server that handles and process stripe webhooks",
            },
            {
              title: "Telegram Bot",
              prompt:
                "Build a web server that listen to telegram bot webhook and reply to it using telegram bot api",
            },
            {
              title: "Discord Bot",
              prompt:
                "Build a discord bot web server that listen to discord bot webhook and reply to it using discordjs sdk",
            },
            {
              title: "Open AI",
              prompt: "Build a micro service that integrate with open ai",
            },
          ].map((template, i) => (
            <TemplateCard
              key={i}
              template={template.title}
              prompt={template.prompt}
              model={model}
              language={language}
            />
          ))}
        </div>
      </div>
      <div className="flex flex-col justify-end items-start w-full h-full px-4 pb-4">
        {user && <UserMenu user={user} />}
      </div>
    </main>
  );
}

const TemplateCard = ({
  template,
  prompt,
  model,
  language,
}: {
  template: string;
  prompt: string;
  model?: string;
  language?: string;
}) => {
  const fetcher = useFetcher();

  const handleOnClick = () => {
    const formData = new FormData();
    formData.append("prompt", prompt);
    formData.append("isTemplate", "true");
    if (model) {
      formData.append("model", model);
    }
    if (language) {
      formData.append("framework_id", language);
    }
    fetcher.submit(formData, { method: "POST" });
  };

  return (
    <Button
      variant="outline"
      type="button"
      onClick={handleOnClick}
      className="border border-input shadow-lg"
    >
      {template} <ArrowUpRight className="!h-4 !w-4" />
    </Button>
  );
};

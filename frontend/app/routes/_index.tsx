import {
  ArrowUpRight,
  Folder,
  GlobeHemisphereWest,
  PaperPlaneTilt,
} from "@phosphor-icons/react";
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

export const action: ActionFunction = async ({ request, params }) => {
  const user = await requireUser(request);

  const formData = await request.formData();

  const submission = parseWithZod(formData, {
    schema: CreateProjectFormSchema,
  });

  if (submission.status !== "success") {
    return redirect("/");
  }

  try {
    // handle creation of project
    const project = await createProject(request, submission.value);

    let headers = new Headers({});

    return redirect(`/${project.data.slug}`, {
      headers,
    });
  } catch (e) {
    return redirect("/");
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
  const defaultModel = available_models.find((model) => model.is_default);
  const [model, setModel] = React.useState<string | undefined>(
    defaultModel?.model
  );

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
            <div className="flex items-center px-3 pb-2">
              <div className="flex gap-2">
                <AIModelPicker2
                  value={model}
                  onSelect={(mode) => {
                    setModel(mode);
                  }}
                />
                <button
                  type="button"
                  className="border border-input h-6 w-6 p-1 inline-flex items-center justify-center"
                >
                  <Folder className="h-5 w-5" />
                </button>
                <button
                  type="button"
                  className="border border-input h-6 w-6 p-1 inline-flex items-center justify-center"
                >
                  <GlobeHemisphereWest className="h-5 w-5" />
                </button>
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
            "Hello world",
            "Stripe webhook",
            "Telegram Bot",
            "Discord Bot",
            "Open AI",
          ].map((template, i) => (
            <TemplateCard key={i} template={template} model={model} />
          ))}
        </div>
      </div>
    </main>
  );
}

const TemplateCard = ({
  template,
  model,
}: {
  template: string;
  model?: string;
}) => {
  const fetcher = useFetcher();

  const handleOnClick = () => {
    const formData = new FormData();
    formData.append("prompt", template);
    formData.append("isTemplate", "true");
    if (model) {
      formData.append("model", model);
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

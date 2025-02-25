import {
  Dialog,
  DialogContent,
  DialogTitle,
  DialogTrigger,
} from "~/components/ui/dialog";
import { Button } from "~/components/ui/button";
import { useFetcher, useLocation, useNavigation } from "@remix-run/react";
import { Label } from "~/components/ui/label";
import { Input } from "~/components/ui/input";
import { Textarea } from "~/components/ui/textarea";
import FormField from "~/components/ui/form-field";
import Paragraph from "~/components/ui/paragraph";
import { useState } from "react";
import { useForm, useInputControl } from "@conform-to/react";
import { parseWithZod } from "@conform-to/zod";
import FormError from "~/components/ui/form-error";
import { z } from "zod";
import { useProject } from "~/hooks/use-project";

const formSchema = z.object({
  name: z.string().min(1),
});

export default function DeleteProjectDialog() {
  const [isOpen, setIsOpen] = useState(false);
  const project = useProject();

  const fetcher = useFetcher();

  const [form, fields] = useForm({
    id: "delete-project",
    shouldValidate: "onBlur",
    shouldRevalidate: "onSubmit",

    onValidate({ formData }) {
      return parseWithZod(formData, {
        schema: formSchema,
      });
    },
  });

  return (
    <Dialog open={isOpen} onOpenChange={(value) => setIsOpen(value)}>
      <DialogTrigger asChild>
        <Button className="!bg-red-500 text-white shadow-none">Delete</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogTitle>Delete Project</DialogTitle>
        <fetcher.Form
          method="post"
          className="grid grid-cols-1 gap-3"
          onSubmit={form.onSubmit}
        >
          <FormField>
            <Input
              key={fields.name.key}
              name={fields.name.name}
              placeholder="Project Name"
              defaultValue={fields.name.initialValue || ""}
            />
            <FormError>{fields.name.errors}</FormError>
          </FormField>

          <div>
            <Paragraph>
              Type in the name of the project{" "}
              <code className="py-0.5 px-1 border-border border">
                {project.name.toLowerCase()}
              </code>
              . Note, this action is irreversible.
            </Paragraph>
          </div>

          <div className="flex gap-2">
            <Button
              variant="outline"
              type="button"
              onClick={() => setIsOpen(false)}
            >
              Close
            </Button>

            <Button
              type="submit"
              name="intent"
              value="create"
              className="!bg-red-500"
            >
              Delete
            </Button>
          </div>
        </fetcher.Form>
      </DialogContent>
    </Dialog>
  );
}

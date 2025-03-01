import {
  type LoaderFunctionArgs,
  redirect,
  type ActionFunction,
} from "@remix-run/node";
import { parseWithZod } from "@conform-to/zod";
import { requireUser } from "~/services/user.server";
import {
  redirectWithSuccessMessage,
  redirectWithErrorMessage,
} from "~/models/message.server";
import { z } from "zod";
import invariant from "tiny-invariant";
import { deleteProject } from "~/services/project.server";

const formSchema = z.object({
  name: z.string().min(1),
});

export const ProjectIdParamSchema = z.object({
  projectId: z.string(),
});

export const action: ActionFunction = async ({ request, params }) => {
  await requireUser(request);

  const { projectId } = ProjectIdParamSchema.parse(params);

  invariant(projectId, "No project found in request.");

  const formData = await request.formData();

  const submission = parseWithZod(formData, {
    schema: formSchema,
  });

  if (submission.status !== "success") {
    return redirectWithErrorMessage(`/`, request, "Invalid form submission");
  }

  try {
    // handle creation of project
    await deleteProject(request, projectId, submission.value);

    return redirectWithSuccessMessage(
      `/`,
      request,
      "Project deleted successfully"
    );
  } catch (e: any) {
    return redirectWithErrorMessage(`/`, request, e.error ?? "Unknown error");
  }
};

export const loader = async ({ request, params }: LoaderFunctionArgs) => {
  const { projectId } = ProjectIdParamSchema.parse(params);

  invariant(projectId, "No project found in request.");

  return redirect("/");
};

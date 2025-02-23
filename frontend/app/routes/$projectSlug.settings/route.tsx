import { LoaderFunctionArgs, MetaFunction } from "@remix-run/node";
import { Outlet } from "@remix-run/react";
import React from "react";
import {
  typedjson,
  UseDataFunctionReturn,
  useTypedLoaderData,
} from "remix-typedjson";
import UserMenu from "~/components/menus/user-menu";
import { useUser } from "~/hooks/use-user";
import { SettingMenuNavItem } from "../settings/route";
import { ProjectSlugParamSchema } from "../$projectSlug/route";
import invariant from "tiny-invariant";

export const loader = async ({ request, params }: LoaderFunctionArgs) => {
  const { projectSlug } = ProjectSlugParamSchema.parse(params);

  invariant(projectSlug, "No project found in request.");

  return typedjson({
    projectSlug,
  });
};

export default function Page() {
  const user = useUser();
  const { projectSlug } = useTypedLoaderData<typeof loader>();

  return (
    <div className="flex flex-col">
      <div className="flex flex-col mb-3 px-4">
        <div className="flex items-center border-border border-b pt-4 pb-2 justify-between">
          <h1 className="text-2xl font-bold">Settings</h1>
          <UserMenu user={user} showHomepage={true} className="ml-auto mr-4" />
        </div>
      </div>
      <SettingMenuNav projectSlug={projectSlug} />
      <Outlet />
    </div>
  );
}

const SettingMenuNav = ({ projectSlug }: { projectSlug: string }) => {
  return (
    <nav className="max-w-md mb-4">
      <ul className="flex gap-3 px-4">
        <SettingMenuNavItem title="General" to={`/${projectSlug}/settings`} />
        <SettingMenuNavItem
          title="Domains"
          to={`/${projectSlug}/settings/domains`}
        />
        <SettingMenuNavItem
          title="Environment Variables"
          to={`/${projectSlug}/settings/envs`}
        />
        <SettingMenuNavItem title="Git" to={`/${projectSlug}/settings/git`} />
      </ul>
    </nav>
  );
};

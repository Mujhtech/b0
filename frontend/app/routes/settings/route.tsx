import { LoaderFunctionArgs } from "@remix-run/node";
import { Link, Outlet, useLocation } from "@remix-run/react";
import React from "react";
import { typedjson } from "remix-typedjson";
import UserMenu from "~/components/menus/user-menu";
import { useUser } from "~/hooks/use-user";
import { cn } from "~/lib/utils";
import { requireUser } from "~/services/user.server";

export const loader = async ({ request, params }: LoaderFunctionArgs) => {
  requireUser(request);

  return typedjson({});
};

export default function Setting() {
  const user = useUser();
  return (
    <div className="flex flex-col">
      <div className="flex flex-col mb-3 px-4">
        <div className="flex items-center border-border border-b pt-4 pb-2 justify-between">
          <h1 className="text-2xl font-bold">Settings</h1>
          <UserMenu user={user} className="ml-auto mr-4" />
        </div>
      </div>
      <SettingMenuNav />
      <Outlet />
    </div>
  );
}

const SettingMenuNav = () => {
  return (
    <nav className="max-w-md mb-4">
      <ul className="flex gap-3 px-4">
        <SettingMenuNavItem title="General" to={"/settings"} />
        <SettingMenuNavItem title="Usage" to={"/settings/usage"} />
        <SettingMenuNavItem title="Billing" to={"/settings/billing"} />
      </ul>
    </nav>
  );
};

export const SettingMenuNavItem = ({
  title,
  to,
}: {
  title: string;
  to: string;
}) => {
  const location = useLocation();

  const isActive = location.pathname == to;

  return (
    <li>
      <Link
        to={to}
        className={cn(
          "text-sm ",
          isActive ? "text-white font-medium" : "text-muted-foreground"
        )}
      >
        {title}
      </Link>
    </li>
  );
};

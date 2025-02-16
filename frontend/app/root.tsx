import {
  Links,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "@remix-run/react";
import type { LinksFunction, LoaderFunctionArgs } from "@remix-run/node";
import { typedjson } from "remix-typedjson";
import "./tailwind.css";
import { env } from "./env.server";
import { getAuthTokenFromSession } from "./services/auth.server";
import { getFeatures } from "./services/feature.server";
import { getUser } from "./services/user.server";
import { User } from "./models/user";
import { Feature } from "./models/feature";
import { Projects } from "./models/project";
import { getProjects } from "./services/project.server";
import {
  commitSession,
  getSession,
  ToastMessage,
} from "./models/message.server";
import { Toast } from "./components/custom-toast";

export const links: LinksFunction = () => [];

export const loader = async ({ request }: LoaderFunctionArgs) => {
  const backendUrl = env.BACKEND_URL;
  const platformUrl = env.PLATFORM_URL;
  let user: User | null = null;
  let feature: Feature | null = null;
  let accessToken: string | null = null;
  let projects: Projects = [];

  const session = await getSession(request.headers.get("cookie"));
  const toastMessage = session.get("toastMessage") as ToastMessage;

  try {
    feature = await getFeatures(request);
  } catch (e) {
    //
  }

  try {
    accessToken = await getAuthTokenFromSession(request);
  } catch (e) {
    //
  }

  try {
    user = await getUser(request);

    projects = await getProjects(request);
  } catch (e) {
    //
  }

  return typedjson(
    {
      user: user,
      feature: feature,
      accessToken: accessToken,
      backendUrl,
      platformUrl,
      toastMessage,
      projects,
    },
    { headers: { "Set-Cookie": await commitSession(session) } }
  );
};

export type RootLoaderType = typeof loader;

export function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" className="h-full" suppressHydrationWarning>
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body
        className="h-full overflow-hidden bg-background text-foreground antialiased !m-0"
        suppressHydrationWarning
      >
        {children}
        <Toast />
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  );
}

export default function App() {
  return <Outlet />;
}

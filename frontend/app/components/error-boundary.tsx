import { isRouteErrorResponse, Link, useRouteError } from "@remix-run/react";
import { motion } from "framer-motion";
import Spline from "@splinetool/react-spline";
import { cn, friendlyErrorDisplay } from "~/lib/utils";
import { LinkButton } from "./ui/button";

type ErrorDisplayOptions = {
  button?: {
    title: string;
    to: string;
  };
  className?: string;
};

export function RouteErrorDisplay(options?: ErrorDisplayOptions) {
  const error = useRouteError();

  return (
    <>
      {isRouteErrorResponse(error) ? (
        <ErrorDisplay
          title={friendlyErrorDisplay(error.status, error.statusText).title}
          message={
            error.data.message ??
            friendlyErrorDisplay(error.status, error.statusText).message
          }
          {...options}
        />
      ) : error instanceof Error ? (
        <ErrorDisplay title={error.name} message={error.message} {...options} />
      ) : (
        <ErrorDisplay
          title="Oops"
          message={JSON.stringify(error)}
          {...options}
        />
      )}
    </>
  );
}

type DisplayOptionsProps = {
  title: string;
  message?: string;
} & ErrorDisplayOptions;

export function ErrorDisplay({
  title,
  message,
  button,
  className,
}: DisplayOptionsProps) {
  return (
    <main className={cn("w-full h-full relative", className)}>
      <div className="flex flex-col justify-center gap-1 items-center h-full w-full">
        <h1 className="font-sans text-3xl">{title}</h1>
        {message && <p className="font-mono">{message}</p>}
        <LinkButton
          variant={"default"}
          to={button ? button.to : "/"}
          className=""
        >
          {button ? button.title : "Go to homepage"}
        </LinkButton>
      </div>
    </main>
  );
}

import React from "react";
import { cn } from "~/lib/utils";

export default function Paragraph({
  className,
  children,
  onClick,
}: {
  className?: string;
  children: React.ReactNode;
  onClick?: () => void;
}) {
  return (
    <p
      className={cn("text-xs text-muted-foreground", className)}
      onClick={onClick}
    >
      {children}
    </p>
  );
}

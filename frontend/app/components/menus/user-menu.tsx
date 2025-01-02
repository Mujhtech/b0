import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "~/components/ui/dropdown-menu";

import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar";
import { Link } from "@remix-run/react";
import { logoutPath } from "~/lib/path";
import { cn } from "~/lib/utils";

export default function UserMenu() {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="focus:outline-none focus-visible:ring-0">
        <UserAvatar />
      </DropdownMenuTrigger>
      <DropdownMenuContent className={cn("ml-4")}>
        <DropdownMenuLabel>My Account</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem>Profile</DropdownMenuItem>
        <DropdownMenuItem>Billing</DropdownMenuItem>
        <DropdownMenuItem>
          <Link to={logoutPath()}>Logout</Link>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

export const UserAvatar = () => {
  return (
    <Avatar className="h-8 w-8 rounded-none p-1 flex items-center md:justify-center border border-border bg-background hover:bg-background hover:border-border hover:border">
      <AvatarImage src={"https://github.com/mattmalin.png"} />
      <AvatarFallback>MM</AvatarFallback>
    </Avatar>
  );
};

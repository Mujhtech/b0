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
import { useUser } from "~/hooks/use-user";
import { User } from "~/models/user";
import { useState } from "react";
import { ProjectsDialog } from "../projects/project-dialog";

export default function UserMenu() {
  const user = useUser();
  const [openProjectDialog, setOpenProjectDialog] = useState(false);
  const handleOpenProjectDialog = () => {
    setOpenProjectDialog(!openProjectDialog);
  };
  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger className="focus:outline-none focus-visible:ring-0">
          <UserAvatar user={user} />
        </DropdownMenuTrigger>
        <DropdownMenuContent className={cn("ml-4")}>
          <DropdownMenuLabel>My Account</DropdownMenuLabel>
          <DropdownMenuSeparator />
          <DropdownMenuItem onClick={handleOpenProjectDialog}>
            Projects
          </DropdownMenuItem>
          <DropdownMenuItem>Profile</DropdownMenuItem>
          <DropdownMenuItem>Billing</DropdownMenuItem>
          <DropdownMenuItem>
            <Link to={logoutPath()}>Logout</Link>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      <ProjectsDialog open={openProjectDialog} setOpen={setOpenProjectDialog} />
    </>
  );
}

export const UserAvatar = ({ user }: { user: User }) => {
  return (
    <Avatar className="h-8 w-8 rounded-none p-1 flex items-center md:justify-center border border-border bg-background hover:bg-background hover:border-border hover:border">
      {user.avatar_url && <AvatarImage src={user?.avatar_url} />}
      <AvatarFallback>{user.name[0]}</AvatarFallback>
    </Avatar>
  );
};

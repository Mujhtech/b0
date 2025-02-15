import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "~/components/ui/dropdown-menu";

import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar";
import { Link, useNavigate } from "@remix-run/react";
import { logoutPath } from "~/lib/path";
import { cn } from "~/lib/utils";
import { User } from "~/models/user";
import { useState } from "react";
import { ProjectsDialog } from "../projects/project-dialog";

export default function UserMenu({
  user,
  className,
  showHomepage,
}: {
  user: User;
  className?: string;
  showHomepage?: boolean;
}) {
  const [openProjectDialog, setOpenProjectDialog] = useState(false);
  const handleOpenProjectDialog = () => {
    setOpenProjectDialog(!openProjectDialog);
  };

  const navigate = useNavigate();
  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger className="focus:outline-none focus-visible:ring-0">
          <UserAvatar user={user} />
        </DropdownMenuTrigger>
        <DropdownMenuContent className={cn("ml-4", className)}>
          <DropdownMenuLabel>My Account</DropdownMenuLabel>
          <DropdownMenuSeparator />
          {showHomepage == true && (
            <DropdownMenuItem
              onClick={() => navigate("/")}
              className="cursor-pointer"
            >
              Homepage
            </DropdownMenuItem>
          )}
          <DropdownMenuItem
            onClick={handleOpenProjectDialog}
            className="cursor-pointer"
          >
            Projects
          </DropdownMenuItem>
          <DropdownMenuItem
            onClick={() => navigate("/settings")}
            className="cursor-pointer"
          >
            Setting
          </DropdownMenuItem>
          <DropdownMenuItem
            onClick={() => navigate("/settings/billing")}
            className="cursor-pointer"
          >
            Billing
          </DropdownMenuItem>
          <DropdownMenuItem className="cursor-pointer">
            <Link to={logoutPath()}>Logout</Link>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      <ProjectsDialog open={openProjectDialog} setOpen={setOpenProjectDialog} />
    </>
  );
}

export const UserAvatar = ({
  user,
  className,
}: {
  user: User;
  className?: string;
}) => {
  return (
    <Avatar
      className={cn(
        "h-8 w-8 rounded-none p-1 flex items-center md:justify-center border border-border bg-background hover:bg-background hover:border-border hover:border",
        className
      )}
    >
      {user.avatar_url && <AvatarImage src={user?.avatar_url} />}
      <AvatarFallback>{user.name[0]}</AvatarFallback>
    </Avatar>
  );
};

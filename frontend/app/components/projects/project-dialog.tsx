import { ArrowUpRight } from "@phosphor-icons/react";
import { Link } from "@remix-run/react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "~/components/ui/dialog";
import { useOptionalProjects } from "~/hooks/use-project";

export function ProjectsDialog({
  open,
  setOpen,
}: {
  open: boolean;
  setOpen: (open: boolean) => void;
}) {
  const projects = useOptionalProjects();
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      {/* <DialogTrigger asChild>
        <Button variant="outline">Edit Profile</Button>
      </DialogTrigger> */}
      <DialogContent className="sm:max-w-[90%] h-[90%] flex flex-col">
        <DialogHeader>
          <DialogTitle>Projects</DialogTitle>
          <DialogDescription className="hidden">
            View and manage your projects
          </DialogDescription>
        </DialogHeader>
        <div className="grid grid-cols-4 gap-4">
          {projects && projects.length > 0 ? (
            <>
              {projects.map((project) => {
                return (
                  <Link to={`/${project.slug}`} key={project.id}>
                    <div className="border border-input p-2 items-center h-full flex gap-1">
                      <p className="font-mono text-sm line-clamp-2 mr-auto">
                        {project.name}
                      </p>
                      <ArrowUpRight className="h-5 w-5 flex-shrink-0" />
                    </div>
                  </Link>
                );
              })}
            </>
          ) : (
            <div></div>
          )}
        </div>
        <DialogFooter></DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

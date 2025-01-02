import { GithubLogo, GoogleLogo } from "@phosphor-icons/react";
import React from "react";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";

export default function Page() {
  return (
    <main className="h-full text-white">
      <div className="relative h-full ">
        <div className="absolute inset-0 w-full h-full z-20 [mask-image:radial-gradient(transparent,white)] pointer-events-none" />

        <div className="flex justify-center w-full h-full ">
          <div className="m-auto min-w-[400px] flex-col flex items-center bg-black z-30 ">
            <div className="flex flex-col px-12 py-20 w-full">
              <h1 className="text-4xl font-sans font-bold">Welcome</h1>
              <p className="text-sm">Create an account or login to continue</p>
              <div className="flex flex-col gap-y-3 mt-6">
                <form
                  method="post"
                  action="/login"
                  className="grid grid-cols-1 gap-4"
                >
                  <Input placeholder="Email Address" />
                  <Button
                    className="w-full"
                    type="button"
                    onClick={() => {
                      //   window.location.href = `${data.backendUrl}/ui/auth/github`;
                      //   window.close();
                    }}
                  >
                    Continue
                  </Button>
                </form>

                <>
                  <p className="text-sm text-center">Or</p>

                  <Button
                    className="w-full"
                    type="button"
                    onClick={() => {
                      //   window.location.href = `${data.backendUrl}/ui/auth/github`;
                      //   window.close();
                    }}
                  >
                    <GithubLogo size={20} className="mr-3 h-5 w-5" /> Continue
                    with Github
                  </Button>

                  <Button
                    className="w-full"
                    type="button"
                    onClick={() => {
                      //   window.location.href = `${data.backendUrl}/ui/auth/google`;
                      //   window.close();
                    }}
                  >
                    <GoogleLogo size={20} className="mr-3 h-5 w-5" />
                    Continue with Google
                  </Button>
                </>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}

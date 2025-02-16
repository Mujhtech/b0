/// <reference types="vitest/config" />
import { vitePlugin as remix } from "@remix-run/dev";
import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";

declare module "@remix-run/node" {
  interface Future {
    v3_singleFetch: true;
  }
}

export default defineConfig({
  resolve: {
    /**
     * @uiw/* packages are not compatible with ESM, as they do not import with `.js` extension.
     * We're using source in TypeScript instead.
     */
    alias: {
      "@uiw/react-codemirror": "@uiw/react-codemirror/src/index.tsx",
      // "@uiw/codemirror-extensions-basic-setup":
      //   "@uiw/codemirror-extensions-basic-setup/src/index.ts",
      // "@uiw/codemirror-theme-abyss": "@uiw/codemirror-theme-abyss/src/index.ts",
      // "@uiw/codemirror-themes": "@uiw/codemirror-themes/src/index.tsx",
    },
    conditions: ["bun"],
  },
  optimizeDeps: {
    /**
     * We need to tell Vite to process @uiw/* and related packages, so we transpile them
     * properly.
     */
    include: [
      // "@uiw/codemirror-extensions-basic-setup",
      "@codemirror/state",
      "@codemirror/view",
      "@codemirror/language",
      "@codemirror/autocomplete",
      "@codemirror/lang-json",
      "@codemirror/commands",
      "@codemirror/lint",
      "@codemirror/search",
      // "@uiw/codemirror-theme-abyss",
      // "@uiw/codemirror-themes",
      "@codemirror/lang-javascript",
    ],
  },
  plugins: [
    remix({
      future: {
        v3_fetcherPersist: true,
        v3_relativeSplatPath: true,
        v3_throwAbortReason: true,
        v3_singleFetch: true,
        v3_lazyRouteDiscovery: true,
      },
    }),
    tsconfigPaths(),
  ],
  test: {
    restoreMocks: true,
  },
});

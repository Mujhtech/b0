import { z } from "zod";
import { useAuthToken, useBackendUrl } from "~/hooks/use-user";
import { ClientRequestOptions, handleClientUnauthorized } from "./api-helper";

const clientRequest = async (options: {
  method: "POST" | "GET" | "PUT" | "DELETE" | "HEAD";
  path: string;
  body: any;
  headers?: any;
  backendBaseUrl?: string;
  accessToken?: string;
}): Promise<Response> => {
  let headers = {
    "Content-Type": "application/json",
    Authorization: `Bearer ${options.accessToken}`,
    ...options.headers,
  };

  const url = `${options.backendBaseUrl}${options.path}`;

  const response = fetch(url, {
    method: options.method,
    headers: headers,
    body:
      options.method == "GET" || options.method == "HEAD"
        ? undefined
        : JSON.stringify({
            ...options.body,
          }),
  });

  return response;
};

async function getRequest<T = any>(
  options: ClientRequestOptions<T>
): Promise<T> {
  return handleClientUnauthorized<T>(
    () => {
      const query = new URLSearchParams(options.query).toString();

      const path = query ? `${options.path}?${query}` : options.path;

      return clientRequest({
        headers: options.headers,
        method: "GET",
        path: path,
        body: {},
        accessToken: options.accessToken,
        backendBaseUrl: options.backendBaseUrl,
      });
    },
    {
      schema: options.schema,
    }
  );
}

async function putRequest<T = any>(
  options: ClientRequestOptions<T>
): Promise<T> {
  return handleClientUnauthorized<T>(
    () => {
      return clientRequest({
        headers: options.headers,
        method: "PUT",
        path: options.path,
        body: options.body,
        accessToken: options.accessToken,
        backendBaseUrl: options.backendBaseUrl,
      });
    },
    {
      schema: options.schema,
    }
  );
}

async function deleteRequest<T = any>(
  options: ClientRequestOptions<T>
): Promise<T> {
  return handleClientUnauthorized<T>(
    () => {
      return clientRequest({
        method: "DELETE",
        path: options.path,
        headers: options.headers,
        body: options.body,
        accessToken: options.accessToken,
        backendBaseUrl: options.backendBaseUrl,
      });
    },
    {
      schema: options.schema,
    }
  );
}

async function postRequest<T = any>(
  options: ClientRequestOptions<T>
): Promise<T> {
  return handleClientUnauthorized<T>(
    () => {
      return clientRequest({
        method: "POST",
        path: options.path,
        body: options.body,
        headers: options.headers,
        accessToken: options.accessToken,
        backendBaseUrl: options.backendBaseUrl,
      });
    },
    {
      schema: options.schema,
    }
  );
}

export const clientApi = {
  post: postRequest,
  get: getRequest,
  put: putRequest,
  delete: deleteRequest,
};

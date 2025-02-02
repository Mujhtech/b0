import { env } from "~/env.server";
import { getAuthSession, getAuthTokenFromSession } from "./auth.server";
import { ClientRequestOptions, handleClientUnauthorized } from "./api-helper";

async function postRequest<T = any>(
  options: ClientRequestOptions<T>
): Promise<T> {
  return handleClientUnauthorized<T>(
    () => {
      return serverRequest({
        method: "POST",
        path: options.path,
        body: options.body,
        request: options.request!,
        headers: options.headers,
      });
    },
    {
      schema: options.schema,
    }
  );
}

async function getRequest<T = any>(
  options: ClientRequestOptions<T>
): Promise<T> {
  return handleClientUnauthorized<T>(
    () => {
      const query = new URLSearchParams(options.query).toString();

      const path = query ? `${options.path}?${query}` : options.path;

      return serverRequest({
        request: options.request!,
        headers: options.headers,
        method: "GET",
        path: path,
        body: {},
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
      return serverRequest({
        request: options.request!,
        headers: options.headers,
        method: "PUT",
        path: options.path,
        body: options.body,
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
      return serverRequest({
        method: "DELETE",
        path: options.path,
        headers: options.headers,
        body: options.body,
        request: options.request!,
      });
    },
    {
      schema: options.schema,
    }
  );
}

const serverRequest = async (options: {
  request: Request;
  method: "POST" | "GET" | "PUT" | "DELETE" | "HEAD";
  path: string;
  body: any;
  headers?: any;
}): Promise<Response> => {
  const accessToken = await getAuthTokenFromSession(options.request);

  let headers = {
    "Content-Type": "application/json",
    Authorization: `Bearer ${accessToken}`,
    ...options.headers,
  };

  const url = `${env.PLATFORM_URL}${options.path}`;

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

export const api = {
  post: postRequest,
  get: getRequest,
  put: putRequest,
  delete: deleteRequest,
};

import { z } from "zod";

export type ClientRequestOptions<T> = {
  request?: Request;
  path: string;
  body?: any;
  query?: any;
  schema?: z.ZodType<T>;
  headers?: any;
  backendBaseUrl?: string;
  accessToken?: string;
};

export async function handleClientUnauthorized<T = any>(
  sendRequest: () => Promise<Response>,
  options: {
    schema?: z.ZodType<T>;
  }
): Promise<T> {
  let response = await sendRequest();

  if (response.status === 401) {
    response = await sendRequest();
  }

  if ([200, 201].includes(response.status) == false) {
    const body = await response.json();

    console.log(body);

    throw body;
  }

  const data = await response.json();

  console.log(data);

  if (options.schema) {
    try {
      return options.schema.parse(data.data);
    } catch (error) {
      return data as T;
    }
  }

  return data as T;
}

export class Request {
  commonHeaders: { [key: string]: string };

  constructor(commonHeaders: { [key: string]: string }) {
    this.commonHeaders = commonHeaders;
  }

  async Raw(
    url: string,
    method: "GET" | "POST" | "PATCH" | "PUT" | "DELETE",
    data?: any,
    headers?: { [key: string]: string }
  ) {
    const response = await fetch(url, {
      method: method,
      headers: { ...this.commonHeaders, ...headers },
      body: data ? JSON.stringify(data) : null,
    });
    return response;
  }

  async Get(url: string, headers?: { [key: string]: string }) {
    return this.Raw(url, "GET", null, headers);
  }
  async Post(
    url: string,
    data = {},
    headers ?: { [key: string]: string }
  ) {
    return this.Raw(url, "POST", data, headers);
  }
  async Delete(url: string, data = {}, headers?: { [key: string]: string }) {
    return this.Raw(url, "DELETE", data, headers);
  }
  async Patch(url: string, data = {}, headers?: { [key: string]: string }) {
    return this.Raw(url, "PATCH", data, headers);
  }
  async Put(url: string, data = {}, headers?: { [key: string]: string }) {
    return this.Raw(url, "PUT", data, headers);
  }
}

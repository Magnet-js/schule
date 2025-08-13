import axios, { AxiosError } from "axios";

export const apiClient = axios.create({
  baseURL: "/api",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

apiClient.interceptors.response.use(
  function (response) {
    return response;
  },
  function (error: AxiosError) {
    let res = error.response;
    if (res?.status == 401) {
      return error;
    }
    return Promise.reject(error);
  }
);

export default { apiClient };

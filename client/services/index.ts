export const FetchMethod = async (url: string, options: RequestInit = {}) => {
  const response = await fetch(url, {
    method: options.method || "GET",
    headers: {
      "Content-Type": "application/json",
    },
    ...options,
  });
  const data = await response.json();
  return data;
};

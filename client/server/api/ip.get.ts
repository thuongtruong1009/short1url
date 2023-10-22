export default defineEventHandler(async (_event) => {
  const config = useRuntimeConfig();
  const ipRes = await fetch(config.IP_RESOLVER as string);
  const ipData = await ipRes.json();
  return ipData;
});

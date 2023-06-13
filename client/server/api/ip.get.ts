export default defineEventHandler(async(event) => {
    const ipRes = await fetch(process.env.NUXT_PUBLIC_IP_RESOLVER as string);
    const ipData = await ipRes.json();
    return ipData;
})
import pkg from "./package.json";

export default defineNuxtConfig({
  app: {
    head: {
      title: pkg.name,
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          hid: "description",
          name: "description",
          content: pkg.description,
        },
        { name: "format-detection", content: "telephone=no" },
      ],
      script: [
        {
          src: "https://cdnjs.cloudflare.com/ajax/libs/qrcodejs/1.0.0/qrcode.min.js",
          defer: true,
        },
      ],
      link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
    },
  },

  vite: {
    css: {
      preprocessorOptions: {
        scss: {
          additionalData:
            '@use "@/assets/_base.scss" as *; @use "@/assets/_abstract.scss" as *;',
        },
      },
    },
  },

  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE,
      appBase: process.env.NUXT_PUBLIC_APP_BASE,
      ipSource: process.env.NUXT_PUBLIC_IP_SOURCE,
    },
  },
});

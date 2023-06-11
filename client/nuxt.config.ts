export default defineNuxtConfig({
  devtools: { enabled: true },

  app: {
    head: {
      title: "Sshort1url",
      meta: [
        { charset: "utf-8" },
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        {
          hid: "description",
          name: "description",
          content:
            "ship your link in an easier way",
        },
        { name: "format-detection", content: "telephone=no" },
      ],
      script: [
        { src: 'https://cdnjs.cloudflare.com/ajax/libs/qrcodejs/1.0.0/qrcode.min.js', defer: true }
      ],
      link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
    },
  },

  vite: {
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: '@use "@/assets/_base.scss" as *; @use "@/assets/_abstract.scss" as *;'
          
        }
      }
    }
  }
})

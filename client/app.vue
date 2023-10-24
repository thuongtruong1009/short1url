<script setup lang="ts">
import { onMounted, ref } from "vue";
import BarCode from "./components/BarCode.vue";
import Footer from "./components/Footer.vue";
import Header from "./components/Header.vue";
import QRCode from "./components/QRCode.vue";
import ShortedList from "./components/ShortedList.vue";
import Bar from "./components/icons/Bar.vue";
import Link from "./components/icons/Link.vue";
import QR from "./components/icons/QR.vue";
import { EBUTTON_OPTION } from "./enums/option";
import { FetchMethod } from "./services";

const config = useRuntimeConfig();

const url = ref<string>("");

const toolOption = ref<string>(EBUTTON_OPTION.SHORTEN);

const shortedUrl = ref<string[]>([]);

const onClearInput = () => {
  url.value = "";
  toolOption.value = EBUTTON_OPTION.SHORTEN;
};

onMounted(() => {
  getAllShorted();
});

const getAllShorted = async () => {
  const ipData = await FetchMethod(config.public.ipSource);
  const allData = await FetchMethod(
    `${config.public.apiBase}/all?ip=${ipData.ip}`,
    {
      method: "POST",
    }
  );

  shortedUrl.value = allData;
};

const onShorten = async () => {
  if (!url.value) return;
  toolOption.value = EBUTTON_OPTION.SHORTEN;

  const shortenData = await FetchMethod(`${config.public.apiBase}`, {
    method: "POST",
    body: JSON.stringify({ url: url.value }),
  });

  shortedUrl.value.push(shortenData.short);
  onClearInput();
};

const onClickOptionBtn = (option: EBUTTON_OPTION) => {
  toolOption.value = option;
};
</script>

<template>
  <main>
    <Header />

    <section>
      <div class="form">
        <input
          type="search"
          name="url"
          id="url"
          placeholder="Enter URL"
          v-model="url"
          required
        />
        <ul>
          <li>
            <button :disabled="!url" class="short_btn" @click="onShorten">
              <Link />
              <span>Shorten</span>
            </button>
          </li>

          <li>
            <button
              :disabled="!url"
              class="qr_btn"
              @click="onClickOptionBtn(EBUTTON_OPTION.QRCODE)"
            >
              <QR />
              <span>QR Code</span>
            </button>
          </li>

          <li>
            <button
              :disabled="!url"
              class="bar_btn"
              @click="onClickOptionBtn(EBUTTON_OPTION.BARCODE)"
            >
              <Bar />
              <span>Bar Code</span>
            </button>
          </li>
        </ul>
      </div>

      <ShortedList :shortedUrl="shortedUrl" />

      <QRCode :text="url" v-if="toolOption === EBUTTON_OPTION.QRCODE" />

      <BarCode :text="url" v-if="toolOption === EBUTTON_OPTION.BARCODE" />
    </section>

    <Footer />
  </main>
</template>

<style lang="scss">
main {
  font-family: "Open Sans", sans-serif;
  font-size: 14px;
  color: #333;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  min-height: 100vh;
  overflow: hidden;
  background-image: url("/landing.png");
  background-repeat: no-repeat;
  background-size: cover;
  background-position: 100% 100%;

  section {
    border-radius: 0.5rem;
    @include flex-center;
    flex-direction: column;

    .form {
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: space-between;
      margin: 1rem;
      padding: 2rem 1rem;
      background: rgba(247, 247, 247, 0.3);
      border-radius: 0.5rem;
      width: 100%;

      input[type="search"] {
        font-size: 0.9rem;
        border: none;
        box-shadow: 0px 5px 10px 0px rgba(0, 0, 0, 0.1);
        padding: 0.6rem 1rem;
        width: 100%;
      }

      ul {
        @include flex-center;
        width: 100%;
        margin-top: 1rem;
        list-style: none;

        li {
          .bar_btn {
            @include button($blue);
          }

          .qr_btn {
            @include button($purple);
          }

          .short_btn {
            @include button($red);
          }
        }
      }
    }
  }
}

@media (min-width: 400px) {
  section {
    width: 320px;
  }
}

@media (min-width: 500px) {
  section {
    width: 420px;
  }
}

@media (min-width: 640px) {
  section {
    width: 500px;
  }
}

@media (min-width: 768px) {
  section {
    width: 640px;
  }
}

@media (min-width: 850px) {
  section {
    width: 768px;
  }
}

@media (min-width: 1280px) {
  section {
    width: 850px;
  }
}
</style>

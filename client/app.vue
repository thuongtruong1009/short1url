<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { EBUTTON_OPTION } from './enums/option';
import QR from './components/icons/QR.vue';
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';
import Clear from './components/icons/Clear.vue';
import Link from './components/icons/Link.vue';
import AnalyzeIcon from './components/icons/Analyze.vue';
import Bar from './components/icons/Bar.vue';
import Download from './components/icons/Download.vue';
import Analyze from './components/Analyze.vue';
import BarCode from './components/BarCode.vue';
import QRCode from './components/QRCode.vue';
import ShortedList from './components/ShortedList.vue';
import { FetchMethod } from './services/fetch';

const config = useRuntimeConfig();

const url = ref<string>('');

const toolOption = ref<string>(null);

const shortedUrl = ref<string[]>([]);

onMounted(() => {
    getAllShorted();
})

const getAllShorted = async () => {
    const ipData = await FetchMethod('https://api.ipify.org?format=json')
    const allData = await FetchMethod(`${config.public.apiBase}/all?ip=${ipData.ip}`, {
        method: 'POST',
    });

    shortedUrl.value = allData;
}

const onClearInput = () => {
    url.value = '';
    toolOption.value = null;
}

const onShorten = async () => {
    if (!url.value) return;
    toolOption.value = EBUTTON_OPTION.SHORTEN

    const shortenData = await FetchMethod(`${config.public.apiBase}/api`, {
        method: 'POST',
        body: JSON.stringify({ url: url.value })
    })

    shortedUrl.value.push(shortenData.short);
    onClearInput();
}

const onClickOptionBtn = (option: EBUTTON_OPTION) => {
    toolOption.value = option
}

</script>

<template>
  <main>
    <Header />

    <section>
        <div class="form">
            <input type="text" name="url" id="url" placeholder="Enter URL" v-model="url" required>
            <ul>
                <li>
                    <button :disabled="!url" @click="onClearInput" class="clear_btn">
                        <Clear />
                        <span>Clear</span>
                    </button>
                </li>
                <li>
                    <button class="analyze_btn" @click="onClickOptionBtn(EBUTTON_OPTION.ANALYZE)">
                        <AnalyzeIcon />
                        <span>Analyze</span>
                    </button>
                </li>
                <li>
                    <button :disabled="!url" class="bar_btn" @click="onClickOptionBtn(EBUTTON_OPTION.BARCODE)">
                        <Bar />
                        <span>Bar Code</span>
                    </button>
                </li>
                <li>
                    <button :disabled="!url" class="qr_btn" @click="onClickOptionBtn(EBUTTON_OPTION.QRCODE)">
                        <QR />
                        <span>QR Code</span>
                    </button>
                </li>
                <li>
                    <button :disabled="!url" class="short_btn" @click="onShorten">
                        <Link />
                        <span>Shorten</span>
                    </button>
                </li>
            </ul>   
        </div>

        <ShortedList :shortedUrl="shortedUrl" />

        <Analyze v-if="toolOption === EBUTTON_OPTION.ANALYZE" />

        <QRCode :text="url" v-if="toolOption === EBUTTON_OPTION.QRCODE" />

        <BarCode :text="url" v-if="toolOption === EBUTTON_OPTION.BARCODE" />

    </section>

    <Footer />
  </main>
</template>

<style lang="scss">

main{
    font-family: 'Open Sans', sans-serif;
    font-size: 14px;
    color: #333;
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
    min-height: 100vh;
    overflow: hidden;
    background-image: url('/landing.png');
    background-repeat: no-repeat;
    background-size: cover;
    background-position: 100% 10%;

    section{
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
            background: rgb(247, 247, 247);
            background-image: url('/bg.png');
            background-repeat: no-repeat;
            background-size: cover;
            background-position: center;
            border-radius: 0.5rem;
            width: 100%;

            input[type="text"] {
                font-size: 0.9rem;
                border: none;
                box-shadow: 0px 5px 10px 0px rgba(0, 0, 0, 0.1); 
                padding: 0.6rem 1rem;
                width: 100%;
            }

            ul{
                @include flex-center;
                width: 100%;
                margin-top: 1rem;
                list-style: none;

                li{
                    .clear_btn {
                        @include button($gray);
                    }

                    .analyze_btn {
                        @include button($green);
                    }

                    .bar_btn {
                        @include button($orange)
                    }

                    .qr_btn {
                        @include button($purple)
                    }

                    .short_btn{
                        @include button($red);
                    }
                }
            }
        }
    }
}

@media (min-width: 320px) {
    section {
        width: 300px;
    }
}

@media (min-width: 576px) {
    section {
        width: 500px;
    }
}

@media (min-width: 768px) {
    section {
        width: 700px;
    }
}

@media (min-width: 992px) {
    section {
        width: 850px
    }
}

@media (min-width: 1024px) {
    section {
        width: 900px;
    }
}

@media (min-width: 1280px) {
    section {
        width: 950px;
    }
}

@media (min-width: 1400px) {
    section {
        width: 1000px;
    }
}
</style>
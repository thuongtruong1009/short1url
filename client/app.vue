<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { EBUTTON_OPTION } from './enums/option';
import QR from './components/icons/QR.vue';
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';
import Clear from './components/icons/Clear.vue';
import Link from './components/icons/Link.vue';
import Proxy from './components/icons/Proxy.vue';
import Analyze from './components/icons/Analyze.vue';
import Bar from './components/icons/Bar.vue';
import Download from './components/icons/Download.vue';
import Check from './components/icons/Check.vue';
import Copy from './components/icons/Copy.vue';
import SpeedTest from './components/SpeedTest.vue';
import BarCode from './components/BarCode.vue';
import QRCode from './components/QRCode.vue';

const config = useRuntimeConfig();

const url = ref<string>('');

const toolOption = ref<string>(null);

const shortedUrl = ref<string[]>([]);

onMounted(() => {
    getAllShorted();
})

const getAllShorted = async () => {
    const ipRes = await fetch('https://api.ipify.org?format=json');
    const ipData = await ipRes.json();

    const allRes = await fetch(`${config.apiBase}/all?ip=${ipData.ip}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
    });
    const allData = await allRes.json();

    shortedUrl.value = allData;
}

const onClearInput = () => {
    if (!url.value) return;
    toolOption.value = null;
    url.value = '';
    toolOption.value = null;
}

const onShorten = async (e: Event) => {
    if (!url.value) return;
    toolOption.value = EBUTTON_OPTION.SHORTEN
    e.preventDefault();

    const shortenResponse = await fetch(`${config.apiBase}/api`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url: url.value })
    });

    const shortenData = await shortenResponse.json();

    shortedUrl.value.unshift(shortenData.short);
    onClearInput();
}

const onBarCode = (e: Event) => {
    e.preventDefault();
    toolOption.value = EBUTTON_OPTION.BARCODE
}

const onClickQr = (e: Event) => {
    e.preventDefault();
    toolOption.value = EBUTTON_OPTION.QRCODE
}

const onAnalyze = (e: Event) => {
    e.preventDefault();
    toolOption.value = EBUTTON_OPTION.ANALYZE
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
                    <div :disabled="!url" @click="onClearInput" class="clear_btn">
                        <Clear />
                        <span>Clear</span>
                    </div>
                </li>
                <li>
                    <div :disabled="!url" class="analyze_btn" @click="onAnalyze">
                        <Analyze />
                        <span>Analyze</span>
                    </div>
                </li>
                <!-- <li>
                    <div :disabled="!url" class="proxy_btn">
                        <Proxy />
                        <span>Proxy</span>
                    </div>
                </li> -->
                <li>
                    <div :disabled="!url" class="bar_btn" @click="onBarCode">
                        <Bar />
                        <span>Bar Code</span>
                    </div>
                </li>
                <li>
                    <div class="qr_btn" @click="onClickQr">
                        <QR />
                        <span>QR Code</span>
                    </div>
                </li>
                <li>
                    <div :disabled="!url" class="short_btn" @click="onShorten">
                        <Link />
                        <span>Shorten</span>
                    </div>
                </li>
            </ul>   
        </div>

        <!-- <ul class="shorten_response" v-if="(toolOption === 'bar') && (shortedUrl?.length > 0)">
            <li v-for="(url, i) in shortedUrl" :key="i">
                <h3>Your shorted URL: </h3>
                <a :href="url" target="_blank">{{url}}</a>
                <button @click="onCopy(url, i)" class="copy_btn">
                    <Copy v-if="shortedUrlIndex !== i" />
                    <span v-if="shortedUrlIndex !== i">Copy</span>
                    <Check v-else />
                </button>
            </li>
        </ul> -->

        <ShortedList :shortedUrl="shortedUrl" v-if="toolOption === EBUTTON_OPTION.SHORTEN" />

        <SpeedTest v-if="toolOption === EBUTTON_OPTION.ANALYZE" />

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
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;

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
                display: flex;
                justify-content: center;
                align-items: center;
                width: 100%;
                margin: 1rem 0;
                list-style: none;

                li{
                    .clear_btn {
                        @include button($gray);
                    }

                    .analyze_btn {
                        @include button($green);
                    }

                    .proxy_btn {
                        @include button($blue);
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
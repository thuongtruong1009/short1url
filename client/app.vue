<script setup lang="ts">
import { ref, onMounted} from 'vue';
import QR from './components/QR.vue';

const url = ref<string>('');

const shortedUrl = ref<string[]>([]);

const clearInput = () => url.value = '';

const onSubmit = async (e: Event) => {
    if (!url.value) return;
    e.preventDefault();
    const response = await fetch('http://localhost:3000/api', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url: url.value })
    });

    const data = await response.json();

    shortedUrl.value.push(data.short);
    clearInput();
}

const generateQrCode = (qrContent) => {
      return new QRCode("qr-code", {
        text: qrContent,
        width: 256,
        height: 256,
        colorDark: "#000000",
        colorLight: "#ffffff",
        correctLevel: QRCode.CorrectLevel.H,
      });
    }

const qrCodeRef = ref(null)
let qrCode = null;


const onCreateQRCode = (e: Event) => {
    e.preventDefault();
    let qrContent = qrCodeRef.value;
    if (qrCode == null) {
        qrCode.clear();
    }
   qrCode = generateQrCode(url.value);
}

// onMounted(() => {
//     new QRCode("qrcode", url.value);
// });

</script>

<template>
  <main>
    <header>
        <img src="/favicon.svg" alt="logo">
        <h1>Short1url</h1>
    </header>

    <blockquote>Ship your link in an easier way</blockquote>

    <section>
        <form @submit.prevent="onSubmit">
            <input type="text" name="url" id="url" placeholder="Enter URL" v-model="url">
            <span @click="clearInput">Clear</span>
            <QR class="qr_btn" @click="onCreateQRCode" />
            <button type="submit" @click="onSubmit" :disabled="url.value">Shorten</button>
        </form>


        <ul class="response">
            <li v-for="(url, i) in shortedUrl" :key="i">
                <h3>Your shorted URL: </h3>
                <a :href="shortedUrl">{{shortedUrl}}</a>
                <button>Copy</button>
            </li>
        </ul>

        <div class="qrcode">
            <div id="qrcode" ref="qrCodeRef" />
        </div>
    </section>

    <footer>
        <p>Copyright <a href="https://github.com/thuongtruong1009">@thuongtruong1009</a>, 2023 - PRESENT</p>
    </footer>
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

    header {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        font-size: 2rem;
        margin-top: 3rem;
        letter-spacing: 0.25rem;
        font-family: 'Lobster Two', cursive;

        img {
            width: 4rem;
            height: 4rem;
            object-fit: cover;
            object-position: center;
            margin-right: 0.5rem;
        }

        h1 {
            color: $yellow;
            text-shadow: 1px 3px 10px rgba(255, 205, 2, 0.4);
        }
    }

    blockquote {
        color: $purple;
        font-style: italic;
        margin: 1rem 0 2rem;
        font-size: 1.2rem;
        font-family: 'Dancing Script', cursive;
    }

    section{
        border-radius: 0.5rem;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;

        form {
            display: flex;
            flex-direction: row;
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
                width: 85%;
            }

            span {
                margin: 0.5rem;
                font-size: 0.9rem;
                cursor: pointer;
                display: block;
            }

            .qr_btn {
                padding: 0.25rem;
                width: 2rem;
                height: 2rem;
                background: white;
                border-radius: 0.5rem;
                cursor: pointer;
            }

            button{
                @include button($red);
            }
        }

        .response {
            width: 100%;
            margin: 2rem 0;

            li {
                display: flex;
                justify-content: space-between;
                align-items: center;
                margin: 1rem 0;
                background: $light-blue;
                padding: 0.5rem;
                border-radius: 0.5rem;
                width: 100%;
                a {
                    font-weight: bold;
                    @include link($purple);
                }

                button {
                    @include button($green);
                    font-size: 1rem;
                }
            }
        }

        .qrcode {
            background: white;
            border-radius: 0.5rem;
            padding: 0.5rem;
            transform: scale(0.5);
        }
    }

    footer{
        position: absolute;
        bottom: 0;
        width: 100%;
        color: white;
        text-align: center;
        padding: 1rem;
        font-size: 0.8rem;
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

@media (min-width: 1200px) {
    section {
        width: 1000px;
    }
}
</style>
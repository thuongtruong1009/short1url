<script setup>
import { ref, onMounted } from 'vue';

const startTime = ref(null);
const endTime = ref(null);
const duration = ref(0);
const imageSize = ref('');
const imageLink = 'https://source.unsplash.com/random?topics=nature';
const speedInBps = ref('');
const speedInKbps = ref('');
const speedInMbps = ref('');

const calculateSpeed = () => {
    const timeDuration = (endTime.value - startTime.value) / 1000;
    duration.value = timeDuration;
    const loadedBits = imageSize.value * 8;
    speedInBps.value = (loadedBits / timeDuration).toFixed(2);
    speedInKbps.value = (speedInBps.value / 1024).toFixed(2);
    speedInMbps.value = (speedInKbps.value / 1024).toFixed(2);
};

const aa = async () => {
    startTime.value = new Date().getTime();
    const image = new Image();
    image.onload = async () => {
    endTime.value = new Date().getTime();
    const response = await fetch(imageLink);
    const contentLength = response.headers.get('content-length');
    imageSize.value = parseInt(contentLength);
    calculateSpeed();
    };
    image.src = imageLink;
};

const ipData = ref(null);

onMounted(async()=> {
  aa()
  const {data, pending, error} = await useFetch('http://localhost:3001/api/ip')
  ipData.value = data
});

// data
// http://ip-api.com/json
// https://jsonip.com/
// https://ipinfo.io/
// ui
// https://speedsmart.net/

</script>

<template>
  <div>
    <p>Speed In Bits: {{ speedInBps }} Kbps</p>
    <p>Speed In Kbs: {{ speedInKbps }} Mbps</p>
    <p>Speed In Mbs: {{ speedInMbps }} Mbps</p>
    <p>Duration: {{ duration }}s</p>

    <p>IP: {{ ipData }}</p>

  </div>
</template>
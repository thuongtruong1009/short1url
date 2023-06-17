<script setup>
import { ref, shallowRef, onMounted, computed } from 'vue';
import Status from './icons/ip/Status.vue';
import Country from './icons/ip/Country.vue';
import Speed from './icons/ip/Speed.vue';
import Region from './icons/ip/Region.vue';
import City from './icons/ip/City.vue';
import Zip from './icons/ip/Zip.vue';
import Lat from './icons/ip/Lat.vue';
import Lon from './icons/ip/Lon.vue';
import Timezone from './icons/ip/Timezone.vue';
import Isp from './icons/ip/Isp.vue';
import Org from './icons/ip/Org.vue';
import As from './icons/ip/As.vue';
import Query from './icons/ip/Query.vue';
import Ping from './icons/ip/Ping.vue';
import { FetchMethod } from '../services/fetch';

const ipItems = shallowRef([]);
const config = useRuntimeConfig();

const colors = ['#FF6633', '#FFB399', '#FF33FF', '#FFFF99', '#00B3E6',
  '#E6B333', '#3366E6', '#999966', '#99FF99', '#B34D4D',
  '#80B300', '#809900', '#E6B3B3', '#6680B3', '#66991A',
  '#FF99E6', '#CCFF1A', '#FF1A66', '#E6331A', '#33FFCC',
  '#66994D', '#B366CC', '#4D8000', '#B33300', '#CC80CC',
  '#991AFF', '#E666FF', '#4DB3FF', '#1AB399',
  '#E666B3', '#33991A', '#CC9999', '#B3B31A', '#00E680','#E6FF80', '#1AFF33', '#999933',
  '#FF3380', '#CCCC00', '#66E64D', '#4D80CC', '#9900B3',
  '#E64D66', '#4DB380', '#FF4D4D', '#99E6E6', '#6666FF'];

const randomColor = () => {
  return colors[Math.floor(Math.random() * colors.length)];
}

const hexToRgb = (hex) => {
  const shorthandRegex = /^#?([a-f\d])([a-f\d])([a-f\d])$/i;
  hex = hex.replace(shorthandRegex, (m, r, g, b) => {
    return r + r + g + g + b + b;
  });
  const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
  return result
    ? `${parseInt(result[1], 16)}, ${parseInt(result[2], 16)}, ${parseInt(result[3], 16)}`
    : null;
}

const randomBackgroundColor = () => {
  return `rgba(${hexToRgb(randomColor())}, 0.1)`;
}
const getIPInfo = async () => {
  const data = await FetchMethod(`${config.public.appBase}/api/ip`)
  // const {data, pending, error} = await useFetch(`${process.env.NUXT_PUBLIC_APP_BASE}/api/ip`)
  // data
  // http://ip-api.com/json
  // https://jsonip.com/
  // https://ipinfo.io/
  // ui
  // https://speedsmart.net/
  console.log('here: ', data)
  ipItems.value = [
    { name: 'status', text: data.status, icon: Status },
    { name: 'country', text: data.country, icon: Country },
    { name: 'region', text: data.regionName, icon: Region },
    { name: 'city', text: data.city, icon: City },
    { name: 'zip', text: data.zip, icon: Zip },
    { name: 'lat', text: data.lat, icon: Lat },
    { name: 'lon', text: data.lon, icon: Lon },
    { name: 'timezone', text: data.timezone, icon: Timezone },
    { name: 'isp', text: data.isp, icon: Isp },
    { name: 'org', text: data.org, icon: Org },
    { name: 'as', text: data.as, icon: As },
    { name: 'query', text: data.query, icon: Query }
  ]
  return data
}

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
    ipItems.value.unshift({ 
      name: 'duration',
      text: duration.value + ' s',
      icon: Ping
    })
    const loadedBits = imageSize.value * 8;
    speedInBps.value = (loadedBits / timeDuration).toFixed(2);
    speedInKbps.value = (speedInBps.value / 1024).toFixed(2);
    speedInMbps.value = (speedInKbps.value / 1024).toFixed(2);
    ipItems.value.unshift({ 
      name: 'speed',
      text: speedInKbps.value + ' Kbps',
      icon: Speed
    })
};

const loadSpeed = async () => {
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

onMounted(()=> {
  loadSpeed()
  getIPInfo()
});

</script>

<template>
  <div class="analyze_container">
    <h3>Your analyze result testing</h3>
    <ul>
      <li v-for="item in ipItems.filter((i) => i.text !=='')" :key="item" :style="{ backgroundColor: randomBackgroundColor() }">
        <div class="cell">
          <div class="info_left">
            <component :is="item.icon" class="info_left-icon" :style="{color: randomColor()}" />
          </div>

          <div class="info_right">
            <h4>{{item.name}}</h4>
            <p>{{item.text}}</p>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<style lang="scss" scoped>
.analyze_container {
  @include flex-center;
  flex-direction: column;
  background: #fff;
  border-radius: 1rem;
  padding: 1rem;

  h3 {
    margin-bottom: 1rem;
  }

  ul {
    list-style: none;
    display: flex;
    justify-content: center;
    flex-wrap: wrap;

    li {
      margin: 0.5rem;
      padding: 0.5rem;
      width: 12rem;
      min-height: 5rem;
      border-radius: 0.5rem;
      display: flex;
      justify-content: flex-start;
    }

    .cell {
      display: flex;
      width: 100%;

      .info_left > .info_left-icon {
        margin-right: 0.5rem;
        width: 1.75rem;
        height: 1.75rem;
      }

      .info_right {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
        word-break: break-word;
        
        h4 {
          font-size: 1.1rem;
          margin-bottom: 0.5rem;
          text-align: center;
          text-transform: capitalize;
        }

        p{
          font-size: 0.8rem;
        }
      }
    }
  }
}
</style>
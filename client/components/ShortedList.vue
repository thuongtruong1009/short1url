<script setup lang="ts">
import { ref } from 'vue';

defineProps<{
    shortedUrl: string[];
}>();

const shortedUrlIndex = ref<number>(null)
const onCopy = (text: string, index: number) => {
    navigator.clipboard.writeText(text)
    shortedUrlIndex.value = index
    setTimeout(() => {
        shortedUrlIndex.value = null
    }, 1000)
}
</script>

<template>
    <ul class="shorten_response" v-if="shortedUrl?.length > 0">
        <li v-for="(url, i) in shortedUrl" :key="i">
            <h3>Your shorted URL: </h3>
            <a :href="url" target="_blank">{{url}}</a>
            <button @click="onCopy(url, i)" class="copy_btn">
                <Copy v-if="shortedUrlIndex !== i" />
                <span v-if="shortedUrlIndex !== i">Copy</span>
                <Check v-else />
            </button>
        </li>
    </ul>
</template>

<style lang="scss" scoped>
    .shorten_response {
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
</style>
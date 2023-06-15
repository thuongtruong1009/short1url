<script setup lang="ts">
import ShortedItem from './ShortedItem.vue';

defineProps<{
    shortedUrl: string[];
}>();

</script>

<template>
    <ul>
        <ShortedItem :url="url" v-for="(url, i) in shortedUrl" :key="i" />
    </ul>
</template>

<style lang="scss" scoped>
ul {
    width: 100%;
    margin: 2rem 0;

    li{
        display: flex;
        width: 100%;
    }
}
</style>
<script setup lang="ts">
import { onMounted } from 'vue';
import JsBarcode from 'jsbarcode';

defineProps<{
    text: string
}>()

onMounted(()=> {
    // https://fengyuanchen.github.io/vue-barcode/
    JsBarcode("#barcode", "1234", {
        lineColor: "#0aa",
        height:40,
        displayValue: false
    });
});

</script>

<template>
    <div>
        <svg id="barcode"></svg>
        <img id="barcode"/>
    </div>
</template>

<style scoped>
svg{
    border-radius: 0.5rem;
}
</style>
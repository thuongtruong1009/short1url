<script setup lang="ts">
import JsBarcode from "jsbarcode";
import { onMounted } from "vue";

defineProps<{
  text: string;
}>();

onMounted(() => {
  JsBarcode("#barcode", "1234", {
    lineColor: "#0aa",
    height: 40,
    displayValue: false,
  });
});
</script>

<template>
  <div>
    <!-- <svg id="barcode"></svg> -->
    <img id="barcode" />
  </div>
</template>

<style scoped>
svg,
img {
  border-radius: 0.5rem;
  width: 22rem;
  height: 8rem;
}
</style>

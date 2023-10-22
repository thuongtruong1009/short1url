<script setup lang="ts">
import { ref } from "vue";
import Check from "./icons/Check.vue";
import Copy from "./icons/Copy.vue";

defineProps<{
  url: string;
}>();

const isCopied = ref<boolean>(false);

const onCopy = (text: string) => {
  navigator.clipboard.writeText(text);
  isCopied.value = true;
  setTimeout(() => {
    isCopied.value = false;
  }, 1000);
};
</script>

<template>
  <li>
    <h3>Your shorted URL:</h3>
    <a :href="url" target="_blank">{{ url }}</a>
    <button @click="onCopy(url)">
      <Copy v-if="!isCopied" />
      <span v-if="!isCopied">Copy</span>
      <Check v-else />
    </button>
  </li>
</template>

<style lang="scss" scoped>
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
</style>

<script setup lang="ts">
import { computed, ref } from "vue";
import Check from "./icons/Check.vue";
import Copy from "./icons/Copy.vue";

const props = defineProps<{
  url: string;
}>();

// const config = useRuntimeConfig();
// const currentUrl = ref<string>(
//   computed(() => String(config.public.apiBase) + props.url) as unknown as string
// );
const currentUrl = ref<string>(computed(() => props.url) as unknown as string);

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
    <h4>Your shorted URL:</h4>
    <a :href="props.url" target="_blank">{{ props.url }}</a>
    <button @click="onCopy(props.url)">
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

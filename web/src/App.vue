<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
import Aslide from './components/Aslide/index.vue';
</script>

<template>
  <div>
    <RouterView />
    <Aslide></Aslide>
  </div>
</template>

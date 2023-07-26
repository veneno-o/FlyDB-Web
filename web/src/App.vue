<script setup lang="ts">
import { RouterView } from 'vue-router'
import {useCounterStore} from "./stores/index"
import Sidebar from "@/components/Sidebar/index.vue"
const locale = useCounterStore()

</script>

<template>
  <el-config-provider :locale="locale">
    <RouterView />
    <Sidebar />
  </el-config-provider>
</template>

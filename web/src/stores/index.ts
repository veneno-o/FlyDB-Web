// 配置国际化,只能支持对组件内的语言国际化
import en from 'element-plus/dist/locale/en.mjs'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'

import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('root', () => {
  const language:any = ref('zh-cn')
  const locale:any = computed(() => (language.value === 'zh-cn' ? zhCn : en))
  const toggle = () => {
    language.value = language.value === 'zh-cn' ? 'en' : 'zh-cn'
  }

  return { toggle, locale }
})

<template>
  <div>
    <n-menu
        :options="menuOptions"
    />
  </div>
</template>

<script setup>
import {computed} from "vue";
import {ToolBox} from "@vicons/carbon";

const menuOptions = computed(() => {
  return [
    {
      label: "小工具",
      key: "tools",
      icon: ToolBox,
    },
    {
      label: "Server",
      key: "server",
      icon: ToolBox,
    },
  ]
})
</script>

<style scoped>

</style>
<template>
  <n-card>
    <n-form
        ref="form"
    >
      <n-form-item label="游戏名称" path="name">
        <n-input v-model:value="form.name" placeholder="输入游戏名称" />
      </n-form-item>
      <n-button attr-type="button" @click="pullData">
        搜索
      </n-button>
    </n-form>
  </n-card>
  <n-card>
    <n-data-table
        :bordered="false"
        :single-line="false"
        :columns="columns"
        :data="data"
    />
  </n-card>
</template>

<script setup>
import {h, ref} from "vue";
import {GetGameSearchSuggest} from '../../../wailsjs/go/service/SteamService'
const form = ref({
  name: null,
})
const handleGameSearch = () => {
  console.log("form ", form.value)
}
const columns = [
  {
    title: '游戏名称',
    render(row) {
      return h("a", {href: row.url, target: "_blank"}, row.name)
    },
  },
  {
    title: '封面',
    render(row) {
      return h("img", {src: row.imageUrl})
    }
  },
  {
    title: '价格',
    key: 'price'
  }
]

const data = ref([])

const pullData = () => {
  GetGameSearchSuggest(form.value.name).then(result => {
    if (result.success) {
      if(result.data != null) {
        data.value = result.data
      } else {
        $message.warning("未找到相关游戏")
      }
    }
    console.log(result)
  })
}
</script>

<style scoped>

</style>
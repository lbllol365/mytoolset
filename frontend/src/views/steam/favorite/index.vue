<template>
  <n-card>
    <n-button attr-type="button" @click="pullData">
      拉取价格及创意工坊信息
    </n-button>
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

const _this = this
const pullData = () => {
  $message.success("拉取成功")
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
  },
  {
    title: '创意工坊物品数量',
    key: 'workshop'
  },
  {
    title: '操作',
    render(row) {
      return h("div", [
        h("n-button", {
          props: {
            type: "primary", size: "small"
          },
          on: {
            click: () => {
              // TODO 调用后端接口取消收藏
            }
          }
        }),
      ])
    }
  }
]

const data = ref([])

</script>

<style scoped>

</style>
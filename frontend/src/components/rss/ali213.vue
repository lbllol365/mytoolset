<template>
  <n-card>
    <n-button @click="pullData">
      拉取最新资源
    </n-button>
    <n-button @click="clearStore">
      清空store
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
import {PullData} from "../../../wailsjs/go/service/RssService";
import {h, ref} from "vue";
import {NIcon} from "naive-ui";
import {includes} from 'lodash'
import useRssStore from "../../store/rss";
import {
  CheckmarkOutline as CheckIcon,
  CloseOutline as CloseIcon
} from "@vicons/ionicons5";

const data = ref([]);
const rssStore = useRssStore();
const pullData = () => {
  let storedRssData = rssStore.getRssData
  if(storedRssData !== null) {
    data.value = storedRssData
  } else {
    PullData().then(result => {
      console.log(result);
      if(result.success) {
        $message.success("拉取成功");
        data.value = result.data
        rssStore.updateRssData(result.data)
      } else {
        $message.error("拉取失败: " + result.msg);
      }
    })
  }
}
const columns = [
  {
    title: "名称",
    key: "title"
  },
  {
    title: "是否正版分流",
    key: "title",
    render(row) {
      if(includes(row.title, "正版分流")) {
        return h(NIcon, {color: '#00FF00', size: "30"}, h(CheckIcon))
      } else {
        return h(NIcon, {color: '#FF0000', size: "30"}, h(CloseIcon))
      }
    },
    resizable: true,
    maxWidth: 30
  },
  {
    title: "链接",
    key: "link",
    render(row) {
      return h("a", {href: row.link, target: "_blank"}, "前往")
    },
    minWidth: 60
  },
  {
    title: "发布时间",
    key: "published"
  }
];

const clearStore = () => {
  rssStore.clearState()
}
</script>

<style scoped>

</style>
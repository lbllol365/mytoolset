<template>
  <n-card>
    <n-tag v-if="dockerClientStatus" :bordered="false" type="success">
      已连接
    </n-tag>
    <n-tag v-if="!dockerClientStatus" :bordered="false" type="error">
      未连接
    </n-tag>
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
import {onBeforeMount, ref, onBeforeUnmount} from "vue";
import {ListImage, PingClient} from "../../../wailsjs/go/service/DockerService";

const columns = [
    {
      title: '名字',
      key: 'name'
    },
    {
      title: '标签',
      key: 'tag'
    },
    {
      title: '大小/MB',
      key: 'size'
    },
    {
      title: '拉取时间',
      key: 'created'
    }
  ]

  const pingClient = () => {
    PingClient().then(result => {
      dockerClientStatus.value = result.success;
    })
  }
  const data = ref([])
  const dockerClientStatus = ref(false)
  const timer = ref(null)
  onBeforeMount(() => {
    PingClient().then(result => {
      dockerClientStatus.value = result.success;
      if(dockerClientStatus) {
        timer.value = setInterval(() => {
          pingClient()
        }, 3000)
        ListImage().then(result => {
          if(result.success) {
            data.value = result.data.reduce((pre, curr) => {
              let nameTag = curr.RepoTags[0].split(':');
              pre.push({
                name: nameTag[0],
                tag: nameTag[1],
                size: (curr.Size / 1000000).toFixed(2),
                created: new Date(curr.Created * 1000).toLocaleString()
              })
              return pre;
            }, []);
          }
        })
      }
    })

  })

  onBeforeUnmount(() => {
    clearInterval(timer)
  })
</script>
<style scoped>

</style>
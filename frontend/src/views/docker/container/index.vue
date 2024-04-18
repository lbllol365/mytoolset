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
import {PingClient} from "../../../../wailsjs/go/service/DockerService.js";


const columns = [
    {
        title: '名称',
        key: 'name'
    },
    {
        title: '来源镜像',
        key: 'image'
    },
    {
        title: '启动命令',
        key: 'command'
    },
    {
        title: '开放端口',
        key: 'ports'
    },
    {
      title: '创建时间',
      key: 'created'
    },
    {
      title: '标签',
      key: 'labels'
    },
    {
        title: '状态',
        key: 'status' // TODO 根据实际值判断，实际展示红绿状态标签 (watch 状态？)
    },

]

const data = ref([])
const timer = ref(null)
const dockerClientStatus = ref(false)
const pingClient = () => {
    PingClient().then(result => {
        dockerClientStatus.value = result.success;
    })
}


</script>

<style scoped>

</style>
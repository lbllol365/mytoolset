<template>
    <div>
        <n-space>
            <n-button @click="handleWarn">
                警告
            </n-button>
            <n-button @click="handleSuccess">
                成功
            </n-button>
            <n-button @click="handleError">
                错误
            </n-button>
        </n-space>
    </div>
</template>

<script setup>

const handleWarn = () => {
    $dialog.warning('警告'
        , () => {$message.success("确定");}
        , () => {$message.error("不确定");
    })
}

const handleSuccess = () => {
    $dialog.success('成功', () => {
        $message.success("耶！");
    });
}

const handleError = () => {
    $dialog.error('错误',() => {
        $message.success("我就知道");
    });
}

</script>

<style scoped>

</style>
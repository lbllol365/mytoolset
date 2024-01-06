<template>
    <n-layout-header>
        <n-form
            ref="formRef"
            :model="model"
            label-placement="top"
            label-width="auto"
            require-mark-placement="right-hanging"
        >
            <n-grid :cols="24" :x-gap="24">
                <n-form-item-gi :span="12" label="name">
                    <n-input v-model:value="model.name" placeholder="输入姓名" />
                </n-form-item-gi>
                <n-form-item-gi :span="12" label="age">
                    <n-input v-model:value="model.age" placeholder="输入年龄" />
                </n-form-item-gi>
                <n-form-item-gi :span="12" label="Datetimeadsadasdasdasd">
                    <n-date-picker v-model:value="model.time" type="datetime" />
                </n-form-item-gi>
                <n-form-item-gi :span="12" label="sex">
                    <n-checkbox-group v-model:value="model.sex">
                        <n-space>
                            <n-checkbox value="man">
                                man
                            </n-checkbox>
                            <n-checkbox value="female">
                                female
                            </n-checkbox>
                        </n-space>
                    </n-checkbox-group>
                </n-form-item-gi>
            </n-grid>
        </n-form>
    </n-layout-header>
    <n-layout-content>
        <n-data-table
            :columns="cols"
            :data="data"
            :bordered="false"
        />
    </n-layout-content>
</template>

<script setup>
import {ref} from "vue";

const  model = ref({
    name: null,
    age: null,
    sex: null,
    time: null,
});
const createColumns = () => {
    return [
        {
            title: "No",
            key: "no"
        },
        {
            title: "Title",
            key: "title"
        },
        {
            title: "Length",
            key: "length"
        }
    ];
};

const data = [
    { no: 3, title: "Wonderwall", length: "4:18" },
    { no: 4, title: "Don't Look Back in Anger", length: "4:48" },
    { no: 12, title: "Champagne Supernova", length: "7:27" }
];

const cols = createColumns();
</script>

<style scoped>

</style>
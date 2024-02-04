<script setup>
import {h, ref} from "vue";
import {NIcon} from "naive-ui";
import {RouterLink} from "vue-router";
import {
  BookOutline as BookIcon,
  LogoRss as RssIcon,
  GameControllerOutline as GameControllerIcon,
  Alarm as AlarmIcon,
  LogoDocker as DockerIcon
} from "@vicons/ionicons5";

function renderIcon(icon) {
  return () => h(NIcon, null, {default: () => h(icon)});
}

const collapsed = ref(true);
const activeKey = ref('tool1');

const menuOptions = [
  {
    label: () => h(
        RouterLink,
        {
          to: {name: "tool1"}
        },
        {default: () => "tool1"}
    ),
    key: "tool1",
    icon: renderIcon(GameControllerIcon)
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {name: "tool2"}
        },
        {default: () => "tool2"}
    ),
    key: "tool2",
    icon: renderIcon(BookIcon)
  },
  {
    label: 'RSS',
    key: "rss",
    icon: renderIcon(RssIcon),
    children: [
      {
        label: () => h(
            RouterLink,
            {
              to: {name: "ali213"}
            },
            {default: () => "游侠"}
        ),
        key: "ali213",
        icon: renderIcon(BookIcon)
      },
    ]
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {name: "steam"}
        },
        {default: () => "steam"}
    ),
    key: "steam",
    icon: renderIcon(AlarmIcon)
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {name: "docker"}
        },
        {default: () => "docker"}
    ),
    key: "docker",
    icon: renderIcon(DockerIcon)
  },
];


</script>

<template>
  <div>
    <n-space vertical size="large">
      <n-layout has-sider position="absolute">
        <n-layout-sider
            bordered
            collapse-mode="width"
            :collapsed-width="64"
            :width="240"
            :collapsed="collapsed"
            show-trigger
            @collapse="collapsed = true"
            @expand="collapsed = false"
        >
          <n-menu
              v-model:value="activeKey"
              :collapsed="collapsed"
              :collapsed-width="64"
              :collapsed-icon-size="22"
              :options="menuOptions"
          />
        </n-layout-sider>
        <n-layout>
          <n-layout-content>
            <router-view class="view main-content"></router-view>
          </n-layout-content>
        </n-layout>
      </n-layout>
    </n-space>
  </div>
</template>

<style>

</style>

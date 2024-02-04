import {computed, createApp} from 'vue'
import App from './App.vue'
import naive from 'naive-ui'
import { createRouter,createWebHashHistory } from 'vue-router'
import tool1 from "./components/tool/tool1.vue";
import tool2 from "./components/tool/tool2.vue";
import ali213 from "./components/rss/ali213.vue";
import steam from "./views/steam/index.vue";
import docker from "./views/docker/index.vue"
import { createDiscreteApi, darkTheme} from 'naive-ui'
import {createPinia} from "pinia";

const pinia = createPinia()

const routes = [
    { path: '/', component: App },
    { path: '/tool1', component: tool1, name: 'tool1' },
    { path: '/tool2', component: tool2, name: 'tool2' },
    {path: '/rss/ali213', component: ali213, name: 'ali213'},
    {path: '/steam/index', component: steam, name: 'steam'},
    {path: '/docker/index', component: docker, name: 'docker'}
]

const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHashHistory(),
    routes, // `routes: routes` 的缩写
})
async function setupApp() {
    const app = createApp(App)
    await setupDiscreteApi();
    app.use(naive);
    app.use(pinia);
    app.use(router);
    app.mount('#app')
}

function setupDialog(dialog) {
    return {
        show(option) {
            option.closable = option.closable === true
            option.autoFocus = option.autoFocus === true
            option.transformOrigin = 'center'
            return dialog.create(option)
        },
        warning: (content, onConfirm, onCancel) => {
            return dialog.warning({
                title: '警告',
                content: content,
                closable: false,
                autoFocus: false,
                transformOrigin: 'center',
                positiveText: '确认',
                negativeText: '取消',
                onPositiveClick: () => {
                    onConfirm && onConfirm()
                },
                onNegativeClick: () => {
                    onCancel && onCancel()
                }
            })
        },
        success: (content, onConfirm) => {
            return dialog.success({
                title: '成功',
                content: content,
                closable: false,
                autoFocus: false,
                transformOrigin: 'center',
                positiveText: '确认',
                negativeText: '取消',
                onPositiveClick: () => {
                    onConfirm && onConfirm()
                }
            })
        },
        error: (content, onConfirm) => {
            return dialog.error({
                title: '失败',
                content: content,
                closable: false,
                autoFocus: false,
                transformOrigin: 'center',
                positiveText: '确认',
                negativeText: '取消',
                onPositiveClick: () => {
                    onConfirm && onConfirm()
                }
            })
        }
    }
}

function setupMessage(message) {
    return {
        error: (content, option = null) => {
            return message.error(content, option)
        },
        info: (content, option = null) => {
            return message.info(content, option)
        },
        loading: (content, option = {}) => {
            option.duration = option.duration != null ? option.duration : 30000
            option.keepAliveOnHover = option.keepAliveOnHover !== undefined ? option.keepAliveOnHover : true
            return message.loading(content, option)
        },
        success: (content, option = null) => {
            return message.success(content, option)
        },
        warning: (content, option = null) => {
            return message.warning(content, option)
        },
    }
}

async function setupDiscreteApi() {
    const configProviderProps = computed(() => ({
        theme: darkTheme
    }))
    const { message, dialog, notification } = createDiscreteApi(['message', 'notification', 'dialog'], {
        configProviderProps,
        messageProviderProps: {
            placement: 'bottom-right',
            keepAliveOnHover: true,
        },
        notificationProviderProps: {
            max: 5,
            placement: 'bottom-right',
            keepAliveOnHover: true,
        },
    })
    window.$message = setupMessage(message)
    window.$dialog = setupDialog(dialog)
}

setupApp()
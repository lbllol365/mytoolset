import {defineStore} from "pinia";
import {PingClient} from "../../wailsjs/go/service/DockerService.js";


const useDockerStore = defineStore('docker', {
    state: () => ({
        clientActive: false,
        imageList: [],
        containerList: [],
    }),
    getters: {
        getDockerClientActive() {

        }
    },
    actions: {
        updateImageList(data) {
            this.imageList = data
        },
        clearImageList() {
            this.updateImageList([])
        },
        updateContainerList(data) {
            this.containerList = data
        },
        clearContainerList() {
            this.updateContainerList([])
        }
    }
})
import { defineStore } from 'pinia'

const useRssStore = defineStore('rss', {
    state: () => ({
        rssData: [],
        lastUpdate: null,
    }),
    getters: {
        getRssData() {
            let needRePullData = this.lastUpdate === null || (new Date().getTime() - this.lastUpdate.getTime()) > 1000 * 60 * 60
            if (needRePullData) {
                return null
            } else {
                return this.rssData
            }
        }
    },
    actions: {
        updateRssData(data) {
            this.rssData = data
            this.lastUpdate = new Date()
        },
        clearState() {
            this.rssData = [];
            this.lastUpdate = null;
        }
    },

})

export default useRssStore
import Vue from 'vue'
import Router from 'vue-router'
import Data from '@/components/Data'
import ResultV2 from '@/components/ResultV2'
Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            component: Data
        },
        {
            path: '/index',
            component: Data
        },
        {
            path: '/result',
            component: ResultV2
        }
    ]
})
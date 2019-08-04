import Vue from 'vue'
import Router from 'vue-router'
import DrawInfo from '@/components/DrawInfo'
import Home from '@/components/Home'
import Data from '@/components/Data'
import ResultV2 from '@/components/ResultV2'
Vue.use(Router)

export default new Router({
    routes: [{
            path: '/data',
            component: Data
        },
        {
            path: '/drawInfo',
            component: DrawInfo
        },
        {
            path: '/',
            component: Home
        },
        {
            path: '/index',
            component: Home
        },
        {
            path: '/result',
            component: ResultV2
        }
    ]
})
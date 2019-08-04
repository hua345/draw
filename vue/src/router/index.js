import Vue from 'vue'
import Router from 'vue-router'
import DrawInfo from '@/components/DrawInfo'
import HomeV2 from '@/components/HomeV2'
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
            component: HomeV2
        },
        {
            path: '/index',
            component: HomeV2
        },
        {
            path: '/result',
            component: ResultV2
        }
    ]
})
import Vue from 'vue'
import Router from 'vue-router'
import DrawInfo from '@/components/DrawInfo'
import Home from '@/components/Home'
import Data from '@/components/Data'
import Result from '@/components/Result'
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
            component: Result
        }
    ]
})
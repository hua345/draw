import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
// 该对象的状态(state)就是全局属性, 类似于组件的data，每个组件都可以访问和修改属性
// 可变的(mutations)类似于组件中的methods, mutations中的每个方法称作为 mutation handler，
// 用来修改state中的值，方法的参数可以有两个(state, payload) state表示全局的单例对象，
// payload(载荷)也就是参数，调用时可以传递参数，该参数是可选的。

var excelData = new Vuex.Store({
    state: {
        excelData: null,
        selectedTypeInfoList: null,
        selectedCompanyName: null,
        drawResult: null,
        isUpload: false
    },
    mutations: {
        SET_EXCEL_DATA(state, payload) {
            state.excelData = payload
        },
        SET_IsUpload(state, payload) {
            state.isUpload = payload
        },
        RESET_EXCEL_DATA(state) {
            state.excelData = null;
        },
        SET_SelectedTypeInfoList(state, payload) {
            state.selectedTypeInfoList = payload
        },
        SET_SelectedCompanyName(state, payload) {
            state.selectedCompanyName = payload
        },
        SET_DrawResult(state, payload) {
            state.drawResult = payload
        },
    },
})

export default excelData
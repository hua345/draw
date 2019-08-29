<template>
  <div id="app">
    <HomeTesseract />
  </div>
</template>

<script>
import HomeTesseract from './components/HomeTesseract.vue'

export default {
  name: 'app',
  components: {
    HomeTesseract
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

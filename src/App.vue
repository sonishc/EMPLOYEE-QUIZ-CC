<template>
  <div id="app">
    <IndexQuiz msg="Welcome to EMPLOYEE-QUIZ-CC App"/>
  </div>
</template>

<script>
import IndexQuiz from './components/IndexQuiz.vue'

export default {
  name: 'app',
  components: {
    IndexQuiz
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

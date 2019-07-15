<template>
  <div class="container">
    <h1 class="bg bg-info text-white text-center">{{ title }}</h1>
    <QuestionsList />
  </div>
</template>

<script>
import QuestionsList from './QuestionsList.vue';

export default {
  name: 'IndexQuiz',
  components: {
    QuestionsList
  },
  props: {
    title: String
  }
};
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

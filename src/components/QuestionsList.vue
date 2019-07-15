<template>
  <div>
    <ul>
      <li v-for="question in questions" v-bind:key="question.id">
        <b>Question:</b> {{ question.text }}
        <br />
        <b>Answer:</b> {{ question.answer }}
        <hr />
      </li>
    </ul>
    <div class="d-flex justify-content-center">
      <button
        class="btn btn btn-primary"
        v-show="!startQuiz"
        @click="fetchQuestions()"
      >
        Start Quize
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'QuestionsList',
  data() {
    return {
      questions: [],
      startQuiz: false
    };
  },
  methods: {
    fetchQuestions: function() {
      const questionsURI = this.serverURI;
      this.$http.get(questionsURI).then(result => {
        this.questions = result.data;
        this.startQuiz = !this.startQuiz;
      });
    }
  }
};
</script>

<style>
ul {
  list-style: none;
}
</style>

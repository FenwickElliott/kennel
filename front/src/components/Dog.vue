<template>
  <div class="Dogs">
    <h1>This is about {{ this.$route.params.collar }}</h1>
      <div v-if="(ready == true)">
        <p>{{dog}}</p>
        <button @click=kill>delete</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dogs',
  data() {
    return {
      ready: false,
      dog: {},
    }
  },
  mounted() {
    fetch('http://localhost:1066/dog/' + this.$route.params.collar)
    .then(res => res.json())
    .then(res => { this.dog = res; this.ready = true; console.log(this.dog) })
    .catch(e => console.log(e))
  },
  methods: {
    kill() {
      fetch('http://localhost:1066/dog/' + this.$route.params.collar, { method: 'DELETE' })
      .then(res => console.log(res))
      .then(this.$router.push({name: 'Dogs'}))
      .catch(e => console.log(e))
    }
  }
}
</script>
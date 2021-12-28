<template>
  <div class="DogsNew">
    <h1>Crate a dog!</h1>
    <!-- <form> -->
        collar: <input type="text" v-model="dog.collar" placeholder="how you bark?"> <br/>
        alive: <input type="checkbox" v-model="dog.alive" placeholder="still kicking?"> <br/>
        toes: <input type="number" v-model="dog.toes" placeholder="how many toeses?"> <br/>
        <button @click="submit()">done</button>
    <!-- </form> -->
  </div>
</template>

<script>
export default {
  name: 'DogsNew',
  data() {
    return {
      ready: false,
      dog: {},
    }
  },
  methods: {
    submit() {
      console.log('submitting:', JSON.stringify(this.dog))
      fetch('http://localhost:1066/dogs', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', },
        body: JSON.stringify(this.dog),
      })
      .then(res => console.log(res))
      .then(this.$router.push({name: 'Dog', params: { collar: this.dog.collar}}))
      .catch(e => console.log(e))
    }
  }
}
</script>
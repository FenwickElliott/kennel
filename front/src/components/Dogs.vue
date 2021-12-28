<template>
  <div class="Dogs">
    <h1>This is an about dogs</h1>
      <div v-if="(ready)">
      <table>
        <th>name</th>
        <th>alive</th>
        <th>toes</th>
        <tr v-for="dog in dogs" v-bind:key="dog.collar">
          <td>{{dog.collar}}</td>
          <td>{{dog.alive}}</td>
          <td>{{dog.toes}}</td>
          <td><router-link :to='{name: "Dog", params: {collar: dog.collar } }'>view / edit</router-link></td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dogs',
  data() {
    return {
      ready: false,
      dogs: [],
    }
  },
  mounted() {
    fetch("http://localhost:1066/dogs")
    .then(res => res.json())
    .then(res => { this.dogs = res; this.ready = true; console.log(this.dogs) })
    .catch(e => console.log(e))
  }
}
</script>
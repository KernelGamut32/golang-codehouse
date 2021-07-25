<template>
  <div id="app">
    <h1>To-Do List</h1>
    <form method="POST" @submit.prevent="sendItem()">
      <input
        type="text"
        size="50"
        v-model="todoitem"
        placeholder="Enter new item"
      />
      <input type="submit" value="Submit" />
    </form>
    <ul>
      <li v-for="item in todolist" v-bind:key="item">{{ item }}</li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  data() {
    return {
      todolist: [],
      todoitem: "",
    };
  },
  mounted: function () {
    this.getList();
  },
  methods: {
    getList() {
      axios.get("http://localhost:8080/api/lists").then((res) => {
        this.todolist = res.data.list;
      });
    },
    async sendItem() {
      const params = new URLSearchParams();
      params.append("item", this.todoitem);
      await axios.post("http://localhost:8080/api/lists", params);
      this.getList();
    },
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

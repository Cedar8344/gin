<template>
  <div>
    <div>MyApp</div>
    <button v-on:click="showMessage">I want help</button>
    <div>{{message}}</div>
    <button v-on:click="addPoint">I want die</button>
    <div>{{message2}}</div>
  </div>
</template>
<script>
import axios from "axios";
const appData = {
  message: "",
  message2: ""
};
export default {
  data() {
    return appData;
  },
  methods: {
    showMessage: showMessage,
    addPoint: addPoint
  }
};
function showMessage() {
  axios.get("/api/v1/hello").then(res => {
    console.log(res);
    appData.message = res.data.message;
  });
}
function addPoint(){
  axios.get("/api/v1/add").then(res => {
  console.log(res); 
  appData.message2 = res.data.message;
});
}
</script>
<style>
</style>

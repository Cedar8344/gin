<template>
  <div>
    <div>MyApp</div>
    <button v-on:click="showMessage">I want help</button>
    <div>{{message}}</div>
    <button v-on:click="addPoint">I want die</button>
    <div>{{message2}}</div>
    <form id="app" v-on:submit.prevent>
      <p>
        <label for="name">Id</label>
        <input type="text" name="id" id="id" v-model="id">
      </p>
      <p>
        <label for="name">Username</label>
        <input type="text" name="username" id="username" v-model="username">
      </p>
      <p>
        <label for="name">Password</label>
        <input type="text" name="password" id="password" v-model="password">
      </p>
      <p>
        <label for="name">First Name</label>
        <input type="text" name="firstname" id="firstname" v-model="firstname">
      </p>
      <p>
        <label for="name">Last Name</label>
        <input type="text" name="lastname" id="lastname" v-model="lastname">
      </p>
    </form>
    <button v-on:click="postFunc">Submit</button>
    <div>{{message3}}</div>
    <button v-on:click="updateFunc">Update</button>
    <div>{{message4}}</div>
    <form id="app" v-on:submit.prevent>
      <p>
        <label for="name">Delete ID:</label>
        <input type="text" name="idd" id="idd" v-model="idd">
      </p>
    </form>
    <button v-on:click="delFunc">Delete</button>
    <div>{{message5}}</div>
  </div>
</template>
<script>
import axios from "axios";
const appData = {
  message: "",
  message2: "",
  message3: "",
  message4: "",
  message5: ""
};
export default {
  data:{
    errors:[],
    id:null,
    idd:null,
    username:null,
    password:null,
    firstname:null,
    lastname:null
  },
  data() {
    return appData;
  },
  methods: {
    showMessage: showMessage,
    addPoint: addPoint,
    postFunc: postFunc,
    updateFunc: updateFunc,
    delFunc: delFunc,
    checkForm:function(e) {
      if(this.name && this.age) return true;
      this.errors = [];
      if(!this.name) this.errors.push("Name required.");
      if(!this.age) this.errors.push("Age required.");
      e.preventDefault();
    }
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
function postFunc(){
  appData.message3 = this.username;
  axios.post("/api/v1/post", {
    Id: this.id,
    Username: this.username,
    Password: this.password,
    Firstname: this.firstname,
    Lastname: this.lastname
  }).then((response) => {
  console.log(response);
  });
  this.id = null;
  this.username = null;
  this.password = null;
  this.firstname = null;
  this.lastname = null;
}
function updateFunc(){
  appData.message4 = this.username;
  axios.put("/api/v1/put", {
    Id: this.id,
    Username: this.username,
    Password: this.password,
    Firstname: this.firstname,
    Lastname: this.lastname
  }).then((response) => {
  console.log(response);
  });
  this.id = null;
  this.username = null;
  this.password = null;
  this.firstname = null;
  this.lastname = null;
}
function delFunc(){
  appData.message5 = this.Id;
  axios.delete("api/v1/delete", {
    data: {Id: this.idd}
  }).then((response) => {
  console.log(response);
  });
  this.id = null;
  this.username = null;
  this.password = null;
  this.firstname = null;
  this.lastname = null;
}
</script>
<style>
</style>

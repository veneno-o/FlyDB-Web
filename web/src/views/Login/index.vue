<template>
  <div class="lgoin">
    <div class="form">
      <p>微信登录</p>
      <div class="img">
        <img src="../../assets/code.png" alt="wx code">
      </div>
      <div class="icons">
        <svg class="icon" aria-hidden="true" @click="open('https://github.com/login')">
          <use xlink:href="#icon-github"></use>
        </svg>
        <svg class="icon" aria-hidden="true" @click="open('https://twitter.com/i/flow/login')">
          <use xlink:href="#icon-tuite"></use>
        </svg>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import "../../assets/iconfont.js"
const open = (url:string)=>{
  const a = document.createElement('a');
  a.href = url;
  a.target = '_blank';
  a.click()
}
</script>

<style scoped>
.lgoin{
  width: 100vw;
  height: 100vh;
  background: url('../../assets/bg.webp') center/cover no-repeat;
  display: flex;
  justify-content: end;
  align-items: center;
  box-sizing: border-box;
  padding: 0 200px;
}
.form{
  padding: 20px;
  background-color: var(--bg-color);
  box-shadow: 0 0 var(--box-shadow-range) var(--box-shadow-color);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  border-radius: var(--border-radius);
}
p{
  margin: 0;
  margin-bottom: 8px;
  text-align: center;
}
.form .img{
  width: 300px;
  height: 300px;
}
img{
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.icons{
  margin-top: 10px;
}
.icons .icon{
  width: 26px;
  height: 26px;
  cursor: pointer;
}
.icons .icon:nth-child(2){
  margin-left: 12px;
}
</style>

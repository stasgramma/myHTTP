<script setup>
import { ref, onMounted } from 'vue'

const message = ref('Загрузка...')

onMounted(async () => {
  try {
    // запрос к твоему Go-бэкенду
    const res = await fetch('http://localhost:8080/api/hello')
    const data = await res.text()
    message.value = data
  } catch (err) {
    message.value = 'Ошибка подключения к серверу'
  }
})
</script>

<template>
  <div style="font-family: sans-serif; text-align: center; margin-top: 50px;">
    <h1>Frontend + Go Backend</h1>
    <p>{{ message }}</p>
  </div>
</template>

<style>
body {
  background: #0d1117;
  color: #fff;
}
</style>


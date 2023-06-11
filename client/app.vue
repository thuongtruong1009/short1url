<script setup lang="ts">
import { ref } from 'vue';

const url = ref<string>('');

const shortedUrl = ref<string>('');

const onSubmit = async (e: Event) => {
    if (!url.value) return;
    e.preventDefault();
    const response = await fetch('http://localhost:3000/api', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url: url.value })
    });

    const data = await response.json();

    shortedUrl.value = data.short;
}

</script>

<template>
  <div>
    <header>
        <img src="/favicon.svg" alt="logo">
        <h1>Short1url</h1>
    </header>

    <main>
        <form @submit.prevent="onSubmit">
            <input type="text" name="url" id="url" placeholder="Enter URL" v-model="url">
            <button type="submit" @click="onSubmit" :disabled="url.value">Shorten</button>
        </form>


        <div class="response">
            <p>Your shorted URL:</p>
            <a :href="shortedUrl">{{shortedUrl || 'dedddededed'}}</a>
        </div>
    </main>

    <footer>
        <p>Copyright <a href="https://github.com/thuongtruong1009">@thuongtruong1009</a>, 2023-PRESENT</p>
    </footer>
  </div>
</template>

<style scoped>

@import url('https://fonts.googleapis.com/css2?family=Lobster+Two&display=swap');

body{
    background-color: #f5f5f5;
    font-family: 'Open Sans', sans-serif;
    font-size: 14px;
    color: #333;
    margin: 0;
    padding: 0;
    display: flex;
    overflow-x: hidden;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
}

@media (min-width: 320px) {
    main {
        width: 300px;
    }
}

@media (min-width: 576px) {
    main {
        width: 500px;
    }
}

@media (min-width: 768px) {
    main {
        width: 700px;
    }
}

@media (min-width: 992px) {
    main {
        width: 900px
    }
}

@media (min-width: 1200px) {
    main {
        width: 1100px;
    }
}

input, button {
    outline: none;
    border: none;
    border-radius: 0.4rem;
}

a{
    text-decoration: none;
    color: #9856d6;
}

header {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    font-family: 'Lobster Two', cursive;
}

header img {
    width: 2rem;
    height: 2rem;
    object-fit: cover;
    object-position: center;
    margin-right: 0.5rem;
}

main{
    background: white;
    border-radius: 1rem;
    margin: 2rem auto;
    padding: 1rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

form {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    margin-bottom: 1rem;
    width: 100%;
}

form > input[type="text"] {
    border: 2px solid #ccc;
    padding: 0.4rem 1rem;
    width: 100%;
}

form > input[type="text"]:focus {
    border: 2px solid #007BFF;
}

form > button {
    background-color: #333;
    color: white;
    margin-left: 0.5rem;
    padding: 0.5rem 1rem;
}

form > button:hover {
    background-color: #555;
    cursor: pointer;
}

.response {
    display: flex;
    align-items: center;
}

.response > p {
    margin-right: 0.5rem;
}



footer{
    position: absolute;
    bottom: 0;
    width: 100%;
    text-align: center;
    padding: 1rem;
    font-size: 0.8rem;
}
</style>
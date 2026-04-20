<template>
    <div class="card">
        <slot/>
    </div>
</template>

<style scoped>
.card{
    padding: 5px 10px;
    border: 1px black solid;
    color: black;
    background: #daf6f4;
    backdrop-filter: blur(5px);
}

@media (prefers-color-scheme: dark) {
    .card{
        color: rgba(255, 255, 255, 0.8);
        background: rgba(2, 12, 12, 0.35);
    }
}
</style>
<script setup lang="ts">
</script>
<template>
  <nav class="relative z-30 flex items-center justify-between">
    <!-- Left Logo -->
    <LogoButton logo="B" text="Bhinno" @click="handleLogoClick" />

    <!-- Right Logo -->
    <AvatarButton logo="B" @click="handleAvatarClick" />
  </nav>
</template>

<script setup lang="ts">
import LogoButton from './LogoButton.vue'
import AvatarButton from './AvatarButton.vue'

function handleLogoClick() {
  console.log('Logo clicked!')
}

function handleAvatarClick() {
  console.log('Avatar clicked!')
}
</script>

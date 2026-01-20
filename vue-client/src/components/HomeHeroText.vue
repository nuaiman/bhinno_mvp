<template>
  <div class="flex flex-col items-center justify-center text-center space-y-8 w-full">
    <div
      class="inline-flex items-center px-2 gap-2 bg-blue-500/10 border border-blue-500/20 py-1.5 rounded-full backdrop-blur-md"
    >
      <slot name="badge">
        <svg class="w-3.5 h-3.5 text-blue-400 fill-blue-400" viewBox="0 0 24 24">
          <path
            d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"
          />
        </svg>
        <span class="text-[10px] uppercase tracking-[0.2em] text-blue-400 font-black">
          Trusted Service Discovery Platform
        </span>
      </slot>
    </div>

    <h1
      class="text-4xl md:text-5xl xl:text-7xl font-bold leading-[1.05] tracking-tighter max-w-2xl"
    >
      <slot name="title"
        >Crafted Services for <br /><span
          class="text-transparent bg-clip-text bg-linear-to-br from-blue-400 via-cyan-300 to-emerald-300"
          >Exceptional</span
        >
        Living</slot
      >
    </h1>

    <p class="text-gray-300 text-lg md:text-xl max-w-lg leading-relaxed font-light mx-auto">
      <slot name="description">
        Discover trusted local service providers across your city — carefully curated for quality,
        reliability, and peace of mind.
      </slot>
    </p>
  </div>
</template>

<script setup lang="ts"></script>

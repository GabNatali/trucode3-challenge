<script setup lang="ts">
import { ref } from 'vue';
interface Props {
  open: boolean
  title: string
}

const props = defineProps<Props>()

const openSections = ref<boolean>(props.open);

const toggleSection = () => {
  openSections.value = !openSections.value
}
</script>

<template>
      <div>
        <button
          @click.prevent.stop="toggleSection"
          class="w-full flex justify-between items-center text-left py-2 font-bold"
        >
          {{ props.title }}
          <span>
            <svg
              v-if="!openSections"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2"
              stroke="currentColor"
              class="w-4 h-4"
            >
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 9l6 6 6-6" />
            </svg>
            <svg
              v-else
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="2"
              stroke="currentColor"
              class="w-4 h-4"
            >
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 15l6-6 6 6" />
            </svg>
          </span>
        </button>
        <div v-if="openSections" class="space-y-2 pl-4">
          <slot></slot>
        </div>
      </div>
</template>


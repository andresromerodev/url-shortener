<script setup>
const url = ref('')
const shortUrl = ref('Empty.')

async function shorten() {
  const response = await $fetch('http://localhost:5000/api/v1/shorten', {
    method: 'POST',
    body: { longUrl: url.value }
  })
  shortUrl.value = response.shortUrl
}

function clear() {
  url.value = ''
  shortUrl.value = 'Empty.'
}

</script>

<template>
  <UContainer>
    <br />
    <h1>URL Shortener</h1>
    <UInput v-model="url" size="sm" @keyup.enter="shorten" />
    <br />
    <UButton @click="shorten">Shorten</UButton>
    <UButton style="margin-left: 4px;" color="red" variant="solid" @click="clear">Clear</UButton>
    <br />
    <br />
    <UCard>
      <ULink
        class="h-32"
        to="/elements/link"
        active-class="text-primary"
        inactive-class="text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-200"
      >
        {{shortUrl}}
    </ULink>
    </UCard>
    <br />
    <UTable :rows="[]" />
    <br />
  </UContainer>
</template>

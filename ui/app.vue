<script setup>
const URL_EMPTY_MESSAGE = "Your new shortened URL will appear here once it's ready."

const url = ref('')
const ip = ref('www')
const shortUrl = ref(URL_EMPTY_MESSAGE)

const userInformation = await $fetch('https://api.ipify.org?format=json')

ip.value = userInformation.ip

async function shorten() {
  const response = await $fetch('http://localhost:5000/api/v1/shorten', {
    method: 'POST',
    body: { longUrl: url.value }
  })
  shortUrl.value = response.shortUrl
}

function clear() {
  url.value = ''
  shortUrl.value = URL_EMPTY_MESSAGE
}

</script>

<template>
  <UContainer>
    <br />
    <h1 style="font-weight: bold;">🔗 URL Shortener</h1>
    <UInput v-model="url" size="sm" @keyup.enter="shorten" placeholder="Enter the URL you want to shorten"/>
    <br />
    <UButton @click="shorten">Shorten</UButton>
    <UButton style="margin-left: 4px;" color="red" variant="solid" @click="clear">Clear</UButton>
    <br />
    <br />
    <h1 style="font-weight: bold;">📤 Copy and Share the Link Below</h1>
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
    <h1 style="font-weight: bold;">🌍 Your Public IP Address is: {{ip}}</h1>
    <br />
    <p>
      We use your public IP address to show a personalized history of the URLs you have shortened. This helps you easily track and manage your links. Rest assured, your privacy is important, and this information is not shared externally.
    </p>
    <br />
    <UTable :rows="[]" />
    <br />
  </UContainer>
</template>

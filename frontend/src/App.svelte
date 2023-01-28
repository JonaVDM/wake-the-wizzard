<script lang="ts">
  import { onMount } from 'svelte';
  import { loadDevices, devices } from './lib/api';
  import DevideCard from './lib/DevideCard.svelte';

  let error = '';

  onMount(async () => {
    try {
      await loadDevices();
    } catch (e) {
      error = 'Error while getting data from the server';
    }
  });
</script>

<main class="container">
  <h1>Wake On Lan</h1>
  <button>Add new</button>

  {#if error}
    <article>
      {error}
    </article>
  {/if}

  {#each $devices as device (device.id)}
    <DevideCard {...device} />
  {/each}
</main>

<style>
  h1 {
    margin-top: 2rem;
    margin-bottom: 2rem;
  }

  article {
    background-color: rgb(179, 0, 0);
    color: white;
  }
</style>

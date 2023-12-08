<script lang="ts">
  import { onMount } from "svelte";
  import { loadDevices, devices } from "./lib/api";
  import DevideCard from "./lib/DevideCard.svelte";
  import NewDevice from "./lib/NewDevice.svelte";
  import Notification from "./lib/Notification.svelte";

  let error = "";
  let showNew = false;

  onMount(async () => {
    try {
      await loadDevices();
    } catch (e) {
      error = "Error while getting data from the server";
    }
  });
</script>

<main class="container mx-auto p-3">
  <div class="flex justify-between">
    <h1 class="text-title font-semibold text-center md:text-left">
      Wake On Lan
    </h1>

    <button on:click={() => (showNew = !showNew)} class="btn btn-secondary"
      >Add new!</button
    >
  </div>

  <Notification />

  {#if showNew}
    <div class="py-2">
      <NewDevice on:close={() => (showNew = false)} />
    </div>
  {/if}

  {#if error}
    <article>
      {error}
    </article>
  {/if}

  {#each $devices as device (device.id)}
    <div class="py-4">
      <DevideCard {...device} />
    </div>
  {/each}
</main>

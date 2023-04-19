<script lang="ts">
  import { slide } from 'svelte/transition';
  import { deleteDevice, wakeDevice } from './api';

  export let id: string;
  export let name: string;
  export let mac: string;

  let loadingDelete = false;
  let loadingWaking = false;
  let showDelete = false;

  const wakePc = async () => {
    loadingWaking = true;
    await wakeDevice(id);
    loadingWaking = false;
  };

  const deletePc = async () => {
    loadingDelete = true;
    await deleteDevice(id);
  };

  const toggleDelete = () => {
    showDelete = !showDelete;
  };
</script>

<div class="bg-base-200 p-4 rounded-lg shadow-md">
  <div class="md:flex justify-between my-2">
    <h2 class="text-lg font-bold">{name}</h2>
    {#if !showDelete}
      <div class="flex gap-2" transition:slide>
        <button class="btn btn-error block flex-grow" on:click={toggleDelete}
          >Delete</button
        >
        <button
          class="btn btn-primary block flex-grow"
          aria-busy={loadingWaking}
          on:click={wakePc}>Wake</button
        >
      </div>
    {/if}
  </div>

  {#if showDelete}
    <div transition:slide>
      <p class="font-bold">Are you sure you want to delete this entry?</p>

      <div class="flex gap-2">
        <button
          aria-busy={loadingDelete}
          on:click={deletePc}
          class="btn btn-error max-md:flex-grow"
        >
          Yes, delete it
        </button>
        <button on:click={toggleDelete} class="btn btn-primary max-md:flex-grow"
          >No, keep it</button
        >
      </div>
    </div>
  {/if}

  <div class="flex flex-col md:flex-row justify-between my-2">
    <p>{mac}</p>
    <p>{id}</p>
  </div>
</div>

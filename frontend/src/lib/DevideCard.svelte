<script lang="ts">
  import { slide } from 'svelte/transition';
  import { deleteDevice, wakeDevice } from './api';

  export let id: string;
  export let name: string;
  export let mac: string;

  let loadingDelete = false;
  let loadingWaking = false;

  const wakePc = async () => {
    loadingWaking = true;
    await wakeDevice(id);
    loadingWaking = false;
  };

  const deletePc = async () => {
    loadingDelete = true;
    await deleteDevice(id);
  };

  let showDelete = false;
  const toggleDelete = () => {
    showDelete = !showDelete;
  };
</script>

<div class="container">
  <article>
    <div class="flex">
      <p>{name}</p>
      <div class="grid half">
        <button class="warning" on:click={toggleDelete}>Delete</button>
        <button aria-busy={loadingWaking} on:click={wakePc}>Wake</button>
      </div>
    </div>

    {#if showDelete}
      <div transition:slide class="delete">
        <p>
          <strong>Are you sure you want to delete this entry?</strong>
        </p>

        <div class="grid">
          <button aria-busy={loadingDelete} on:click={deletePc} class="warning">
            Yes, delete it
          </button>
          <button on:click={toggleDelete}>No, keep it</button>
        </div>
      </div>
    {/if}

    <footer class="flex">
      <p>{mac}</p>
      <p>{id}</p>
    </footer>
  </article>
</div>

<style>
  p,
  button {
    margin-bottom: 0;
    border: 0;
  }
  .flex {
    display: flex;
    justify-content: space-between;
  }

  .half {
    width: 50%;
  }

  .warning {
    background-color: rgb(202, 6, 6);
  }

  .delete {
    padding-top: 3rem;
  }
</style>

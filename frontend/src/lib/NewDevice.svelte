<script lang="ts">
  import { slide, fly } from 'svelte/transition';
  import { addDevice } from './api';

  let showForm = false;
  let mac = '';
  let name = '';

  const submit = () => {
    addDevice(name, mac);
  };
</script>

<article>
  <div class="flex">
    <h3>Add a new device</h3>
    {#if !showForm}
      <button transition:fly class="half" on:click={() => (showForm = true)}>
        Show
      </button>
    {/if}
  </div>
  <form>
    {#if showForm}
      <div transition:slide>
        <div class="grid">
          <label>
            Name
            <input type="text" bind:value={name} />
          </label>

          <label>
            Mac
            <input type="text" bind:value={mac} />
          </label>
        </div>

        <button on:click|preventDefault={submit}>Create</button>
      </div>
    {/if}
  </form>
</article>

<style>
  .flex {
    display: flex;
    justify-content: space-between;
  }

  .half {
    width: 50%;
  }
</style>

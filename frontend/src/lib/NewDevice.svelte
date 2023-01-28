<script lang="ts">
  import { slide, fly } from 'svelte/transition';
  import { addDevice } from './api';

  let showForm = false;
  let mac = '';
  let name = '';
  let loading = false;
  let error = '';

  const submit = async () => {
    try {
      loading = true;
      await addDevice(name, mac);
      loading = false;
      showForm = false;
    } catch (e) {
      loading = false;
      error = 'Could not create new entry';
    }
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

        {#if error != ''}
          <p class="error">{error}</p>
        {/if}
        <button aria-busy={loading} on:click|preventDefault={submit}>
          Create
        </button>
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

  .error {
    color: rgb(192, 2, 2);
  }
</style>

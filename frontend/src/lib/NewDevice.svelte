<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { slide } from "svelte/transition";
  import { addDevice } from "./api";

  let mac = "";
  let name = "";
  let loading = false;
  let error = "";

  const dispatch = createEventDispatcher();

  const submit = async () => {
    try {
      loading = true;
      await addDevice(name, mac);
      loading = false;
      dispatch("close");
    } catch (e) {
      loading = false;
      error = e;
    }
  };
</script>

<div class="bg-base-200 p-4 rounded-lg shadow-md">
  <form on:submit|preventDefault={submit}>
    <div transition:slide>
      <label class="my-2 block">
        Name
        <input
          class="input input-bordered w-full"
          type="text"
          bind:value={name}
        />
      </label>

      <label class="my-2 block">
        Mac
        <input
          type="text"
          bind:value={mac}
          class="input input-bordered w-full"
        />
      </label>

      {#if error != ""}
        <p class="text-error">{error}</p>
      {/if}
      <button class="btn btn-primary my-2">Create</button>
    </div>
  </form>
</div>

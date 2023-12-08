<script lang="ts">
  import { onMount } from "svelte";
  import { slide } from "svelte/transition";
  import { error, message } from "../lib/api";

  onMount(() => {
    const timeout = 5000;

    const unsubError = error.subscribe((str) => {
      if (str && str != "") {
        setTimeout(() => {
          error.set("");
        }, timeout);
      }
    });

    const unsubMessage = message.subscribe((str) => {
      if (str && str != "") {
        setTimeout(() => {
          message.set("");
        }, timeout);
      }
    });

    return () => {
      unsubError();
      unsubMessage();
    };
  });
</script>

{#if $error}
  <div class="alert alert-error my-2" transition:slide>
    <span>
      {$error}
    </span>
  </div>
{/if}

{#if $message}
  <div class="alert alert-success my-2" transition:slide>
    <span>
      {$message}
    </span>
  </div>
{/if}

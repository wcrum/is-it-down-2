<script>
  import { onMount } from "svelte";

  let apiData;

  onMount(() => {
    const interval = setInterval(async () => {
      const response = await fetch("/your-api-endpoint");
      apiData = await response.json();
    }, 5000); // Fetch every 5 seconds

    return () => clearInterval(interval); // Cleanup on component unmount
  });
</script>

<h1>Welcome to SvelteKit</h1>
<p>
  Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation
</p>

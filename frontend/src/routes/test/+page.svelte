<script>
  import { fromTheme } from "tailwind-merge";
  import Chart from "./Chart.svelte";
  import { get } from "svelte/store";

  async function getData(url) {
    try {
      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`Response status: ${response.status}`);
      }

      const json = await response.json();

      return {datasets: [{data: json}]};
    } catch (error) {
      console.error(error.message);
    }
    return;
  }

  let test = {
    google: getData("http://localhost:8080/latency/google"),
    insta: getData("http://localhost:8080/latency/insta")
  }
</script>

<body>
  {#await test.google}
    <p>...waiting</p>
  {:then data}
    <Chart {data} />
  {:catch error}
    <p style="color: red">{error.message}</p>
  {/await}
  {#await test.insta}
    <p>...waiting</p>
  {:then data}
    <Chart {data} />
  {:catch error}
    <p style="color: red">{error.message}</p>
  {/await}
</body>

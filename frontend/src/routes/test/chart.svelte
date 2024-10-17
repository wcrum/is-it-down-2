<script>
  import { Line } from "svelte-chartjs";
  import "moment"
  import "chartjs-adapter-moment"

  export let data;

  import 'chart.js/auto';

</script>

<body>
  <Line {data} options={{ responsive: true, scales: {x: { type: "time" }}}}/>
</body>

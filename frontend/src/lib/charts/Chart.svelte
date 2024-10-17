<script context="module">
	import Chart from 'chart.js/auto'
</script>

<script>
	export let config
	
	const handleChart = (element, config) => {
		let theChart = new Chart(element, config)
		
		return {
			update(config) {
				theChart.destroy()
				theChart = new Chart(element, config)
			},
			destroy() {
				theChart.destroy()
			}
		}
	}
</script>

<body>
    <canvas use:handleChart={config} />
</body>
var ctx = document.getElementById("lang-chart");

var myChart = new Chart(ctx, {
	type: 'bar',
	data: {
		labels: ["Red", "Blue", "Yellow"],
		datasets: [{
			labels: "# of Votes",
			data: [12, 10, 3],
			backgroundColor: [
				'rgba(255, 99, 132, 0.2)',
				'rgba(54, 162, 235, 0.2)',
				'rgba(255, 99, 132, 0.2)'
			], 
		}]
	}
});

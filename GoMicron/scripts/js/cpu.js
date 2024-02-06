function getCPUData(data) {
  var CPU_Percent = data["CPU_Percent"];
  performanceData.push(CPU_Percent);
  chart.data.datasets[0].data = performanceData.slice(-50);
  chart.update();
}
const performanceData = [];
const chart = new Chart(document.getElementById("line-chart"), {
  width: 300,
  height: 200,
  type: "line",
  data: {
      labels: Array(50).fill(""),
      datasets: [{
          data: performanceData,
          borderColor: "#0e4768b7",
          fill: true,
          backgroundColor: "#3C9CCE",
          tension: 0.5,
          cubicInterpolationMode: "monotone"
      }]
  },
  options: {
      scales: {
          y: {
              min: 0
          }
      },
      title: {
          display: false
      },
      plugins: {
          legend: {
              display: false
          }
      },
      animation: {},
      elements: {
          point: {
              radius: 0,
              opacity: 0
          }}}});
function getRamData(data) {
    var RAM_Percent = data["RAM_Percent"];
    memoryData.push(RAM_Percent);
    if (memoryData.length > 8) {
        memoryData.shift(); }
    const startLabel = memoryData.length > 8 ? memoryData.length - 7 : 1;
    memchart.data.labels = Array.from({ length: 8 }, (_, i) => `Point ${startLabel + i}`);
    memchart.data.datasets[0].data = memoryData;
    memchart.update();
}
const memoryData = [];
const canvas = document.getElementById("bar-chart");
const ctx = canvas.getContext("2d");
const memchart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: Array.from({ length: 8 }, (_, i) => `Point ${i + 1}`),
        datasets: [{
            label: 'Memory Performance',
            data: memoryData,
            backgroundColor: '#355f778b', 
            borderColor: '#0e4768b7',
            borderWidth: 3,
        }]
    },
    options: {
        scales: {
            y: {
                min: 0,
                max: 100,
                ticks: {
                    stepSize: 5
                },
                title: {
                    display: true,
                    text: 'Performance (%)'
                }
            },
            x: {
                title: {}
            }}}});
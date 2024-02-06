const memoryData = [70, 80, 60, 90, 70, 80, 60, 90];
const canvas = document.getElementById("mem-chart");
const ctx = canvas.getContext("2d");
const memchart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: Array.from({ length: memoryData.length }, (_, i) => `Point ${i + 1}`),
        datasets: [{
            label: 'Memory Performance',
            data: memoryData,
            backgroundColor: '#355f778b', 
            borderColor: '#0e4768b7',
            borderWidth: 3,
        }]},
    options: {
        scales: {
            y: {
                min: 0,
                max: 100,
                ticks: {
                    stepSize: 5 },
                title: {
                    display: true,
                    text: 'Performance (%)'}},
            x: {
                title: {}}}}});

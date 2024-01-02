const ctx = document.getElementById('myChart').getContext('2d');
const diskData = {labels: ['Free Disk Space ', 'Used Disk Space'],
    datasets: [{
        data: [474000, 202000], 
        backgroundColor: ['#355f778b', '#0e4768b7'],}],};
const chart = new Chart(ctx, {
    type: 'pie',
    data: diskData,
    options: {
        maintainAspectRatio: true, 
        responsive: true, 
        scales: {
            y: {
                display: false,
            },x: {
                display: false,}},
        title: {
            display: false,},
        plugins: {
            legend: {
                position: 'bottom',},},
        hoverOffset: 5,
        animation: {
          animateRotate: true, 
          animateScale: true, },
        elements: {
            point: {
                radius: 0,
                opacity: 0}}}});
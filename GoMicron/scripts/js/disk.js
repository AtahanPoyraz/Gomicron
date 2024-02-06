function getDiskData(data) {
    const Disk_Percent = data["Disk_Percent"];

    const usedSpacePercent = Disk_Percent;
    const freeSpacePercent = 100 - Disk_Percent;

    piechart.data.datasets[0].data = [freeSpacePercent, usedSpacePercent];
    piechart.update();
}
const Ctx = document.getElementById('pie-chart').getContext('2d');
const diskData = {
    labels: ['Free Space % ', 'Used Space %'],
    datasets: [{
        data: [0, 0],
        backgroundColor: ['#355f77', '#0e4768']
    }]
};
const piechart = new Chart(Ctx, {
    type: 'pie',
    data: diskData,
    options: {
        maintainAspectRatio: true,
        responsive: true,
        plugins: {
            legend: {
                position: 'bottom',
            },
        },
        hoverOffset: 5,
        animation: {
            animateRotate: true,
            animateScale: true,
        },
        elements: {
            point: {
                radius: 0,
                opacity: 0
            }}}});
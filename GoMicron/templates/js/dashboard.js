const blob = new Blob(
    [
      `let time = performance.now()
             let iterations = 0;
             while(true){
                  iterations++;
                  let now = performance.now()
                  if(now - time > 100){
                      postMessage(iterations)
                      time = performance.now()
                      iterations = 0
                  }
             }`,
    ], {
      type: "text/javascript"
    }
  );
  
  const processWorker = new Worker(window.URL.createObjectURL(blob));
  
  const performanceData = [];
  
  processWorker.onmessage = (e) => {
    const cpuUsage = +e.data / 100000
    performanceData.push(cpuUsage);
    chart.data.datasets[0].data = performanceData.slice(-50);
    chart.update();
  };
  
  const chart = new Chart(document.getElementById("line-chart"), {
    width: 400,
    height: 200,
    type: "line",
    data: {
      labels: Array(50).fill(""),
      datasets: [{
        data: performanceData,
        borderColor: "#0e4768b7",
        fill: !0,
        fillColor: "#3C9CCE",
        tension: 10,
        cubicInterpolationMode: "monotone"
      }]
    },
    options: {
      scales: {
        y: {
          min: 0,
        }
      },
      title: {
        display: !1
      },
      plugins: {
        legend: {
          display: !1
        }
      },
      animation: {},
      elements: {
        point: {
          radius: 0,
          opacity: 0
        }
      }
    }
  });
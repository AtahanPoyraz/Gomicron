const { exec } = require('child_process');

function getCPUInfo(callback) {
    const CPU_INFO = {};
    exec('python scripts/cpu.py', (error, stdout, stderr) => {
        if (error) {
            console.error(`ERROR in getCPUInfo: ${error.message}`);
            callback(error, null);
            return;
        }

        const info = JSON.parse(stdout);
        CPU_INFO['CPU'] = info;
        
        callback(null, CPU_INFO);
    });
}

module.exports = { getCPUInfo };

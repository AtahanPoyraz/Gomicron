const { exec } = require('child_process');

function getMemoryInfo(callback) {
    const MEMORY_INFO = {};
    exec('python scripts/memory.py', (error, stdout, stderr) => {
        if (error) {
            console.error(`ERROR in getMemoryInfo: ${error.message}`);
            callback(error, null);
            return;
        }

        const info = JSON.parse(stdout);
        MEMORY_INFO['Memory'] = info;
        
        callback(null,  MEMORY_INFO);
    });
}

module.exports = { getMemoryInfo };
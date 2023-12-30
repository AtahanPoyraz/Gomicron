const { exec } = require('child_process');

function getDiskInfo(callback) {
    const DISK_INFO = {};
    exec('python scripts/disk.py', (error, stdout, stderr) => {
        if (error) {
            console.error(`ERROR in getDiskInfo: ${error.message}`);
            callback(error, null);
            return;
        }

        const info = JSON.parse(stdout);
        DISK_INFO['Disk'] = info;

        callback(null, DISK_INFO);
    });
}

module.exports = { getDiskInfo };
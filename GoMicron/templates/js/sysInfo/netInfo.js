const { exec } = require('child_process');

function getNetInfo(callback) {
    const NET_INFO = {};
    exec('python scripts/net.py', (error, stdout, stderr) => {
        if (error) {
            console.error(`ERROR in getNetInfo: ${error.message}`);
            callback(error, null);
            return;
        }

        const info = JSON.parse(stdout);
        NET_INFO['Net'] = info;

        callback(null, NET_INFO);
    });
}

module.exports = { getNetInfo };
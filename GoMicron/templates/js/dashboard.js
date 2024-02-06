//const baseUrl = 'http://localhost:1357';
var baseUrl = 'http://' + window.location.hostname + ':' + window.location.port;
function getStats(cookie) {
    fetch(baseUrl + '/gomicron/server/stats/', {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'same-origin',
    }).then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error('Network response was not ok');
        }
    }).then(data => {
        getCPUData(data);
        getRamData(data);
        getDiskData(data);
    }).catch(error => {
        console.log('Fetch error:', error);
    });
}
const cookies = document.cookie.split('; ').reduce((prev, current) => {
    const [name, value] = current.split('=');
    prev[name] = value;
    return prev;
}, {});
let intervalID;
const intervalInputs = document.getElementsByClassName('updateInterval');
for (let i = 0; i < intervalInputs.length; i++) {
    const intervalValue = intervalInputs[i].value;
    const cookie = cookies['G0M1CR0N4UTHK3Y'];
    function handleInputChange() {
        const newIntervalValue = this.value;

        clearInterval(intervalID);

        intervalID = setInterval(() => {
            getStats(cookie);
        }, newIntervalValue * 1325);}
    handleInputChange.call(intervalInputs[i]);
    intervalInputs[i].addEventListener('input', handleInputChange);}

import psutil
import json

def get_network_info():
    network_info = psutil.net_io_counters(pernic=True)

    network_info_json = {}
    for _, stats in network_info.items():
        network_info_json = {
            "NETbytes_sent": stats.bytes_sent,
            "NETbytes_recv": stats.bytes_recv,
            "NETpackets_sent": stats.packets_sent,
            "NETpackets_recv": stats.packets_recv,
        }

    print(json.dumps(network_info_json, indent=2))

if __name__ == "__main__":
    get_network_info()

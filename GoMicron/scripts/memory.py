import psutil
import json

def get_memory_info():
    memory_info = psutil.virtual_memory()

    swap_info = psutil.swap_memory()

    memory_info_json = {
        "MEMTotal": memory_info.total,
        "MEMAvailable": memory_info.available,
        "MEMUsed": memory_info.used,
        "MEMPercent": memory_info.percent,
        "MEMSwap_Total": swap_info.total,
        "MEMSwap_Used": swap_info.used,
        "MEMSwap_Free": swap_info.free,
        "MEMSwap_Percent": swap_info.percent
    }

    print(json.dumps(memory_info_json, indent=2))

if __name__ == "__main__":
    get_memory_info()
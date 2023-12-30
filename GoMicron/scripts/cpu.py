import psutil
import platform
import os
import json

def get_cpu_info():
    cpu_model = platform.processor()
    cpu_count = os.cpu_count()
    cpu_usage = psutil.cpu_percent(interval=1)
    cpu_freq = psutil.cpu_freq().current

    cpu_info_json = {
        "CPUModel": cpu_model,
        "CPUCount": cpu_count,
        "CPUUsage": cpu_usage,
        "CPUFreq": cpu_freq
    }

    print(json.dumps(cpu_info_json, indent=2))

if __name__ == "__main__":
    get_cpu_info()
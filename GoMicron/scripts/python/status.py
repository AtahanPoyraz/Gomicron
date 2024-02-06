import psutil as ps
import json
import os
import platform

def get_rootdir() -> str:
    if platform.system().lower() == "windows":
        return os.path.dirname(os.path.abspath(__file__)).split("\\")[0] + "\\"
    else:
        return os.path.dirname(os.path.abspath(__file__)).split("/")[0] + "/"

ROOT_DIR = get_rootdir()

class PerformanceStats:

    @staticmethod
    def CPU_Usage() -> float:
        cpu_percent = ps.cpu_percent(0)
        return cpu_percent

    @staticmethod
    def RAM_Usage() -> float:
        ram_percent = ps.virtual_memory().percent
        return ram_percent

    @staticmethod
    def Disk_Usage() -> float:
        disk_percent = ps.disk_usage(ROOT_DIR).percent
        return disk_percent

if __name__ == "__main__":
    stats = {
        "CPU_Percent": PerformanceStats.CPU_Usage(),
        "RAM_Percent": PerformanceStats.RAM_Usage(),
        "Disk_Percent": PerformanceStats.Disk_Usage()
    }

    print(json.dumps(stats))
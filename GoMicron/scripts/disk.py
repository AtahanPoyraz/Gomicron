import psutil
import json

def get_disk_info():
    disk_partitions = psutil.disk_partitions(all=True)
    disk_info = {}

    for partition in disk_partitions:
        usage = psutil.disk_usage(partition.mountpoint)

        disk_info = {
            "DiskTotal": usage.total,
            "DiskUsed": usage.used,
            "DiskFree": usage.free,
            "DiskPercent": usage.percent
        }

    print(json.dumps(disk_info, indent=2))

if __name__ == "__main__":
    get_disk_info()
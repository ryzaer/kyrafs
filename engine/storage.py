import shutil
import re

DEFAULT_RESERVED_FREE_SPACE = 20 * 1024 * 1024  # 20 MB


def parse_size(value: str) -> int:
    """
    Convert:
        20MB
        1GB
        500KB
        2TB

    return bytes
    """

    if value is None:
        return DEFAULT_RESERVED_FREE_SPACE

    value = value.strip()

    if value == "":
        return DEFAULT_RESERVED_FREE_SPACE

    m = re.match(r"^(\d+)\s*(KB|MB|GB|TB)$", value.upper())

    if not m:
        return DEFAULT_RESERVED_FREE_SPACE

    number = int(m.group(1))
    unit = m.group(2)

    table = {
        "KB": 1024,
        "MB": 1024 ** 2,
        "GB": 1024 ** 3,
        "TB": 1024 ** 4,
    }

    return number * table[unit]


def get_disk_info(path):
    """
    Return:
        total
        used
        free
    """

    usage = shutil.disk_usage(path)

    return {
        "total": usage.total,
        "used": usage.used,
        "free": usage.free
    }


def check_reserved_free_space(volume_path, upload_size, reserved_value):

    reserved = parse_size(reserved_value)

    disk = get_disk_info(volume_path)

    remain = disk["free"] - upload_size

    if remain < reserved:

        reserved_mb = reserved / (1024 * 1024)

        return {
            "success": False,
            "error": {
                "code": "STORAGE001",
                "message": f"Remaining storage space is below the reserved free space ({reserved_mb:.0f} MB). Please add a new volume in kyrafs.ini."
            }
        }

    return {
        "success": True
    }
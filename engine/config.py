#!/usr/bin/env python3

import configparser
import os

CONFIG_FILE = os.path.join(
    os.path.dirname(os.path.dirname(__file__)),
    "product",
    "win",
    "kyrafs.ini"
)


class Config:

    def __init__(self):

        self.cfg = configparser.ConfigParser()
        self.cfg.read(CONFIG_FILE)

    def access_key(self):

        return self.cfg.get(
            "storage",
            "access_key",
            fallback=""
        ).strip()

    def reserved_free_space(self):

        return self.cfg.get(
            "storage",
            "reserved_free_space",
            fallback=""
        ).strip()

    def volumes(self):

        result = []

        for section in self.cfg.sections():

            if not section.startswith("volume:"):
                continue

            name = section.split(":", 1)[1]

            path = self.cfg.get(
                section,
                "path",
                fallback=""
            ).strip()

            result.append({
                "name": name,
                "path": path
            })

        return result
    
from storage import check_reserved_free_space


def select_volume(cfg, upload_size):

    last_error = None

    for volume in cfg.volumes():

        path = volume["path"]

        if not os.path.isdir(path):
            continue

        result = check_reserved_free_space(
            path,
            upload_size,
            cfg.reserved_free_space()
        )

        if result["success"]:
            return {
                "success": True,
                "name": volume["name"],
                "path": path
            }

        last_error = result

    if last_error:
        return last_error

    return {
        "success": False,
        "error": {
            "code": "STORAGE002",
            "message": "No available storage volume found. Please check kyrafs.ini."
        }
    }

# if __name__ == "__main__":
#     cfg = Config()
#     print(cfg.volumes())
#     print(select_volume(cfg, 1024))
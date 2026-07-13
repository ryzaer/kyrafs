import json
import os
import time

from config import Config, select_volume
from hash import generate_hash, generate_privilege_key
from helper import (
    mime_type,
    create_storage,
    move_file,
    write_meta,
)

def run_put(filepath):

    if not os.path.isfile(filepath):
        print(json.dumps({
            "success": False,
            "error": {
                "code": "PUT001",
                "message": "File not found."
            }
        }))
        return

    filename = os.path.basename(filepath)
    filesize = os.path.getsize(filepath)

    cfg = Config()

    volume = select_volume(cfg, filesize)

    if not volume["success"]:
        print(json.dumps(volume))
        return

    hash_value = generate_hash()

    privilege_key = generate_privilege_key()

    folder = create_storage(
        volume["path"],
        hash_value
    )

    move_file(
        filepath,
        folder
    )

    meta = {
        "hash": hash_value,
        "filename": filename,
        "mime": mime_type(filename),
        "size": filesize,
        "secure_key": "",
        "privilege_key": privilege_key,
        "created_at": int(time.time()),
        "last_access": 0,
        "download_count": 0
    }

    write_meta(
        folder,
        meta
    )

    print(json.dumps({
        "success": True,
        "hash": hash_value,
        "filename": filename,
        "mime": meta["mime"],
        "size": filesize,
        "secure_key": "",
        "privilege_key": privilege_key,
        "saved_to": volume["name"]
    }))
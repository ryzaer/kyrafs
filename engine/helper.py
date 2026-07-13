#!/usr/bin/env python3

import json
import mimetypes
import os
import shutil


def hash_to_path(hash_value: str) -> str:
    """
    A10Bf3n5jk7hko9Pm2QaX7Ld

    ->

    A10B/f3n5/jk7hko9Pm2QaX7Ld
    """

    return os.path.join(
        hash_value[:4],
        hash_value[4:8],
        hash_value[8:]
    )


def create_storage(volume_path: str, hash_value: str):

    folder = os.path.join(
        volume_path,
        hash_to_path(hash_value)
    )

    os.makedirs(folder, exist_ok=True)

    return folder


def move_file(source: str, destination_folder: str):

    destination = os.path.join(
        destination_folder,
        "file"
    )

    shutil.move(source, destination)

    return destination


def mime_type(filename: str):

    mime, _ = mimetypes.guess_type(filename)

    if mime is None:
        return "application/octet-stream"

    return mime


def write_meta(folder: str, meta: dict):

    meta_file = os.path.join(
        folder,
        "meta.json"
    )

    with open(meta_file, "w", encoding="utf-8") as fp:
        json.dump(
            meta,
            fp,
            indent=4,
            ensure_ascii=False
        )


def read_meta(folder: str):

    meta_file = os.path.join(
        folder,
        "meta.json"
    )

    with open(meta_file, "r", encoding="utf-8") as fp:
        return json.load(fp)
    
# if __name__ == "__main__":

#     h = "A10Bf3n5jk7hko9Pm2QaX7Ld"

#     print(hash_to_path(h))
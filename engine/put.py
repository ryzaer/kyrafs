#!/usr/bin/env python3

import json
import os
import sys


def main():

    if len(sys.argv) != 3:
        print(json.dumps({
            "success": False,
            "message": "FILE_REQUIRED"
        }))
        sys.exit(1)

    filepath = sys.argv[2]

    if not os.path.isfile(filepath):
        print(json.dumps({
            "success": False,
            "message": "FILE_NOT_FOUND"
        }))
        sys.exit(1)

    filename = os.path.basename(filepath)
    size = os.path.getsize(filepath)

    result = {
        "success": True,
        "filename": filename,
        "size": size,
        "message": "Python engine ready"
    }

    print(json.dumps(result))


if __name__ == "__main__":
    main()
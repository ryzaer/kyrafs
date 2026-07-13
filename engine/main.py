#!/usr/bin/env python3

import sys

from put import run_put

def main():

    if len(sys.argv) < 3:
        print('{"success":false,"message":"INVALID_ARGUMENT"}')
        return

    command = sys.argv[1]

    if command == "put":
        run_put(sys.argv[2])
        return

    print('{"success":false,"message":"UNKNOWN_COMMAND"}')

if __name__ == "__main__":
    main()
#!/usr/bin/env python3

import sys

from put import main as put


def main():

    if len(sys.argv) < 2:
        print("Command required")
        sys.exit(1)

    command = sys.argv[1]

    if command == "put":
        put()
    else:
        print("Unknown command")


if __name__ == "__main__":
    main()
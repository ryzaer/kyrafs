#!/usr/bin/env python3

import secrets
import string

BASE62 = string.ascii_uppercase + string.ascii_lowercase + string.digits


def random_string(length: int) -> str:
    return ''.join(secrets.choice(BASE62) for _ in range(length))


def split4(value: str) -> str:
    """
    ABCDEFGHIJKLMNOP
    ->
    ABCD-EFGH-IJKL-MNOP
    """
    return "-".join(
        value[i:i + 4]
        for i in range(0, len(value), 4)
    )


def generate_hash() -> str:
    """
    24 chars

    Example:
    A10Bf3n5jk7hko9Pm2QaX7Ld
    """
    return random_string(24)


def generate_privilege_key() -> str:
    """
    Example:
    PK-W8N2-XK4M-Q9P7
    """
    return "PK-" + split4(random_string(12))


def generate_secure_key() -> str:
    """
    Example:
    SK-BM7P-VT2X-K8Q4
    """
    return "SK-" + split4(random_string(12))

# if __name__ == "__main__":

#     print(generate_hash())
#     print(generate_privilege_key())
#     print(generate_secure_key())
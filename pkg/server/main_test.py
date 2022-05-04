import socket
from datetime import date, datetime
from ipaddress import IPv4Address


def time_to_int(t: datetime) -> int:
    return int(datetime.now().timestamp())


def int_to_time(n: int) -> datetime:
    return datetime.fromtimestamp(n)

def int_to_hex(n: int) -> str:
    return n.to_bytes(4, "big").hex()

def hex_to_int(s: str) -> int:
    return int.from_bytes(bytes.fromhex(s), "big")

def time_to_hex(t: datetime) -> str:
    return int_to_hex(time_to_int(t))


def hex_to_time(s: str) -> datetime:
    return int_to_time(hex_to_int(s))


def ip_to_hex(s: str) -> str:
    return int_to_hex(IPv4Address(s))

def hex_to_ip(s: str) -> str:
    return str(IPv4Address(hex_to_int(s)))


if __name__ == "__main__":
    # conn = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, 0)
    # conn.connect(("127.0.0.1", 30514))
    # conn.send("你好".encode("utf-8"))
    # conn.close()
    print(int_to_hex(6147900))

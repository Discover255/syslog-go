import socket


if __name__ == "__main__":
    conn = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, 0)
    conn.connect(("127.0.0.1", 30514))
    conn.send("你好".encode("utf-8"))
    conn.close()

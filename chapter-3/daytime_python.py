#!/usr/bin/env python3

from datetime import datetime
import socket

HOST = '127.0.0.1'  # Standard loopback interface address (localhost)
PORT = 37        # Port to listen on (non-privileged ports are > 1023)

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
	s.bind((HOST, PORT))
	s.listen()

	print("Listen on :37")
	while True:
		conn, addr = s.accept()
		print("Accept new connection")
		with conn:
			print('Connected by', addr)
			conn.sendall(datetime.utcnow().isoformat().encode())
		conn.close()
		print("Connection closed")


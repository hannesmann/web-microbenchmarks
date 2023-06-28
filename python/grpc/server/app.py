import grpc

from service_pb2 import *
from service_pb2_grpc import *
from concurrent.futures import *

class BenchmarkImpl(BenchmarkService):
	def Benchmark(self, _request, _context):
		return Response(data="r")

GRPC_ADDRESS = "127.0.0.1"
GRPC_PORT = "9500"

if __name__ == "__main__":
	try: 
		print("Python gRPC server started")

		server = grpc.server(ThreadPoolExecutor(max_workers=4))
		add_BenchmarkServiceServicer_to_server(BenchmarkImpl(), server)

		server.add_insecure_port(f"{GRPC_ADDRESS}:{GRPC_PORT}")
		server.start()

		server.wait_for_termination()
	except KeyboardInterrupt:
		print("Python gRPC server stopped")
		exit()

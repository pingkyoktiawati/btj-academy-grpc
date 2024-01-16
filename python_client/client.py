import grpc
import calc_pb2
import calc_pb2_grpc
import sys

# gRPC server address
address = 'localhost:50051'

def run():
    # Check if two numbers are provided as command line arguments
    if len(sys.argv) != 3:
        sys.exit("2 numbers expected: n1 n2")

    # Parse command line arguments
    n1, n2 = map(int, sys.argv[1:3])

    # Set up a connection to the gRPC server
    with grpc.insecure_channel(address) as channel:
        # Create a gRPC client
        stub = calc_pb2_grpc.CalculatorStub(channel)

        # Contact the server and print out its response
        try:
            add_response = stub.Add(calc_pb2.AddRequest(n1=n1, n2=n2))
            print(f"{n1} + {n2} = {add_response.r}")

            subtract_response = stub.Subtract(calc_pb2.SubtractRequest(n1=n1, n2=n2))
            print(f"{n1} - {n2} = {subtract_response.r}")

            multiply_response = stub.Multiply(calc_pb2.MultiplyRequest(n1=n1, n2=n2))
            print(f"{n1} * {n2} = {multiply_response.r}")

            divide_response = stub.Divide(calc_pb2.DivideRequest(n1=n1, n2=n2))
            print(f"{n1} / {n2} = {divide_response.r:.2f}")

        except grpc.RpcError as e:
            if "Cannot divide by zero" in str(e):
                print("Dividing error: Cannot divide by zero")
            else:
                print(f"Error: {e.code()}: {e.details()}")

if __name__ == '__main__':
    run()

import grpc
import notif_pb2
import notif_pb2_grpc

def run():
    # Set up a connection to the gRPC server
    channel = grpc.insecure_channel('localhost:50051')
    notif_stub = notif_pb2_grpc.NotifServiceStub(channel)

    print("What topics do you want to get notifications about?")
    print("1. Quote Service")
    print("   This topic will give you random quotes")
    print("2. Image Service")
    print("   This topic will give you random images")
    print("0. Exit")

    # User input to select the topic
    choice = input("Please select the topic: ")

    if choice == '0':
        print("Exit")
        return

    try:
        if choice == '1':
            # Get user input for the number of quotes
            num_quotes = int(input("Enter the number of quotes you want: "))
            if num_quotes < 1:
                print("Invalid number of quotes. Please enter a positive integer.")
                return
            
            # Create a request message for the Quote Service
            request = notif_pb2.QuoteRequest(num_quote=num_quotes)
            
            # Call the gRPC method for the Quote Service
            responses = notif_stub.GetQuote(request)
            for response in responses:
                print(f"Generated Quote: {response.generated_quote}")

        elif choice == '2':
            # Get user input for the number of images
            num_images = int(input("Enter the number of images you want: "))
            if num_images < 1:
                print("Invalid number of images. Please enter a positive integer.")
                return
            
            # Create a request message for the Image Service
            request = notif_pb2.ImagesRequest(num_image=num_images)
            
            # Call the gRPC method for the Image Service
            responses = notif_stub.GenerateImages(request)
            for response in responses:
                print(f"Generated Image: {response.generated_image_urls}")

        else:
            print("Invalid choice. Please enter 1 or 2.")

    except grpc.RpcError as e:
        print(f"RPC error: {e}")

    finally:
        # Close the gRPC channel
        channel.close()

if __name__ == '__main__':
    run()


# Build a Simple gRPC-based Push Notification Application

Membuat project simple push notification dengan Golang di Windows.



## Instalasi Golang

1. **Unduh Golang:**
   - Buka [situs resmi Golang](https://golang.org/dl/) dan unduh versi terbaru.
   - Ikuti petunjuk instalasi yang disediakan.

2. **Konfigurasi:**
   - Setelah instalasi Golang, tentukan direktori tempat proyek Go akan disimpan. 

3. **Verifikasi Instalasi:**
   - Buka terminal dan ketik `go version` untuk memastikan instalasi telah berhasil.

## Instalasi Protocol Buffers
1. **Unduh Protocol Buffers:**
   - Buka [situs resmi Protobuf](https://developers.google.com/protocol-buffers) dan unduh versi terbaru.
   - Ikuti petunjuk instalasi yang disediakan.

2. **Konfigurasi:**
   - Setelah instalasi Protobuf, tentukan direktori tempat proyek Protobuf dengan langkah-langkah [berikut](https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/).

3. **Verifikasi Instalasi:**
   - Buka terminal dan ketik `protoc --version` untuk memastikan instalasi berhasil.

## Setup Project gRPC

1. **Inisialisasi modul Go**:
   ```
   go mod init github.com/pingkyoktiawati/grpc-calculator
   ```
   Nantinya, akan terbuat file `go.mod` secara otomatis.

2. **Instalasi plugin Protocol Buffers (proto):**

   Jalankan command `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28` untuk menghasilkan kode Go dari file proto dan command `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2` untuk menghasilkan kode Go yang mendukung implementasi server dan klien untuk layanan gRPC yang didefinisikan dalam file proto.

3. **Instalasi gRPC-Go package:**

    Jalankan command `go mod tidy` lalu jalankan `go get google.golang.org/grpc` untuk mengunduh paket gRPC beserta dependensinya dari [repositori resmi](https://github.com/grpc/grpc-go) dan menginstalnya ke dalam workspace. Nantinya, akan terbuat file `go.sum` secara otomatis.

4. **Buat file proto**:

   Buat direktori `proto` lalu buat file `notif.proto` dalam direktori tersebut.

5. **Generate kode Go dari file proto:**
   ```
   protoc --go_out=. --go-grpc_out=. calculator/calcpb/calculator.proto
   ```
   Nantinya, akan terbuat file `notif.pb.go` dan `notif_grpc.pb.go` secara otomatis.

6. **Buat file untuk server:**

   Buat direktori `go_server` lalu buat file `main.go` dalam direktori tersebut.

7. **Buat file untuk client:**

   Buat direktori `python_client` lalu instal paket gRPC dan alat-alat yang berkaitan dengan pengembangan gRPC di lingkungan Python dengan command `pip install grpcio grpcio-tools`. Selanjutnya jalankan command berikut untuk generate kode Python dari file proto.

   ```
   python -m grpc_tools.protoc --proto_path=../proto --python_out=. --grpc_python_out=. ../proto/notif.proto
   ```
   Nantinya, akan terbuat file `notif_pb2go` dan `notif_pb2_grpc.py` secara otomatis. Selanjutnya buat file `main.go` dalam direktori tersebut.

   _Note: setup python client menggunakan virtual environment._

## Run Server and Client

Menjalankan Go server dan Python client menggunakan VSCode.

1. Clone repository:

   ```git clone https://github.com/pingkyoktiawati/grpc-calculator.git```

2. Buka folder `grpc-calculator/push_notifications` melalui VSCode

3. Split terminal (dual terminal)
   
   a) Pada terminal pertama, run Go server dengan command 
      
      ```go run main.go```
   
   b) Pada terminal kedua, aktifkan virtual environment terlebih dahulu (`venv/Scripts/activate`) lalu run Python server dengan command 
      
      ```python main.py```


## Test Case

Simple push notification ini akan memiliki 2 topik, yaitu:
1. Quote Service

   Topik ini akan memberikan random quotes sesuai dengan banyaknya quotes yang diinginkan oleh user.

2. Image Service

   Topik ini akan memberikan random URL gambar makanan sesuai dengan banyaknya gambar yang diinginkan oleh user.

Misalkan user memilih topik pertama dan menginginkan satu quotes, maka output pada sisi user akan seperti berikut:

```
Generated Quote: Life doesn't get easier or more forgiving, we get stronger and more resilient. (Steve Maraboli)
```

Lalu misalkan user memilih topik kedua dan menginginkan dua URL gambar makanan, maka output pada sisi user akan seperti berikut:

```
Generated Image: https://foodish-api.com/images/biryani/biryani12.jpg
Generated Image: https://foodish-api.com/images/dosa/dosa78.jpg
```

Sedangkan output pada sisi server akan seperti berikut:

```
2024/01/18 14:38:36 Starting gRPC server on port 50051
2024/01/18 14:38:52 Received request for 1 quotes
2024/01/18 14:39:10 Received request for 2 images
```


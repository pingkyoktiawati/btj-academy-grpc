
# Build a Simple Calculator gRPC Service

Membuat project simple calculator dengan Golang di Windows.



## Instalasi Golang

1. **Unduh Golang:**
   - Buka [situs resmi Golang](https://golang.org/dl/) dan unduh versi terbaru.
   - Ikuti petunjuk instalasi yang disediakan.

2. **Konfigurasi:**
   - Setelah instalasi Golang, tentukan direktori tempat proyek Go akan disimpan. 

3. **Verifikasi Instalasi:**
   - Buka terminal dan ketik `go version` untuk memastikan instalasi telah berhasil.

## Instalasi Protocol Buffers (Protobuf)
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
   Ini akan membuat file `go.mod` secara otomatis.

2. **Buat file `main.go`**:

   Ini akan membuat file `go.sum` secara otomatis setelah menjalankan command `go mod tidy` pada terminal.

3. **Instalasi plugin Protocol Buffers (proto):**

   Jalankan command `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28` untuk menghasilkan kode Go dari file proto dan command `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2` untuk menghasilkan kode Go yang mendukung implementasi server dan klien untuk layanan gRPC yang didefinisikan dalam file proto.

4. **Buat file proto**:

   Buat direktori `calculator/calcpb` lalu buat file `calc.proto` dalam direktori tersebut.

5. **Generate kode Go dari file proto:**
   ```
   protoc --go_out=. --go-grpc_out=. calculator/calcpb/calculator.proto
   ```
   Ini akan membuat file `calc.pb.go` dan `calc_grpc.pb.go` secara otomatis.

6. **Buat file untuk server:**

   Buat direktori `calculator_server` lalu buat file `main.go` dalam direktori tersebut. Lalu jalankan command `go run calculator_server/main.go` untuk menjalankan server.

7. **Buat file untuk client:**

   Buat direktori `calculator_client` lalu buat file `main.go` dalam direktori tersebut. Lalu jalankan command `go run calculator_client/main.go <n1> <n2>` dan ganti `<n1>` dan `<n2>` dengan angka yang ingin dioperasikan.

## Test Case

Simple calculator ini akan menjalankan 4 operasi sekaligus setiap user mengirimkan request, yaitu penjumlahan, pengurangan, perkalian, dan pembagian. 

Misalkan user memasukkan angka 100 dan 3 maka command yang dijalankan adalah `go run calculator_client/main.go 100 3`. Nantinya output pada sisi user akan seperti berikut:
```
2024/01/12 22:30:02 100 + 3 = 103
2024/01/12 22:30:02 100 - 3 = 97
2024/01/12 22:30:02 100 * 3 = 300
2024/01/12 22:30:02 100 / 3 = 33.33
```
Sedangkan output pada sisi server akan seperti berikut:
```
2024/01/12 22:29:50 server listening at [::]:50051
2024/01/12 22:30:02 Received request: 100 + 3
2024/01/12 22:30:02 Received request: 100 - 3
2024/01/12 22:30:02 Received request: 100 * 3
2024/01/12 22:30:02 Received request: 100 / 3
```


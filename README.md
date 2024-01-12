
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

4. **Instalasi plugin Protocol Buffers (proto):**

Jalankan command `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28` untuk menghasilkan kode Go dari file proto dan command `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2` untuk menghasilkan kode Go yang mendukung implementasi server dan klien untuk layanan gRPC yang didefinisikan dalam file proto.

5. **Buat file proto**:

Buat direktori `calculator/calcpb` lalu buat file `calc.proto` dalam direktori tersebut.

6. **Generate kode Go dari file proto:**
```
protoc --go_out=. --go-grpc_out=. calculator/calcpb/calculator.proto
```
Ini akan membuat file `calc.pb.go` dan `calc_grpc.pb.go` secara otomatis.


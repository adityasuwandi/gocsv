# gocsv
A CSV Writer API for [goventory](https://github.com/haruelrovix/goventory)

### Prerequisite
1. Follow `How to Run` section on [goventory](https://github.com/haruelrovix/goventory#how-to-run) first ‼️
1. [Go](https://golang.org/)
```sh
$ go version
go version go1.10 darwin/amd64
```
2. [SQLite](https://www.sqlite.org/index.html)
```sh
$ sqlite3 version
SQLite version 3.16.0 2016-11-04 19:09:39
```

3. [curl](https://github.com/curl/curl)
```sh
$ curl --version
curl 7.51.0 (x86_64-apple-darwin16.0) libcurl/7.51.0 SecureTransport zlib/1.2.8
Protocols: dict file ftp ftps gopher http https imap imaps ldap ldaps pop3 pop3s rtsp smb smbs smtp smtps telnet tftp
Features: AsynchDNS IPv6 Largefile GSS-API Kerberos SPNEGO NTLM NTLM_WB SSL libz UnixSockets
```

### How to Run
1. Clone this repository
```sh
$ git clone https://github.com/haruelrovix/gocsv.git && cd gocsv
```

2. Execute `gocsv.sh`
```sh
$ ./gocsv.sh
```

3. If it asks to accept incoming network connections, allow it.
<img src="https://i.imgur.com/FqfijBf.png" alt="Accept incoming network connections" width="30%" />

4. `gocsv` listening on port 4000
```sh
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /nilaibarang              --> gocsv/api.ExportItemReport (3 handlers)
[GIN-debug] GET    /penjualan                --> gocsv/api.ExportSalesReport (3 handlers)
[GIN-debug] Environment variable PORT="4000"
[GIN-debug] Listening and serving HTTP on :4000
```

### Test the API
Have you heard about [REST Client for VS Code](https://github.com/Huachao/vscode-restclient) ?

1. Laporan Nilai Barang: `http://0.0.0.0:4000/nilaibarang`

<img src="https://user-images.githubusercontent.com/17120764/37029410-ef46bc66-2169-11e8-9c6b-524221ada805.png" title="Laporan Nilai Barang response in CSV format" width=500 />

2. Laporan Penjualan: `http://0.0.0.0:4000/penjualan?startdate=2017-12-01&enddate=2017-12-31`

![image](https://user-images.githubusercontent.com/17120764/37030594-170cfb94-216e-11e8-9917-11a2236425ba.png)

For the default, `Harga` is in integer format. If you want to put `Rp` in front of it, add `rupiah=true` on the request.

`http://0.0.0.0:4000/penjualan?startdate=2017-12-01&enddate=2017-12-31&rupiah=true`

![image](https://user-images.githubusercontent.com/17120764/37030709-95aa8804-216e-11e8-8d2c-30d03c6c86dd.png)

Another optional parameter is `prettifydate`. In Toko Ijah alternate universe, there is no timezone. So it just removes `T` and `Z` from the `timestamp`.

`http://0.0.0.0:8080/penjualan?startdate=2017-12-01&enddate=2017-12-01&prettifydate=true`

<img src="https://user-images.githubusercontent.com/17120764/37024304-02ed24a8-215b-11e8-829d-9b3f7cc153f7.png" width="350" />

### Debugging
VS Code and [Delve](https://github.com/derekparker/delve), a debugger for the Go programming language.

<img src="https://user-images.githubusercontent.com/17120764/37030914-3f38f4a0-216f-11e8-80fe-6a164306ed79.png" alt="Debugging" width=500 />

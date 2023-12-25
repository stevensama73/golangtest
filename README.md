# Backend Golang

## Requirements
1. VSCode
2. Go v1.2.1
3. Mysql

## Step
1. Clone branch main dan socket
2. Untuk yang main pindah ke folder src lalu "go run main.go", kalo yang socket langsung "go run main.go"
3. Buat database dengan nama "dbprovider", kemudian import table dari file "test.sql" yang ada di root folder main
4. Install extensions https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer, kemudian restart vscode, dan kemudian jalankan html yang ada di main di folder src/templates, dengan klik kanan kemudian open with live server
5. Buka page html dengan http://localhost:5500/src/templates/index.html untuk halaman index, http://localhost:5500/src/templates/input.html untuk halaman input, dan  http://localhost:5500/src/templates/output.html untuk halaman output
6. Harus login dulu di index baru bisa ke halaman input atau output, untuk login menggunakan one tap dan akan muncul pop up di sebelah kanan untuk login, jika browser tidak memunculkan pop up gunakan browser lain

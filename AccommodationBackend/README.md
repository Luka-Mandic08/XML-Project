Kreiranje mikroservisa:
-kopiras postojeci mikroservis
-obrises go.mod
-go mod init x_service u tom servisu
-go get common
-na dnu go mod upisati replace common => ../common
-u common folderu napraviti proto fajl za mikroservis i pokrenuti proto komandu
-izmeniti dockerFile i dodati u docker-compose sta treba
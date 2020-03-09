# PDTM
pdtm est un  programme de sauvegarde de fichier à distance sur différent serveur. 
Il sépare un fichier en plusieurs chunk et les sauvegardes en doubles sur 4 serveur, puis envoie l'index a un 5e serveur.
Avant l'envoi, il vas chiffrer votre fichier pour plus de sécurité.
*/!\ Pour le moment PDTM de fonctionne que pour un unique fichier, avec 4 serveur sauvegarde, 1 serveur liste et 1 client.*

## Intallation Programme
Chaque fichier contient les scripts a intaller sur les machines désignées par le nom du fichier

## Installation Client
Votre client doit installer Go pour faire fonctionner le programme

## Installation Serveur
Les serveurs doivent avoir une base de donnée Redis, ainsi que Go installé

## Installation Serveur Liste
Le serveur Liste doit avoir lui aussi un environnement Go pour faire fonctionner le programme

## Configurer les scripts

### Client
Rentrez les ips ainsi que les ports de vos serveurs dans chaque scripts a la places des ips exemples

### Serveur sauvegarde
* Dans Redis.go, si vous avez personalisez l'intallation de votre bdd redis, changer les information dans la première ligne du main
* Dans givedata.go, rentrez l'ip ansi que le port de votre client dans la fonction sendToClient()

### Serveur liste
* Dans savelink.go, si vous avez personalisez l'intallation de votre bdd redis, changer les information dans la première ligne du main
* Dans getlink.go, rentrez l'ip de votre client dans la fonction sendToClient

## Utilisation

### Sauvegarde fichier
* lancez le script redis.go sur les serveur de sauvegarde et savelink.go sur le serveur liste
* lancez le script savefile.go chez le client pour lancer la sauvegarde

### Recupération fichier
* lancez givedata.go sur le serveur sauvegarde
* lancez getfile.go sur le client
* lancez getlink.go sur le serveur liste

  

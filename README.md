GO Project

MONITORING

Le projet a une finalité de monitoring.
Dans cette version, une API va :
- Se connecter sur un serveur via un login (username/password), et récupérer un token d'authentification.
- Une fois la transmission du token réalisée, elle va se connecter sur un dashboard côté client.
- Le dashboard va récupérer les informations (température CPU / GPU) de la machine, et les afficher sur le dashboard.


TODO:
Graphique
Cache ?
Stockage des données ?
SQL LITE



-----------------------------------

Initialisation de GO_API

Mise en place de la libraire :
- go get LIBRAIRIE_NAME
- go mod tidy

Utilisation de WSL + VCXSRV afin de réaliser la compilation, et l'affichage de la fenêtre.


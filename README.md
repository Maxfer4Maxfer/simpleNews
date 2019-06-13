# simpleNews Service
**simpleNews Service** is a simple REST API service written on Go.
**simpleNews Service** consists of two microservice apps.
A first app is called "newssvc". Its give REST APIs to a user for create, view and delete news.
A second app is called "newsrepo". It stores news in a storage. A storage type can be inmem or database. It depends on the settings. By default MySQL database is used.
Interaction between the two services goes through NATS message query. 

**simpleNews Service** can be used as a template or as a project for study purpose.

## Docker-compose Deploy
* [Docker Compose deploy](https://github.com/Maxfer4Maxfer/simpleNews/blob/master/docs/docker-compose-deploy.md)

## Access the service
```bash
curl -X POST http://<simpleNews_IP>/news
```

Show information about a particular news
```bash
curl  -X GET http://<simpleNews_IP>/news/{news_id}
```

Show all news
```bash
curl  -X GET http://<simpleNews_IP>/news
```

Delete all news
```bash
curl  -X DELETE http://<simpleNews_IP>/news/
```

## Donations
 If you want to support this project, please consider donating:
 * PayPal: https://paypal.me/MaxFe
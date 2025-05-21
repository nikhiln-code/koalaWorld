### 📦 Docker Compose Services

The following services are defined in [`docker-compose.yml`](../docker-compose.yml):

```bash
# Start all services
docker-compose up --build
# Start a single service
docker-compose up --build backend
```

## After the starting of the service you can access the following services at the following ports:

Note: The frontend service is comment out from the docker compose for now, still working on it.

| Service    | Port | Description          |
| ---------- | ---- | -------------------- |
| `frontend` | 3000 | Next.js frontend app |
| `backend`  | 8085 | Go Gin backend API   |

You can access the api endpoints defined in [Sample-API's](./api/curl_examples.md)

Go back to [README](./../README.md)

# Mini Chat

## Build and Run the Containers

Open a terminal and type this command below:

```bash
docker-compose up --build
```

This command will:

- Build the chatservice container.
- Pull the latest cassandra and redis images.
- Start all services together in a single network.

## Cassandra table messages

```bash
CREATE TABLE IF NOT EXISTS messages (
    id UUID PRIMARY KEY,
    sender_id UUID,
    recipient_id UUID,
    content TEXT,
    timestamp TIMESTAMP
);
```

## Testing the Application

Once all containers are up and running:

- Open a terminal and use `curl` or a tool like Postman to test the endpoints.
- Test user registration and login using the `/register` and `/login` endpoints.
- Test sending and retrieving messages using the `/send` and `/messages` endpoints.

### Example Testing Commands:

#### Register a User

```bash
curl -X POST -d '{"username": "user1", "password": "password1"}' http://localhost:8080/register
```

#### Login a User

```bash
curl -X POST -d '{"username": "user1", "password": "password1"}' http://localhost:8080/login
```

#### Send a Message

```bash
curl -X POST -d '{"sender_id": "user1_id", "recipient_id": "user2_id", "content": "Hello!"}' http://localhost:8080/send
```

#### Retrieve Messages

```bash
curl -X GET 'http://localhost:8080/messages?sender_id=user1_id&recipient_id=user2_id'
```

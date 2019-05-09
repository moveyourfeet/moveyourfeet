# Mobile BFF

The purpose of this service is to:

- be the single point of connection for the mobile app

## Debug

After starting it with  `cd ./mobileBFF && make dev` it will be on `http://localhost:8088/`

Use this online socket.io tester: 
[socketio-client-tool](http://amritb.github.io/socketio-client-tool/#url=aHR0cDovL2xvY2FsaG9zdDo4MDg4Lw==&opt=&events=)

Emit message to topic: `create-game`

```json
{
    "name":"my New Game",
    "owner":"user1",
    "playingTime":"30",
    "gameField": {
        "latitude": 9.32,
        "longitude": 56.11,
        "radius": 10
    }
}
```

## Endpoints

| Env           | URL                                                                |
|---------------|--------------------------------------------------------------------|
| Local dev dns | [http://mobilebff.localtest.me](http://mobilebff.localtest.me)     |
| Local ip      | [http://localhost:8003/](http://localhost:8003/)                   |


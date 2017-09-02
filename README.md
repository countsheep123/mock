### Mock

#### Install

```
go get -u github.com/countsheep123/mock/...
```

#### Usage

```
mock -config config.json
```

#### Sample config

```
[
  {
    "port": 8080,
    "endpoints": [
      {
        "endpoint": "/apps",
        "methods": [
          {
            "method": "POST",
            "status": 200,
            "response": {
              "id": "0000000002"
            }
          },
          {
            "method": "GET",
            "status": 200,
            "response": [
              {
                "id": "0000000001",
                "name": "game"
              },
              {
                "id": "0000000002",
                "name": "social"
              },
              {
                "id": "0000000003",
                "name": "entertainment"
              }
            ]
          }
        ]
      }
    ]
  }
]
```

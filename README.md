# go-go-power-rangers
A simple golang API whose purpose we will discover later...

```bash
# POST
curl -X POST -H "Content-Type: application/json" -d '{
  "title": "Power Ranger",
  "description": "blabla"
  }' http://localhost:8080/seasons

# GET
curl http://localhost:8080/seasons
```
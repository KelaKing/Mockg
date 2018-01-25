# Mockg

--
Mock JSON responses


### How to use

```
go get github.com/kelaking/mockg
mkdir temp
echo "{ "msg": "success" }" > ./temp/foo.json
mockg ./temp/
```
try open `http://localhost:8080/foo.json`

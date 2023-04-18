# Collector


## to upload a file
```
curl --location 'http://localhost:5000/api/upload' \
--header 'Content-Type: multipart/form-data' \
--form 'coverageReport=@"File with path to be uploaded"'
```


## to download
```
curl -o <Archive Name>.zip --location 'http://localhost:5000/api/download'
```
# url_shorting


```
mkdir uploads - Create folder for qrcode images

make container - Create containers
```

`make migration - up migration up`

`make run - Go run project`


http://localhost:8080/swagger/index.html#/URL/post_api_v1_url swagger file

server - http://13.234.59.34:8000/swagger/index.html


## Create URL 

```
	
Url

Example Value
Model
{
  "custom_url": "custom_url",
  "exp_count": "10", (ixtiyoriy, yozmasa cheksiz)
  "exp_time": "2006-01-02 15:04:05", (ixtiyoriy, yozmasa cheksiz)
  "org_path": "https://www.google.com/" (https://url/)
}

```
## Response
```
	
Response body
Download
{
  "id": "05c49a15-5f78-4cbf-a1a0-eee47a745ea8",
  "user_id": "05734a3a-5621-451e-ad3d-7ea6f0e9836c",
  "org_path": "https://url/",
  "short_path": "Ns6UXjcW",
  "counter": 0,
  "status": 0,
  "qr_code_path": "uploads/43a/378380fd-5b1e-44dd-b53f-1ea0c298db75/40339b96-e70b-4221-8e8b-583db03d97ef.jpg",
  "created_at": 1677437483,
  "updated_at": 0
}
```
## Request
```
http://localhost:8080/Ns6UXjcW
```

## QRCOde
```
http://localhost:8080/uploads/43a/378380fd-5b1e-44dd-b53f-1ea0c298db75/40339b96-e70b-4221-8e8b-583db03d97ef.jpg
```

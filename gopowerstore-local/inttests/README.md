# GoPowerStore integration tests

You should export all ENV variables required for API client initialization before run this tests.

It's also possible to create GOPOWERSTORE_TEST.env which will hold all required vars

_GOPOWERSTORE_TEST.env example:_
```shell
GOPOWERSTORE_INSECURE=true
GOPOWERSTORE_HTTP_TIMEOUT=60
GOPOWERSTORE_APIURL=https://127.0.0.1/api/rest
GOPOWERSTORE_USERNAME=admin
GOPOWERSTORE_PASSWORD=Password
GOPOWERSTORE_DEBUG=true
```
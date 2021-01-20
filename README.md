# NgProxy

Very basic (slightly inefficient) reverse proxy for Angular app that's hosted on Google Cloud Storage

If a path doesn't have a . in it, then a HEAD request is made to check if the path is not a 404.
If the path is a 404 the proxy returns the index.html of the proxy URL (to allow angular routing to work)

## Usage
To build install go and then run
```go build```

To run
```./ngproxy -port 3000 -target https://storage.googleapis.com/parenting-app-ui-master1```

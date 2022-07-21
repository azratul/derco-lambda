# Lambda with Go & Gin Template

Single-endpoint microservice transformed into a "lambda function"(FaaS in AWS)


## To test in local


1. Load env vars (DO NOT USE "source" command, because it's not POSIX compliant!)


```
. .env
```


2. Run the service


```
go run .
```


3. That's all folks!



## To test in AWS


1. Nothing, just push to the required branch (how to push?!!! really?!!!)


```
git checkout <REQUIRED_BRANCH>
```


```
git commit -am "my first push :)"
```


```
git push
```

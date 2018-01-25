# go-web-framework

## Introduction
- A simple web framework using [httprouter](https://github.com/julienschmidt/httprouter) and net/http
- Using Object-Oriented
- You can add middlewares in `PatchHandle`
- For your new Handler just implement need http method, other methods will return `StatusMethodNotAllowed: 405`, the router will help you to accomplish the methods
- For usage see `src/example`

## Why
While using tornado or django, we just need to implement the methods we need, and one path support all http methods.
But in go we need to add different handle func for each http method by ourselves. This simple frame work can help to let you just need to implement the methods you need.
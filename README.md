# go-web-framework

## Introduction
- A simple web framework using [httprouter](https://github.com/julienschmidt/httprouter) and net/http
- Using Object-Oriented
- You can add middlewares in `PatchHandle`
- For your new Handler just implement need http method, other methods will return `StatusMethodNotAllowed: 405`
- For usage see `src/example`

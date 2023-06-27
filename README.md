# **_RESTful_Blog_**

# solution :

I have used the monolithic design in this project and tried to implement it in a very simple manner. We have two controllers named "Post" and "Category" that utilize their respective services. Each service uses the required repository. Additionally, there is a middleware called "Auth" that hard codes the access check for logging in, and all APIs are versioned.

* If you don't have Golang installed on your system, you need to install Golang first.

    1- Visit the official Golang website: https://golang.org/dl/

  2 - On the downloads page, you will see different versions of Golang available for various operating systems. Choose the appropriate version for your operating system. Golang supports Windows, macOS, and Linux.

Run Make command to start easy project :

To download go library.
```makefile
    make go-mod  
```
Run Test.

```makefile
    make test
```

To run project:
```makefile
    make run 
```

Run project with docker.

```makefile
    make docker
```


you can see the swagger in local:
```swagger
http://localhost:<port-serve>/swagger/index.html
```
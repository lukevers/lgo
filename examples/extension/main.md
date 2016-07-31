This is an example where our file ends in the extension `.md` instead of `.lgo`.

As always, let's first set our package name before we do anything else. Since we want to build and run this package instead of using it as a package, we're going to set our package to `main`.

```
package main
```

We're importing the `log` package because in this example all we're going to do is create a basic logger and print something out. At least it's a little more exciting than hello world. The `os` package is also being imported here because we know that we're going to be printing out with the logger to the OS.

```
import (
    "log"
    "os"
)
```

Open our main function:

```
func main() {
```

Now now we can reate our logger. This logger should print to [stdout](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.2), have the prefix `"lgo: "`, and use the [default](https://godoc.org/log#pkg-constants) flags for a `*Logger`.

```
l := log.New(os.Stdout, "lgo: ", log.LstdFlags)
```

And we'll just use it and print out something fun:

```
l.Println("Kittens kittens everywhere!!")
```

And lastly let's just close our main function

```
}
```

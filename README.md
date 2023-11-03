# raft a golang Single Page Application powered through wasm

## install

```sh
    go install https://github.com/AlexTheProgrammer/raft
```

## run

```sh
    raft dev # start development server

    raft build # bundle site assets
```

# philosophy & decisions
In a world of very sophiscated frameworks and libraries for web development
raft offers the opposite approach.

raft's code base will be kept as simple as possible so that any developer can fork it and extend with minimal effort.


## configuration
do I want another xxx.config.ts file that requires me to read some article to understand how it works.
    ....
No.

so configuration is not supported, if you want to configure raft, fork it. Or better don't use it at all and build your own tool.

## features that I actually want
- build (make an artifact that I can host on any cdn)
- dev (run a local server that watches for changes)


do I want a bunch of additional features that cater to edge use cases to make the product more self sufficient.
    ... 
No.

these are the only features I generally use in any web framework/library until I find that the framework holds me back.
I propose a different approach to adding more features. 
If/when I find that raft holds me back. Instead of adding features I'm going to enable piping of rafts output to
another tool.

then it becomes not raft's problem. I believe this is in line with the Unix philosophy of minimal modular software.

## testing 
but, what about unit testing?

yep, even unit testing, not supported, if you want that you're gonna have to figure it out for yourself.

but, my view on unit testing is that you shouldn't be testing if a div contains a certain value.
you should be testing that your logic that generates that value works, because that's your code. 
the div rendering that value correctly, that's raft's / the html parser's (i.e. the browser's) job
and because raft is a go module and your site is just a go program you can use go's inbuilt testing framework 
like you would for any go project to test your code.




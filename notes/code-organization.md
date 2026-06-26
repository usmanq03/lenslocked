# Code Organization

- Some codebases are easy to navigate
- Others are hard to follow
- Code organization is often the difference

A good code structure will:
- Make it easier to guess where bugs are
- Make it easier to add new features
- And more

## Flat Structure

All code is in a single package

Files can still seperate code:

```
myapp/
    gallery_store.go
    gallery_handler.go
    gallery_templates.go
    user_store.go
    user_handler.go
    user_templates.go
    router.go
    ....
```

## Seperation of concerns

Seperate code based on duties

HTML & CSS is a traditional example:
- HTML focuses on overall structure
- CSS focuses on styling it

Model-View-Controller (MVC) is a popular structure following this strategy. In a web app:
- `models` => data, logic, rules; usually databases
- `views` => rendering things; usually html
- `controllers` => connects it all. accepts user input, passes to models to do stuff, then passes data to views to render things; usually handlers

```bash
myapp/
    controller/
        user_handler.go
        gallery_handler.go
        ---
    views/
        user_templates.go
        ...
    models/
        user_store.go
        ...
```

Doesnt need to be named `models` `views` and `controllers`.

[Buffalo] (https://gobufallo.io/en/) uses a variation of MVC, but deosn't name directories exactly

Very common in frameworks b/c it is predictable. Makes it easier for a dev to guess where certain files will be

Ruby on Rails, Django, etc use MVC.


## Dependency-based structure

Ben Johnson writes about this at [Go Beyond] (https://www.gobeyond.dev/standard-package-layout)

Structured based on dependencies, but with a common set of interfaces and types they all work with.

```bash
myapp/
    user.go #defines a User type that isnt locked into any specific data store
    user_store.go # defines a UserStore interface

    psql/
        user_store.go # implement the UserStore interface with a Postgresql specific implementation
```

## Many more

- Domain driven design
- Onion architecture


## Which should you use?

No such thing as one-size fits all.
Each has pros & cons.

Flat can work well with microservices since each is small.

Flat won't work well at Google with a massive monorepo.


Experience also matters:
- Globals are bad, but why?
- Testable code only makes sense after you try to test untestable code.



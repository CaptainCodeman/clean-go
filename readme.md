# Clean Architecture in Go

There are a number of different application architectures that are all simlar variations on the same theme which is to have clean separation
of concerns and dependecies that follow the best practices of "the dependency invesion principle":

A. High-level modules should not depend on low-level modules. Both should depend on abstractions.
B. Abstractions should not depend on details. Details should depend on abstractions.

Variations include:

* [The Clean Architecture](https://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html) advocated by Robert Martin ('Uncle Bob')
* Ports & Adapters or [Hexagonal Architecture](http://alistair.cockburn.us/Hexagonal+architecture) by Alistair Cockburn
* [Onion Architecture](http://jeffreypalermo.com/blog/the-onion-architecture-part-1/) by Jeffrey Palermo

From more in-depth practical application of many of the ideas I can strongly recommend the excellent book [Implementing Domain-Driven Design]
(http://www.amazon.com/Implementing-Domain-Driven-Design-Vaughn-Vernon/dp/0321834577) by Vaughn Vernon that goes into far greater detail.

Besides the clean codebase the approaches also bring other advantages - significant parts of the system can be unit tested quickly and easily
without having to fire up the full web stack, something that is often difficult when the dependencies go the wrong way (if you need a database
and a web-server running to make your tests work, you're doing it wrong).

I'd used it before in the world of .NET but forgot about it after moving to coding more in Python. After switching languages again (yeah, right)
to the wonderful world of go I came across a blog post which re-ignited my interest in it:
[Applying The Clean Architecture to Go applications](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)

It's a great read but I found the example a little overly-complex with the database code especially being a significant distraction and at the
same time it was light on some issues I wanted to resolve such as switching between different storage typea and web UI.

I've also been looking for a way to make my application usable both standalone and on AppEngine as well as being easier to test so this
seemed like a good opportunity to do some experimenting and this is what I came up with.

## Dependency Rings

We've all heard of n-tier or layered architecture, especially if youve' come from the world of .NET or Java and it's unfair that it get's a bad
rap. Probably because it was often so poorly implemented with the typical mistake of everything relying on the database layer at the bottom which
made software rigid, difficult to test and closely tied to the vendors database implementation (hardly surprising that they promoted it so hard).

Reversing the dependencies though has a wonderful transformative effect on your code. Here is my interpretation of the layers or rings implemented
using the Go language (Golang for Google).

### Domain

At the center of the dependencies is the domain. These are the business objects or entities and should represent and encapsulate the fundamental
business rules such as "can a closed customer account create a new order?". There is usually a single root object that represents the system and
which has the factory methods to create other objects (which in turn may have their own methods to create others). This is where the domain-driven
design lives.

Looking at this should give you an idea of the business model for the application and what the system is working with. This package allows code
such as unit tests to excercise the core parts of the app for testing to ensure that basic rules are enforced.

### Engine / Use-Cases

The application level rules and use-cases orchestrate the domain model and add richer rules and logic including persistence. I prefer the term
engine for this package becase it is the engine of what the app actually does. The rules implemented at this level should not affect the domain
model rules but obviously may depend on them. The rules of the application also shouldn't rely on the UI or the persistence frameworks being used.

Why would the business rules change depending on what UI framework is the new flavour of the month or if we want to change from an RDBMS to MongoDB
or some cloud datastore? Those are implementation details that pull the levers of the use cases or are used by the engine via abstract interfaces.

### Interface Adapters

These are concerned with converting data from a form that the use-cases handle to whatever the external framework and drivers use. A use-case
may expect a request struct with a set of parameters and return a response struct with the results. The public facing part of the app is more
likely to expect to send requests as JSON or http form posts and return JSON or rendered HTML. The database may return results in a structure
that needs to be adapted into something the rest of the app can understand.

### Frameworks and Drivers

These are the ports that allow the system to talk to 'outside things' which could be a database for persistence or a web server for the UI.
None of the inner use cases or domain entities should know about the implementation of these layers and they may change over time because ...
well, we used to store data in SQL, then MongoDB and now cloud datastores. Changing the storage should not change the application or any of
the business rules. I tend to call these "providers".

# Run

## App Engine

    goapp serve

## Standalone

    mongod --config /usr/local/etc/mongod.conf
    go run app.go

## Run Tests

    ginkgo watch -cover domain

    go tool cover -html=domain/domain.coverprofile
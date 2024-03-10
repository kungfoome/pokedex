# What is this?

This is a project from the boot.dev golang series. This is a REPL built in golang that uses the PokeAPI to gather information. We created a pokedex to store pokemon and it's information in. We can look up locations, catch pokemon, store pokemon in a pokedex, and view information about pokemon that were caught. 

# Goals
- Learn how to parse JSON in Go
- Practice making HTTP requests in Go
- Learn how to build a CLI tool that makes interacting with a back-end server easier
- Get hands-on practice with local Go development and tooling
- Learn about caching and how to use it to improve performance

# My Thoughts on The Project

Overall I thought it offered a good experience into creating a CLI and interacting with JSON using Golang. I felt like at times it was a bit difficult to figure out how to get something started and what the best approach was. This lead me into creating something that works, but could be much better. I did think that the topics that were explored were great topics to know in Golang and it sort of touched a bit on some key topics.

# What Would I Improve

- I would improve the overall structure of the code. It is a bit messy right now and this means there is a lot of duplicate code where there should not be.
    - Using a context to pass information about the api connection.
    - Being able to pass certain parameters to a function and know what they are. Right now it is just a slice of parameters.
    - Having a config, but don't store context information in there and only store them in a context in which it relates to the api itself. For example, the only api that needs to know about next map and previous map is the map command. Other commands do not need that context.
- I would dive more into types and how I can best use them. I had to use a work around so I could create a function where a type is a reciever. There is probably a better and simplier way to accomplish this, as I just want to print some information.
- I would do better error handling. Right now I just default to a Fatal error if the response is 300 or greater. For instance, if you get a 404, the program should not throw a fatal error and simply just say something does not exist.
- It would be nice for the CLI to support the "up" and "down" arrow to cycle through commands.
- I would like to add some more unit tests and get some more practice there.
- Create a pokedex that is persistent and saved to disk so it can be reloaded.
- Improve the catch algorithm of pokemon, right it is pretty simplistic. I could always expand on this and add more functionality like catching the same type of pokemon and also catching pokemon of different levels.

I would also like to try doing this Charm after adding a few features myself around terminal interaction. It would be nice to see the differences there and what Charm offers out of the box for you. We might be able to create a mini pokemon game in the terminal with this approach as well. We could explore a map interactively just using ASCII.
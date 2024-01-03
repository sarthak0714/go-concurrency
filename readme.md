# Concurrency 

## Explaining with an example
* So imagine a game server which needs player data
* The data comes in multiple Fetch functions fetching player info,history,friends etc.
* In an Syncronous environment the data would be fetched in time = total execution time of all fetch
* But in an async environment the total time = execution time of the slowest function which will be way faster than a sync env
* goroutines assign func to a thread
* which ensure parellel execution 

## time without conc

``` bash
361.555865ms
```
## time with conc

``` bash
180.255046ms
```

### Hence concurrency helps xd 
Thanks Anthony!

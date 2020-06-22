# [Golang] [Beginner] Quiz Application
[Thanks Notion!](https://www.notion.so/Golang-Beginner-Quiz-Application-816cc1670b8f4d8cb6205e9383bca969)
Created: Apr 07, 2020 1:25 AM

# Introduction

A simple application to get to learn golang a little more! This was my first time using Golang, and wow is it so different than what I'm used to (Javascript/React)!

The only similarities I've seen so far are closures. Everything else has been quite different. From the way we initialize variables, leave out the curly braces in if/else, and the for statements. The switch statement is now a 'select' statement now. Importing packages and setting up the small application was also quite different from javascript. Geez we don't really have classes either! We use structs and slices. There are no more convenient functional programming tools like javascript's .map, or .filter anymore.  Perhaps most importantly, Golang allows us to use go routines and channels. These are all ideas I am still trying to wrap my head around! 

I wanted to learn more about golang because I wanted to challenge myself to learn something new and to do something more "backend," to see how I like it!

# TIL

## Concurrency

## Goroutines

Quoted basically everywhere, goroutines are essentially (they are like, but not exactly) a "lightweight thread managed by the Go runtime." Less complex than creating an actual thread. Usually accompanied by the word "go," closures, and are immediately invoked. (IIFE from js much?) 

Threads by themself having some factors that make them slower. These are the three points in which threads can be costly:

1. large stack size (> 1MB) 
2. restore the register, can be a hit to performance
3. threads setup and teardown require communication with OS, which is slow

Luckily, Golang's go-routines only exist in the virtual space of Go's runtime. This helps combat point number 3! We thus rely on the Go runtime to manage their lifecycles. Go routines also only allocate 2KB of data when starting out, and can dynamically allocate more or less as needed. This is significantly less than the 1MB data allocation when starting out for threads. This solves problem 1! Go routines also only update three registers at most, compared to all of the registers when using threads. This solves problem 2!

Here's a cute picture that illustrates how goroutines communicate with Go's runtime and the OS:

![Golang%20Beginner%20Quiz%20Application%2014b6cb33fa6b4842bd0d8e79b89cd67d/Untitled.png](Golang%20Beginner%20Quiz%20Application%2014b6cb33fa6b4842bd0d8e79b89cd67d/Untitled.png)

A great tool for dealing with concurrency.

### Sources

- Straight from the source: [https://tour.golang.org/concurrency/1](https://tour.golang.org/concurrency/1)
- Go routines in more detail: [https://medium.com/the-polyglot-programmer/what-are-goroutines-and-how-do-they-actually-work-f2a734f6f991](https://medium.com/the-polyglot-programmer/what-are-goroutines-and-how-do-they-actually-work-f2a734f6f991)

## Channels

Channels provide a way for two goroutines to communicate with one another. 

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

### Sources

- [https://www.golang-book.com/books/intro/10](https://www.golang-book.com/books/intro/10)

## Structs

These are my objects. Importing objects was slightly difficult at first. In the end, because this application was so simple, I decided to use structs in the same file, instead of separating the code. 

## Some new packages

Using the packages provided by Golang just hammered home that I am truly a beginner at Golang! It was great to start with this simple application, because it introduced to me some very basic packages. This will save me time for hunting packages in the future! 

# Resources For More Go!

**Rob Pike - 'Concurrency Is Not Parallelism'**

[https://www.youtube.com/watch?v=cN_DpYBzKso&t=441s](https://www.youtube.com/watch?v=cN_DpYBzKso&t=441s)

**Analysis of a Go runtime scheduler**
[http://www1.cs.columbia.edu/~aho/cs6998/reports/12-12-11_DeshpandeSponslerWeiss_GO.pdf](http://www1.cs.columbia.edu/~aho/cs6998/reports/12-12-11_DeshpandeSponslerWeiss_GO.pdf)

**Five things that make Go Fast**

[https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast](https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast)

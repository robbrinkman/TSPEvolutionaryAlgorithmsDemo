Evolutionary Algorithms Traveling Salesman Problem
==========

This is the Evolutionairy Algorithm Demo for the Traveling Salesman Problem that I use to demo the simplicity and power of Evolutionairy Algoritms.

To get it, just clone this repository by using your favorite git gui or by executing 

```
git clone https://github.com/bknopper/TSPEvolutionaryAlgorithmsDemo.git
```

### Running the Demo - Backend
Because this demo uses [Gradle](https://gradle.org/) as package manager / build automation system and has the Gradle wrapper incorporated you can run the backend of this demo by simply executing the gradle wrapper.

For example, from the TSPEADemo directory, run:

```bash
./gradlew run
```
This should start resolving all the dependency and start the backend of the demo for you using Spring Boot.

### Running the Frontend
Since the frontend is based on angular, we need some form of http serving. My favorite is the python SimpleHttpServer and you can let it start serving by running the following from the `TSPEADemo/frontend/app` dir:

```bash
python -m SimpleHTTPServer 8001
```

That's it! Now you can navigate to `http://localhost:8001` to start playing. 


Have fun!

For the talk I did on J-Fall where this demo is used, see: [Evolutionairy Algorithms J-Fall talk](http://www.youtube.com/watch?v=5LUqjnwbp5c)

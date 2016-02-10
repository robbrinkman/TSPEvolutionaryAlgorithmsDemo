angular.module('myApp').directive('demo', ['$routeParams', function ($routeParams) {
    return {
        restrict: 'E',
        templateUrl: 'components/demo/demo.html',

        scope: {},
        controller: function ($scope, $log, $rootScope, $location, $http, $timeout) {

            $scope.backendErrors = false;

            var container = $("#flot-moving-line-chart");

            var maximumPoints = 300;

            $scope.fitnessPlot = {};
            $scope.fitnessPlot.ys = [];
            for (var i = 0; i < maximumPoints; i++) {
                $scope.fitnessPlot.ys.push(0);
            }

            // zip the generated y values with the x values

            var res = [];
            for (var i = 0; i < $scope.fitnessPlot.ys.length; i++) {
                res.push([i, $scope.fitnessPlot.ys[i]])
            }

            $scope.fitnessPlot.series = [{
                data: res,
                lines: {
                    fill: true
                }
            }];

            //

            $scope.fitnessPlot.plot = $.plot(container, $scope.fitnessPlot.series, {
                grid: {
                    borderWidth: 1,
                    minBorderMargin: 20,
                    labelMargin: 10,
                    backgroundColor: {
                        colors: ["#fff", "#e4f4f4"]
                    },
                    margin: {
                        top: 8,
                        bottom: 20,
                        left: 20
                    },
                    markings: function (axes) {
                        var markings = [];
                        var xaxis = axes.xaxis;
                        for (var x = Math.floor(xaxis.min); x < xaxis.max; x += xaxis.tickSize * 2) {
                            markings.push({
                                xaxis: {
                                    from: x,
                                    to: x + xaxis.tickSize
                                },
                                color: "rgba(232, 232, 255, 0.2)"
                            });
                        }
                        return markings;
                    }
                },
                xaxis: {
                    tickFormatter: function () {
                        return "";
                    }
                },
                yaxis: {
                    min: 0,
                    max: 25000
                },
                legend: {
                    show: true
                }
            });

            $scope.algorithmStillRunning = false;
            $scope.settings = {};
            $scope.settings.selectedOption = {
                id: 1,
                name: "low",
                settings: {
                    "mutationProbability": 25,
                    "populationSize": 500,
                    "nrOfGenerations": 5000,
                    "fitnessThreshold": 11000,
                    "parentSelectionSize": 60,
                    "parentPoolSize": 100
                }
            };
            $scope.settings.availableOptions = [
                {
                    id: 0,
                    name: "none",
                    settings: $scope.settings.selectedOption.settings
                },
                {
                    id: 1,
                    name: "low",
                    settings: {
                        "mutationProbability": 25,
                        "populationSize": 500,
                        "nrOfGenerations": 5000,
                        "fitnessThreshold": 11000,
                        "parentSelectionSize": 60,
                        "parentPoolSize": 100
                    }
                },
                {
                    id: 2,
                    name: "medium",
                    settings: {
                        "mutationProbability": 50,
                        "populationSize": 75,
                        "nrOfGenerations": 10000,
                        "fitnessThreshold": 11000,
                        "parentSelectionSize": 30,
                        "parentPoolSize": 50
                    }
                },
                {
                    id: 3,
                    name: "high",
                    settings: {
                        "mutationProbability": 90,
                        "populationSize": 1000,
                        "nrOfGenerations": 20000,
                        "fitnessThreshold": 11000,
                        "parentSelectionSize": 6,
                        "parentPoolSize": 10
                    }
                }
            ];
            $scope.setPresetToNone = function () {
                console.log("change!");
                $scope.settings.selectedOption = {
                    id: 0,
                    name: "none",
                    settings: $scope.settings.selectedOption.settings
                }
            },

                $scope.run = {"generation": 0, "bestFitness": 0};

            function setNewValues(response) {
                $scope.areas = response.data.route;
                $scope.run.bestFitness = response.data.fitness;
                $scope.run.generation = response.data.generation;

                $scope.fitnessPlot.ys = $scope.fitnessPlot.ys.slice(1);
                $scope.fitnessPlot.ys.push($scope.run.bestFitness);

                // zip the generated y values with the x values

                var res = [];
                for (var i = 0; i < $scope.fitnessPlot.ys.length; i++) {
                    res.push([i, $scope.fitnessPlot.ys[i]])
                }

                $scope.fitnessPlot.series[0].data = res;
                $scope.fitnessPlot.plot.setData($scope.fitnessPlot.series);
                $scope.fitnessPlot.plot.draw();
            }

            $scope.getBest = function () {
                return $http.get('http://localhost:8080/api/currentBest').then(function (response) {
                    setNewValues(response);
                }, function (response) {
                    $scope.algorithmStillRunning = false;

                    // try to get the last rating
                    $http.get('http://localhost:8080/api/latestBest').then(function (response) {
                        setNewValues(response);
                    });
                });
            };

            $scope.start = function () {
                var body = $scope.settings.selectedOption.settings;

                $scope.startAlgorithm(body).then(function () {
                    $scope.algorithmStillRunning = true;
                    $scope.updateCurrentBest();
                });

            };

            $scope.stop = function () {
                $http.post('http://localhost:8080/api/stopAlgorithm').
                    then(function (response) {
                        // we're good!
                    });
            }

            $scope.updateCurrentBest = function () {
                if ($scope.algorithmStillRunning) {
                    $scope.getBest().then(function (response) {
                        var route = $scope.updateRoute();

                        if ($scope.flightPath) {
                            $scope.flightPath.setMap(null);
                        }

                        $scope.flightPath = new google.maps.Polyline({
                            path: route,
                            geodesic: true,
                            strokeColor: '#FFFF66',
                            strokeOpacity: 1.0,
                            strokeWeight: 2
                        });

                        $scope.flightPath.setMap($scope.map);

                        $timeout($scope.updateCurrentBest, 250);
                    });
                }
            };

            $scope.updateRoute = function () {
                var route = [];
                for (var j = 0; j < $scope.areas.length; j++) {

                    var city = $scope.areas[j];

                    city.location = new google.maps.LatLng(city.latitude, city.longitude);
                    var marker = new google.maps.Marker({
                        position: city.location,
                        map: $scope.map,
                        title: city.name,

                    });

                    route.push({lat: city.latitude, lng: city.longitude})
                }
                return route;
            }

            $scope.startAlgorithm = function (body) {
                return $http.post('http://localhost:8080/api/startAlgorithm', body).
                    then(function (response) {
                        $scope.backendErrors = false;
                        return response;
                    }, function (response) {
                        if (response.status === -1) {
                            $scope.backendErrors = true;
                        }
                    });
            }

            $scope.stillRunning = function () {
                return $http.get('http://localhost:8080/api/stillRunning').then(function (response) {
                    return response.data;
                });
            };

            $scope.initialize = function () {
                $scope.backendErrors = false;
                if ($scope.algorithmStillRunning) {
                    $scope.stopAlgorithm;
                }
                $scope.algorithmStillRunning = false;
                $scope.run = {"generation": 0, "bestFitness": 0};
                var mapOptions = {
                    center: new google.maps.LatLng(50.521401, 9.623885),
                    zoom: 5,
                    mapTypeId: google.maps.MapTypeId.SATELLITE
                };
                $scope.map = new google.maps.Map(document.getElementById("map-canvas"),
                    mapOptions);

                $scope.fitnessPlot.ys = [];
                for (var i = 0; i < maximumPoints; i++) {
                    $scope.fitnessPlot.ys.push(0);
                }

                // zip the generated y values with the x values

                var res = [];
                for (var i = 0; i < $scope.fitnessPlot.ys.length; i++) {
                    res.push([i, $scope.fitnessPlot.ys[i]])
                }
                $scope.fitnessPlot.series[0].data = res;
                $scope.fitnessPlot.plot.setData($scope.fitnessPlot.series);
                $scope.fitnessPlot.plot.draw();

                return $http.get('http://localhost:8080/api/getCities').then(function (response) {
                    $scope.areas = response.data;
                    $scope.updateRoute();
                });

            };
            google.maps.event.addDomListener(window, 'load', $scope.initialize);
            $scope.initialize();
        }
    };
}]);
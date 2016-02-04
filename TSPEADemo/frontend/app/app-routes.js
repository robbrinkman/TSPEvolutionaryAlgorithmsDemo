var appModule = angular.module('myApp');

appModule.config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.
            when('/home', {
                templateUrl: 'partials/demo.html'
            }).
            otherwise({
                redirectTo: '/home'
            });
    }
]);

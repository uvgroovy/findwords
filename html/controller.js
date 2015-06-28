var phonecatApp = angular.module('findwordsApp', []);

phonecatApp.controller('WordListController', function ($scope, $http) {
  $scope.letters = "";
  $scope.words = [];
  $scope.errormsg = "";

  $scope.$watch("letters", function(newValue, oldValue) {

    config = {"params" : {"letters":newValue}}
    $http.get('/words', config).
      success(function(data, status, headers, config) {
        $scope.words = data
      }).
      error(function(data, status, headers, config) {
        $scope.words = []
        $scope.errormsg = data
      });
    });


});

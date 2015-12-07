angular.module("ChristmasTracker")

.controller("newPersonCtrl", function($rootScope, $scope, $location, $routeParams, person, currentPerson) {
	$scope.newPerson = new person();
	
	$scope.refreshCurrentPerson = function() {
		refreshPerson = person.current();
		
		return refreshPerson.$promise.then( function() {
			angular.copy( refreshPerson, currentPerson );
			$rootScope.currentPerson = currentPerson;
		});
	}
	
	$scope.createUser = function() {
		$scope.saving = true;
		
		return $scope.newPerson.$save().then( function() {
			return $scope.newPerson.$link();
		}).then( function() {
			return $scope.refreshCurrentPerson();
		}).then( function () {
			$location.path("/");
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	}
	
	$scope.linkUser = function( link ) {
		$scope.newPerson.Name = link;
		$scope.saving = true;
		
		return $scope.newPerson.$link().then( function() {
			return $scope.refreshCurrentPerson();
		}).then( function () {
			$location.path("/");
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	}
	
	$scope.link = $routeParams["link"];
})
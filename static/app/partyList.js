angular.module("ChristmasTracker")
.controller("partyListCtrl", function($rootScope, $scope, party) {
	$scope.getParties = function() {
		$scope.parties = party.query()
		$scope.parties.$promise.catch(function(reason) {
			$rootScope.errorMessage = reason;
		})
		
		return $scope.parties.$promise;
	};
	
	$scope.createParty = function() {
		return new party();
	};
	
	$scope.saveParty = function( newParty ) {
		newParty.$save().then( function() {
			$scope.getParties();
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	};
	
	$scope.getParties();
})
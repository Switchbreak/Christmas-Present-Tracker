angular.module("ChristmasTracker")
.controller("partyCtrl", function($rootScope, $scope, $routeParams, party, invitedPeople, comment) {
	$scope.getParty = function() {
		$scope.party = party.get( { title: $routeParams.title } );
		$scope.party.$promise.catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	}
	
	$scope.getInvitedPeople = function() {
		$scope.invitedPeople = invitedPeople.query({ title: $routeParams.title })
		$scope.invitedPeople.$promise.catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	};
	
	$scope.init = function() {
		$scope.partyTitle = $routeParams.title;
		
		$scope.getParty();
		$scope.getInvitedPeople();
	}
	
	$scope.invite = function() {
		if (!$scope.people)
			$scope.getPeople();
		
		$scope.invitedPerson = null;
	};
	
	$scope.invitePerson = function( person ) {
		party.$invite( { person: person.Name } ).then( function() {
			$scope.getInvitedPeople();
		}).catch(function( reason ) {
			$rootScope.errorMessage = reason;
		});
	};
	
	$scope.init();
})
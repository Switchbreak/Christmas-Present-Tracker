angular.module("ChristmasTracker")
.controller("inviteCtrl", function($rootScope, $scope, $timeout, $location, person) {
	$scope.getPeople = function() {
		$scope.people = person.query();
		return $scope.people.$promise.catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	};
	
	$scope.personInvited = function( checkPerson ) {
		for( var i = 0; i < $scope.$parent.invitedPeople.length; i++ ) {
			if( $scope.$parent.invitedPeople[i].Name == checkPerson.Name )
				return false;
		};
		
		return true;
	}
	
	$scope.invite = function() {
		if( !$scope.people )
			$scope.getPeople();
		
		$scope.invitedPerson = new person();
		$scope.newPerson = new person();
	};
	
	$scope.invitePerson = function( invitedPerson ) {
		invitedPerson.saving = true;
		$scope.$parent.invitedPeople.push( invitedPerson );
		
		return $scope.$parent.party.$invite( { person: invitedPerson.Name } ).then(function() {
			invitedPerson.saving = false;
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		})
	};
	
	$scope.createPerson = function( newPerson ) {
		newPerson.saving = true;
		$scope.$parent.invitedPeople.push( newPerson );
		$scope.people.push( newPerson );
		
		return newPerson.$save().then( function() {
			return $scope.$parent.party.$invite( { person: newPerson.Name } );
		}).then( function() {
			newPerson.saving = false;
			$("#RegistrationLinkModal").modal();
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		})
	};	
	
	$scope.scrollToBottom = function( element ) {
		element.scrollTop(element.prop("scrollHeight")); 
	};
	
	$scope.host = $location.absUrl().replace($location.url(), "");
})
.directive("invitePerson", function() {
	return {
		scope: {
			party: '=',
		},
		controller: "inviteCtrl",
		templateUrl: 'static/components/invitePerson.html',
	}
})
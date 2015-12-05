angular.module("ChristmasTracker", ["ngRoute", "ngResource"])
	.config(function($routeProvider, $locationProvider) {
		$routeProvider
			.when("/", {controller: "partyListCtrl", templateUrl: "static/parties.html", resolve: { currentPerson: function( currentPerson ) { return currentPerson.queryPromise } } })
			.when("/party/:title", {controller: "partyCtrl", templateUrl: "static/party.html", resolve: { currentPerson: function( currentPerson ) { return currentPerson.queryPromise } } })
		$locationProvider.html5Mode(true);
	})
	
	.run(function($rootScope, currentPerson) {
		currentPerson.queryPromise.then(function() {
			$rootScope.currentPerson = currentPerson;
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	})
	
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
	.controller("partyCtrl", function($rootScope, $scope, $routeParams, party, invitedPeople, comment, person, currentPerson) {
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
		
		$scope.getPeople = function() {
			$scope.people = person.query();
			$scope.people.$promise.catch(function(reason) {
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
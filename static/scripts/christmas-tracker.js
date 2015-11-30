angular.module("ChristmasTracker", ["ngRoute", "ngResource"])
	.config(function($routeProvider, $locationProvider) {
		$routeProvider
			.when("/", {controller: "partyList", templateUrl: "static/parties.html", resolve: { currentPerson: function( currentPerson ) { return currentPerson.queryPromise } } });
		$locationProvider.html5Mode(true);
	})
	
	.factory("party", function($resource) {
		return $resource("/api/party/:title", { title: "@Title" }, {
			invite: { method: "POST", url: "/api/party/:title/invite/:person" }
		});
	})
	.factory("person", function($resource) {
		return $resource("/api/person/:name", { name: "@Name" }, {
			current: { method: "GET", url: "/api/person/current" },
			link: { method: "POST", url: "/api/person/:name/link" },
		});
	})
	.factory("currentPerson", function($resource, $q, person) {
		currentPerson = person.current()
		
		currentPerson.queryPromise = currentPerson.$promise.catch(function(reason) {
			if( reason.status == 404 ) {
				var newPerson = new person({ Name: "New Person" });
				return newPerson.$save().then( function() {
					return newPerson.$link( newPerson ).$promise;
				}).catch( function (reason) {
					return $q.reject(reason);
				});
			} else {
				return $q.reject(reason);
			}
		});
		
		return currentPerson;
	})
	
	.run(function($rootScope, currentPerson) {
		currentPerson.queryPromise.catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	})
	
	.controller("partyList", function($rootScope, $scope, party) {
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
angular.module("ChristmasTracker", ["ngRoute", "ngResource"])
	.config(function($routeProvider, $locationProvider) {
		$routeProvider
			.when("/", {controller: "partyListCtrl", templateUrl: "static/parties.html", resolve: { currentPerson: function( currentPerson ) { return currentPerson.$promise } } })
			.when("/new-person", {controller: "newPersonCtrl", templateUrl: "static/newPerson.html" })
			.when("/party/:title", {controller: "partyCtrl", templateUrl: "static/party.html", resolve: { currentPerson: function( currentPerson ) { return currentPerson.$promise } } })
			.when("/party/:title/:name", {controller: "personCtrl", templateUrl: "static/person.html", resolve: { currentPerson: function( currentPerson ) { return currentPerson.$promise } } })
		$locationProvider.html5Mode(true);
	})
	
	.run(function($rootScope, $location, currentPerson) {
		currentPerson.$promise.then(function() {
			$rootScope.currentPerson = currentPerson;
		}).catch(function(reason) {
			if( reason.status == 404 ) {
				$location.path("/new-person");
			} else {
				$rootScope.errorMessage = reason;
			}
		});
	})
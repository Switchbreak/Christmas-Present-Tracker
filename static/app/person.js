angular.module("ChristmasTracker")
.controller("personCtrl", function($rootScope, $scope, $routeParams, $timeout, $location, party, person, wishlist, boughtItem) {
	$scope.getParty = function() {
		$scope.party = party.get( { title: $routeParams.title } );
		$scope.party.$promise.catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	};
	
	$scope.getPerson = function() {
		$scope.person = person.get( { name: $routeParams.name } );
		$scope.person.$promise.catch(function(reason) {
			$rootScope.errorMessage = reason;
		})
	};
	
	$scope.getWishlist = function() {
		$scope.wishlist = wishlist.query( { party: $routeParams.title, person: $routeParams.name } );
		$scope.wishlist.$promise.catch(function(reason) {
			$rootScope.errorMessage = reason;
		})
		
		$scope.newWishlistItem = new wishlist();
	};
	
	$scope.saveWishlistItem = function( wishlistItem ) {
		wishlistItem.saving = true;
		$scope.wishlist.push( wishlistItem );
		$timeout( function() { $scope.scrollToBottom($('#Wishlist')) } );
		
		wishlistItem.$save( { party: $routeParams.title, person: $routeParams.name } ).then(function(data) {
			wishlistItem.saving = false;
			wishlistItem.ID = data.ID;
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
		
		$scope.newWishlistItem = new wishlist();
	};
	
	$scope.deleteWishlistItem = function( wishlistItem ) {
		$scope.wishlist.splice( $scope.wishlist.indexOf( wishlistItem ), 1 );
		
		wishlistItem.$delete( { party: $routeParams.title, person: $routeParams.name } ).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	};
	
	$scope.getBoughtItems = function() {
		$scope.boughtItems = boughtItem.query( { party: $routeParams.title, person: $routeParams.name } );
		$scope.boughtItems.$promise.catch(function(reason) {
			$rootScope.errorMessage = reason;
		})
		
		$scope.newBoughtItem = new boughtItem();
	};
	
	$scope.saveBoughtItem = function( newBoughtItem ) {
		newBoughtItem.saving = true;
		$scope.boughtItems.push( newBoughtItem );
		$timeout( function() { $scope.scrollToBottom($('#BoughtItems')) } );
		
		newBoughtItem.$save( { party: $routeParams.title, person: $routeParams.name } ).then(function(data) {
			newBoughtItem.saving = false;
			newBoughtItem.ID = data.ID;
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
		
		$scope.newBoughtItem = new boughtItem();
	};
	
	$scope.deleteBoughtItem = function( newBoughtItem ) {
		$scope.boughtItems.splice( $scope.boughtItems.indexOf( newBoughtItem ), 1 );
		
		newBoughtItem.$delete( { party: $routeParams.title, person: $routeParams.name } ).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
	};
	
	$scope.scrollToBottom = function( element ) {
		element.scrollTop(element.prop("scrollHeight")); 
	};
	
	$scope.init = function() {
		$scope.partyTitle = $routeParams.title;
		$scope.personName = $routeParams.name;
		$scope.viewingSelf = ($scope.personName == $rootScope.currentPerson.Name);
		$scope.host = $location.absUrl().replace($location.url(), "");
		
		$scope.getParty();
		$scope.getPerson();
		$scope.getWishlist();
		if( !$scope.viewingSelf )
			$scope.getBoughtItems();
	}
	
	$scope.init();
})
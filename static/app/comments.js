angular.module("ChristmasTracker")
.controller("commentsCtrl", function($rootScope, $scope, $timeout, comment) {
	$scope.getComments = function() {
		$scope.comments = comment.query({ title: $scope.party })
		$scope.comments.$promise.then(function() {
			$timeout( function() { $scope.scrollToBottom($('#comments')) } );
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
		
		$scope.newComment = new comment();
	};
	
	$scope.saveComment = function( newComment ) {
		newComment.Author = currentPerson.Name;
		newComment.Date = new Date();
		newComment.saving = true;
		$scope.comments.push( newComment );
		$timeout( function() { $scope.scrollToBottom($('#comments')) });
		
		newComment.$save( { title: $scope.party } ).then( function() {
			newComment.saving = false;
		}).catch(function(reason) {
			$rootScope.errorMessage = reason;
		});
		
		$scope.newComment = new comment();
	};
	
	$scope.scrollToBottom = function( element ) {
		element.scrollTop(element.prop("scrollHeight")); 
	};
	
	$scope.getComments();
})
.directive("comments", function() {
	return {
		scope: {
			party: '=',
			person: '=',
		},
		controller: "commentsCtrl",
		templateUrl: 'static/comments.html',
	}
})
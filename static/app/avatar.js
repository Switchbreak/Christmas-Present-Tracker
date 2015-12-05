var hexCodes = "0123456789abcdef";

angular.module("ChristmasTracker")
.directive("avatar", function() {
	return {
		scope: {
			name: '=name',
		},
		controller: function($scope) {
			$scope.avatarColor = function( name ) {
				if( !name )
					return;
				
				var rng = new Math.seedrandom(name);
				
				var avatarColor = "#";
				for( var i = 0; i < 6; i++ ) {
					avatarColor += hexCodes[Math.floor( rng.quick() * 16 )];
				}
				
				return avatarColor;
			};
			
			$scope.initials = function( name ) {
				if( !name )
					return;
				
				var names = name.split(" ");
				var initials = "";
				for( var i = 0; i < names.length; i++ ) {
					initials += names[i].substr(0, 1).toUpperCase();
				}
				
				return initials;
			};
		},
		template: '<span class="avatar-no-image" style="background: {{avatarColor(name)}}" title="{{name}}"><h4 class="avatar-text">{{initials(name)}}</h4></span>',
	}
});
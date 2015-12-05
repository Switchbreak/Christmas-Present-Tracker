angular.module("ChristmasTracker")
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
.factory("invitedPeople", function($resource) {
	return $resource("/api/party/:title/invited");
})
.factory("comment", function($resource) {
	return $resource("/api/party/:title/:person/comment/:id", { id: "@ID" });
})

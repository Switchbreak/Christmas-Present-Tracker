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

.factory("currentPerson", function($resource, $q, $location, person) {
	currentPerson = person.current();
	
	return currentPerson;
})

.factory("invitedPeople", function($resource) {
	return $resource("/api/party/:title/invited");
})

.factory("comment", function($resource) {
	return $resource("/api/party/:title/:person/comment/:id", { id: "@ID" });
})

.factory("wishlist", function($resource) {
	return $resource("/api/party/:party/:person/wishlist/:id", { id: "@ID" });
})

.factory("boughtItem", function($resource) {
	return $resource("/api/party/:party/:person/bought/:id", { id: "@ID" });
})
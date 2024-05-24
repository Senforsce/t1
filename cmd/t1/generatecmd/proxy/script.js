let src = new EventSource("/_t1/reload/events");
src.onmessage = (event) => {
	if (event && event.data === "reload") {
		window.location.reload();
	}
};

const socket = new WebSocket("/ws");
const pointStyle = {
	stroke: "rgba(255, 0, 0, 1)",
	fill: "rgba(255, 0, 0, 0.8)"
}
var config = {
	form: false,
	datapath: "/static/data/",
	center: [0, 0, 0],
	constellations: {
		names: false,
		lines: false,
	},
	dsos: {
		show: false,
	},
};
Celestial.display(config);

let dynamicPoints = {
	type: "FeatureCollection",
	features: []
};

let audioSystem;
document.addEventListener('click', async () => {
	if (!audioSystem) {
		audioSystem = await initializeAudio();
		// Start the ambient pad
		audioSystem.ambience.start();
	}
});

Celestial.add({
	type: "point",
	callback: function() {
		if (error) return console.warn(error);
		console.log("Add callback initiated");
	},
	redraw: function() {
		dynamicPoints.features.forEach(function(point) {
			if (Celestial.clip(point.geometry.coordinates)) {
				// Get point coordinates
				var pt = Celestial.mapProjection(point.geometry.coordinates);
				// Object radius in pixels
				var r = Math.pow(parseInt(point.properties.dim || 5) * 0.5, 0.5);
				// Draw on canvas
				Celestial.setStyle(pointStyle);
				Celestial.context.beginPath();
				Celestial.context.arc(pt[0], pt[1], r, 0, 2 * Math.PI);
				Celestial.context.closePath();
				Celestial.context.stroke();
				Celestial.context.fill();
			}
		});
	},
});

// Receive data from the websocket
// and add it to the dynamicPoints object
// then redraw the map
socket.addEventListener("message", function(event) {
	const data = JSON.parse(event.data);
	dynamicPoints.features = dynamicPoints.features.concat(data);
	let len = dynamicPoints.features.length;
	if (len > 100) {
		dynamicPoints.features = dynamicPoints.features.slice(len - 100, len);
	}
	Celestial.redraw();
	playNote()
});

/**
 * Plays a random note from the scale
 */
function playNote() {
	if (!audioSystem) return; // Guard against events before initialization

	audioSystem.spaceSound.playRandomNote(0.4);
}

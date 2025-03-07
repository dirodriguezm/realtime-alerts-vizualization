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
				paintPoint(point)
			}
		});
	},
});

/**
 * Paints a point in celestial map
 * @param {Object} point 
 */
function paintPoint(point) {
	// Get point coordinates
	var pt = Celestial.mapProjection(point.geometry.coordinates);

	// Object radius in pixels
	if (point.properties.currentRadius === undefined) {
		point.properties.currentRadius = Math.pow(point.properties.dim, 0.5)
	}
	//point.properties.currentRadius = point.properties.currentRadius * 0.9
	let r = point.properties.currentRadius

	// Draw on canvas
	if (point.properties.style === undefined) {
		point.properties.style = JSON.parse(JSON.stringify(pointStyle))
	}
	Celestial.setStyle(point.properties.style);
	Celestial.context.beginPath();
	Celestial.context.arc(pt[0], pt[1], r, 0, 2 * Math.PI);
	Celestial.context.closePath();
	Celestial.context.stroke();
	Celestial.context.fill();
}

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


function reduceOpacity() {
	// Set style to each point
	dynamicPoints.features = dynamicPoints.features.map((point) => {
		let currentPointStyle = JSON.parse(JSON.stringify(point.properties.style));
		if (point.properties.style === undefined) {
			currentPointStyle = JSON.parse(JSON.stringify(pointStyle));
		}
		let fillOpacity = parseFloat(currentPointStyle.fill.match(/[\d\.]+/g)[3])
		let strokeOpacity = parseFloat(currentPointStyle.stroke.match(/[\d\.]+/g)[3])
		// reduce opacity by 10%
		fillOpacity = fillOpacity * 0.9
		strokeOpacity = strokeOpacity * 0.9

		// delete point if opacity is less than 0.1
		if (fillOpacity < 0.1) {
			return null
		}

		// set new opacity values
		currentPointStyle.fill = `rgba(255, 0, 0, ${fillOpacity})`
		currentPointStyle.stroke = `rgba(255, 0, 0, ${strokeOpacity})`
		point.properties.style = currentPointStyle
		return point
	}).filter((point) => point !== null)
	if (dynamicPoints.features.length) {
		Celestial.redraw();
	}
}

setInterval(reduceOpacity, 500);

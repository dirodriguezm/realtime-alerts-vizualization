class SpaceSound {
	constructor() {
		// Initialize effects chain for ethereal space sound
		this.reverb = new Tone.Reverb({
			decay: 5,
			wet: 0.5
		}).toDestination();

		this.delay = new Tone.PingPongDelay({
			delayTime: "8n",
			feedback: 0.2,
			wet: 0.3
		}).toDestination();

		// Initialize synth with space-appropriate settings
		this.synth = new Tone.PolySynth(Tone.FMSynth, {
			harmonicity: 2,
			modulationIndex: 3,
			oscillator: {
				type: "sine"
			},
			envelope: {
				attack: 0.8,
				decay: 0.4,
				sustain: 0.5,
				release: 1
			},
			modulation: {
				type: "triangle"
			}
		}).connect(this.reverb).connect(this.delay);

		// Lydian scale starting from C4
		this.scale = ["C4", "D4", "E4", "F#4", "G4", "A4", "B4", "C5"];

		// Keep track of last played note to avoid repetition
		this.lastPlayedIndex = -1;

		// Sound management properties
		this.isPlaying = false;
		this.maxSimultaneousSounds = 10; // Maximum number of sounds that can play simultaneously
		this.activeSounds = 0;
	}

	/**
	 * Checks if a new sound can be played
	 * @returns {boolean}
	 */
	canPlayNewSound() {
		return this.activeSounds < this.maxSimultaneousSounds;
	}


	/**
	 * Play sound if not too many sounds are playing
	 * @param {Function} soundFunction - Function that produces the sound
	 * @returns {boolean} - Whether the sound was played
	 */
	playSound(soundFunction) {
		if (!this.canPlayNewSound()) {
			console.log('Too many sounds playing. Skipping sound.');
			return false;
		}

		try {
			this.activeSounds++;
			soundFunction();
		} finally {
			this.activeSounds--;
		}
		return true;
	}

	/**
	 * Plays a random note from the scale
	 * @param {number} duration - note duration in seconds (default: 2)
	 */
	playRandomNote(duration = 2) {
		const noteFunction = () => {
			let noteIndex;
			do {
				noteIndex = Math.floor(Math.random() * this.scale.length);
			} while (noteIndex === this.lastPlayedIndex && this.scale.length > 1);

			this.lastPlayedIndex = noteIndex;
			this.synth.triggerAttackRelease(this.scale[noteIndex], duration);
		};

		return this.playSound(noteFunction);
	}

	/**
	 * Update sound limiting parameters
	 * @param {Object} params - Parameters to update
	 */
	updateParameters({
		maxSimultaneousSounds
	} = {}) {
		if (maxSimultaneousSounds !== undefined) this.maxSimultaneousSounds = maxSimultaneousSounds;
	}

	/**
	 * Get current sound system status
	 */
	getStatus() {
		return {
			activeSounds: this.activeSounds,
			isPlaying: this.isPlaying,
		};
	}
}

async function initializeAudio() {
	const ambience = new SpaceAmbience();
	const spaceSound = new SpaceSound(ambience);

	// Start audio context (must be triggered by user interaction)
	await Tone.start();

	return { spaceSound, ambience };
};


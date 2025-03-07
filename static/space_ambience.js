class SpaceAmbience {
	constructor() {
		// Create effects for the pad
		this.reverb = new Tone.Reverb({
			decay: 8,
			wet: 0.6
		}).toDestination();

		this.chorus = new Tone.Chorus({
			frequency: 0.1,
			delayTime: 4,
			depth: 0.7,
			wet: 0.5
		}).connect(this.reverb);

		// Main pad synth
		this.padSynth = new Tone.PolySynth(Tone.FMSynth, {
			harmonicity: 1.5,
			modulationIndex: 1,
			oscillator: {
				type: "sine"
			},
			envelope: {
				attack: 2,
				decay: 0.8,
				sustain: 1,
				release: 4
			},
			modulation: {
				type: "triangle"
			},
			modulationEnvelope: {
				attack: 4,
				decay: 0.8,
				sustain: 1,
				release: 4
			},
			volume: -15 // Lower volume for background
		}).connect(this.chorus);

		// Simple chord progression in C Lydian
		this.chordProgression = [
			["C2", "C3", "D3", "E3"],
			["D3", "A3", "D4", "F#4"]
			["B2", "B3", "D3", "F#3"],
			["E2", "B3", "E3", "G3"],
		];

		this.isPlaying = false;
		this.currentChord = 0;
		this.loop = null;
	}

	start() {
		if (this.isPlaying) return;

		// Set up the loop
		this.loop = new Tone.Loop((time) => {
			const chord = this.chordProgression[this.currentChord];
			this.padSynth.triggerAttackRelease(chord, "4n", time);
			this.currentChord = (this.currentChord + 1) % this.chordProgression.length;
		}, "8n").start(0);

		// Start transport if it's not already running
		if (Tone.Transport.state !== "started") {
			Tone.Transport.bpm.value = 40; // Slow BPM for ambient feel
			Tone.Transport.start();
		}

		this.isPlaying = true;
	}

	stop() {
		if (!this.isPlaying) return;

		if (this.loop) {
			this.loop.stop();
			this.loop.dispose();
			this.loop = null;
		}

		this.isPlaying = false;
		this.currentChord = 0;
	}
}

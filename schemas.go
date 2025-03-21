// Code generated by avro/gen. DO NOT EDIT.
package main

// avro alert schema.
type ZtfAlertCandidate struct {
	// Observation Julian date at start of exposure [days].
	Jd float64 `avro:"jd" json:"jd"`
	// Filter ID (1=g; 2=R; 3=i).
	Fid int `avro:"fid" json:"fid"`
	// Processing ID for science image to facilitate archive retrieval.
	Pid int64 `avro:"pid" json:"pid"`
	// Expected 5-sigma mag limit in difference image based on global noise estimate [mag].
	Diffmaglim *float32 `avro:"diffmaglim" json:"diffmaglim"`
	// filename of positive (sci minus ref) difference image.
	Pdiffimfilename *string `avro:"pdiffimfilename" json:"pdiffimfilename"`
	// Principal investigator attached to program ID.
	Programpi *string `avro:"programpi" json:"programpi"`
	// Program ID: encodes either public, collab, or caltech mode.
	Programid int `avro:"programid" json:"programid"`
	// Candidate ID from operations DB.
	Candid int64 `avro:"candid" json:"candid"`
	// t or 1 => candidate is from positive (sci minus ref) subtraction; f or 0 => candidate is from negative (ref minus sci) subtraction.
	Isdiffpos string `avro:"isdiffpos" json:"isdiffpos"`
	// Internal pipeline table extraction ID.
	Tblid *int64 `avro:"tblid" json:"tblid"`
	// Night ID.
	Nid *int `avro:"nid" json:"nid"`
	// Readout channel ID [00 .. 63].
	Rcid *int `avro:"rcid" json:"rcid"`
	// ZTF field ID.
	Field *int `avro:"field" json:"field"`
	// x-image position of candidate [pixels].
	Xpos *float32 `avro:"xpos" json:"xpos"`
	// y-image position of candidate [pixels].
	Ypos *float32 `avro:"ypos" json:"ypos"`
	// Right Ascension of candidate; J2000 [deg].
	Ra float64 `avro:"ra" json:"ra"`
	// Declination of candidate; J2000 [deg].
	Dec float64 `avro:"dec" json:"dec"`
	// Magnitude from PSF-fit photometry [mag].
	Magpsf float32 `avro:"magpsf" json:"magpsf"`
	// 1-sigma uncertainty in magpsf [mag].
	Sigmapsf float32 `avro:"sigmapsf" json:"sigmapsf"`
	// Reduced chi-square for PSF-fit.
	Chipsf *float32 `avro:"chipsf" json:"chipsf"`
	// Aperture mag using 14 pixel diameter aperture [mag].
	Magap *float32 `avro:"magap" json:"magap"`
	// 1-sigma uncertainty in magap [mag].
	Sigmagap *float32 `avro:"sigmagap" json:"sigmagap"`
	// distance to nearest source in reference image PSF-catalog [pixels].
	Distnr *float32 `avro:"distnr" json:"distnr"`
	// magnitude of nearest source in reference image PSF-catalog [mag].
	Magnr *float32 `avro:"magnr" json:"magnr"`
	// 1-sigma uncertainty in magnr [mag].
	Sigmagnr *float32 `avro:"sigmagnr" json:"sigmagnr"`
	// DAOPhot chi parameter of nearest source in reference image PSF-catalog.
	Chinr *float32 `avro:"chinr" json:"chinr"`
	// DAOPhot sharp parameter of nearest source in reference image PSF-catalog.
	Sharpnr *float32 `avro:"sharpnr" json:"sharpnr"`
	// Local sky background estimate [DN].
	Sky *float32 `avro:"sky" json:"sky"`
	// Difference: magap - magpsf [mag].
	Magdiff *float32 `avro:"magdiff" json:"magdiff"`
	// Full Width Half Max assuming a Gaussian core, from SExtractor [pixels].
	Fwhm *float32 `avro:"fwhm" json:"fwhm"`
	// Star/Galaxy classification score from SExtractor.
	Classtar *float32 `avro:"classtar" json:"classtar"`
	// Distance to nearest edge in image [pixels].
	Mindtoedge *float32 `avro:"mindtoedge" json:"mindtoedge"`
	// Difference: diffmaglim - magap [mag].
	Magfromlim *float32 `avro:"magfromlim" json:"magfromlim"`
	// Ratio: difffwhm / fwhm.
	Seeratio *float32 `avro:"seeratio" json:"seeratio"`
	// Windowed profile RMS afloat major axis from SExtractor [pixels].
	Aimage *float32 `avro:"aimage" json:"aimage"`
	// Windowed profile RMS afloat minor axis from SExtractor [pixels].
	Bimage *float32 `avro:"bimage" json:"bimage"`
	// Ratio: aimage / fwhm.
	Aimagerat *float32 `avro:"aimagerat" json:"aimagerat"`
	// Ratio: bimage / fwhm.
	Bimagerat *float32 `avro:"bimagerat" json:"bimagerat"`
	// Ratio: aimage / bimage.
	Elong *float32 `avro:"elong" json:"elong"`
	// number of negative pixels in a 5 x 5 pixel stamp.
	Nneg *int `avro:"nneg" json:"nneg"`
	// number of prior-tagged bad pixels in a 5 x 5 pixel stamp.
	Nbad *int `avro:"nbad" json:"nbad"`
	// RealBogus quality score from Random Forest classifier; range is 0 to 1 where closer to 1 is more reliable.
	Rb *float32 `avro:"rb" json:"rb"`
	// distance to nearest known solar system object if exists within 30 arcsec [arcsec].
	Ssdistnr *float32 `avro:"ssdistnr" json:"ssdistnr"`
	// magnitude of nearest known solar system object if exists within 30 arcsec (usually V-band from MPC archive) [mag].
	Ssmagnr *float32 `avro:"ssmagnr" json:"ssmagnr"`
	// name of nearest known solar system object if exists within 30 arcsec (from MPC archive).
	Ssnamenr *string `avro:"ssnamenr" json:"ssnamenr"`
	// Ratio: sum(pixels) / sum(|pixels|) in a 5 x 5 pixel stamp where stamp is first median-filtered to mitigate outliers.
	Sumrat *float32 `avro:"sumrat" json:"sumrat"`
	// Aperture mag using 18 pixel diameter aperture [mag].
	Magapbig *float32 `avro:"magapbig" json:"magapbig"`
	// 1-sigma uncertainty in magapbig [mag].
	Sigmagapbig *float32 `avro:"sigmagapbig" json:"sigmagapbig"`
	// Right Ascension of nearest source in reference image PSF-catalog; J2000 [deg].
	Ranr float64 `avro:"ranr" json:"ranr"`
	// Declination of nearest source in reference image PSF-catalog; J2000 [deg].
	Decnr float64 `avro:"decnr" json:"decnr"`
	// g-band PSF-fit magnitude of closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Sgmag1 *float32 `avro:"sgmag1" json:"sgmag1"`
	// r-band PSF-fit magnitude of closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Srmag1 *float32 `avro:"srmag1" json:"srmag1"`
	// i-band PSF-fit magnitude of closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Simag1 *float32 `avro:"simag1" json:"simag1"`
	// z-band PSF-fit magnitude of closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Szmag1 *float32 `avro:"szmag1" json:"szmag1"`
	// Star/Galaxy score of closest source from PS1 catalog; if exists within 30 arcsec: 0 <= sgscore <= 1 where closer to 1 implies higher likelihood of being a star.
	Sgscore1 *float32 `avro:"sgscore1" json:"sgscore1"`
	// Distance to closest source from PS1 catalog; if exists within 30 arcsec [arcsec].
	Distpsnr1 *float32 `avro:"distpsnr1" json:"distpsnr1"`
	// Number of spatially-coincident detections falling within 1.5 arcsec going back to beginning of survey; only detections that fell on the same field and readout-channel ID where the input candidate was observed are counted. All raw detections down to a photometric S/N of ~ 3 are included.
	Ndethist int `avro:"ndethist" json:"ndethist"`
	// Number of times input candidate position fell on any field and readout-channel going back to beginning of survey.
	Ncovhist int `avro:"ncovhist" json:"ncovhist"`
	// Earliest Julian date of epoch corresponding to ndethist [days].
	Jdstarthist *float64 `avro:"jdstarthist" json:"jdstarthist"`
	// Latest Julian date of epoch corresponding to ndethist [days].
	Jdendhist *float64 `avro:"jdendhist" json:"jdendhist"`
	// Peak-pixel signal-to-noise ratio in point source matched-filtered detection image.
	Scorr *float64 `avro:"scorr" json:"scorr"`
	// 1 => candidate is from a Target-of-Opportunity (ToO) exposure; 0 => candidate is from a non-ToO exposure.
	Tooflag *int `avro:"tooflag" json:"tooflag"`
	// Object ID of closest source from PS1 catalog; if exists within 30 arcsec.
	Objectidps1 *int64 `avro:"objectidps1" json:"objectidps1"`
	// Object ID of second closest source from PS1 catalog; if exists within 30 arcsec.
	Objectidps2 *int64 `avro:"objectidps2" json:"objectidps2"`
	// g-band PSF-fit magnitude of second closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Sgmag2 *float32 `avro:"sgmag2" json:"sgmag2"`
	// r-band PSF-fit magnitude of second closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Srmag2 *float32 `avro:"srmag2" json:"srmag2"`
	// i-band PSF-fit magnitude of second closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Simag2 *float32 `avro:"simag2" json:"simag2"`
	// z-band PSF-fit magnitude of second closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Szmag2 *float32 `avro:"szmag2" json:"szmag2"`
	// Star/Galaxy score of second closest source from PS1 catalog; if exists within 30 arcsec: 0 <= sgscore <= 1 where closer to 1 implies higher likelihood of being a star.
	Sgscore2 *float32 `avro:"sgscore2" json:"sgscore2"`
	// Distance to second closest source from PS1 catalog; if exists within 30 arcsec [arcsec].
	Distpsnr2 *float32 `avro:"distpsnr2" json:"distpsnr2"`
	// Object ID of third closest source from PS1 catalog; if exists within 30 arcsec.
	Objectidps3 *int64 `avro:"objectidps3" json:"objectidps3"`
	// g-band PSF-fit magnitude of third closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Sgmag3 *float32 `avro:"sgmag3" json:"sgmag3"`
	// r-band PSF-fit magnitude of third closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Srmag3 *float32 `avro:"srmag3" json:"srmag3"`
	// i-band PSF-fit magnitude of third closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Simag3 *float32 `avro:"simag3" json:"simag3"`
	// z-band PSF-fit magnitude of third closest source from PS1 catalog; if exists within 30 arcsec [mag].
	Szmag3 *float32 `avro:"szmag3" json:"szmag3"`
	// Star/Galaxy score of third closest source from PS1 catalog; if exists within 30 arcsec: 0 <= sgscore <= 1 where closer to 1 implies higher likelihood of being a star.
	Sgscore3 *float32 `avro:"sgscore3" json:"sgscore3"`
	// Distance to third closest source from PS1 catalog; if exists within 30 arcsec [arcsec].
	Distpsnr3 *float32 `avro:"distpsnr3" json:"distpsnr3"`
	// Number of source matches from PS1 catalog falling within 30 arcsec.
	Nmtchps int `avro:"nmtchps" json:"nmtchps"`
	// Processing ID for reference image to facilitate archive retrieval.
	Rfid int64 `avro:"rfid" json:"rfid"`
	// Observation Julian date of earliest exposure used to generate reference image [days].
	Jdstartref float64 `avro:"jdstartref" json:"jdstartref"`
	// Observation Julian date of latest exposure used to generate reference image [days].
	Jdendref float64 `avro:"jdendref" json:"jdendref"`
	// Number of frames (epochal images) used to generate reference image.
	Nframesref int `avro:"nframesref" json:"nframesref"`
	// version of Random Forest classifier model used to assign RealBogus (rb) quality score.
	Rbversion string `avro:"rbversion" json:"rbversion"`
	// Ratio: D/stddev(D) on event position where D = difference image.
	Dsnrms *float32 `avro:"dsnrms" json:"dsnrms"`
	// Ratio: S/stddev(S) on event position where S = image of convolution: D (x) PSF(D).
	Ssnrms *float32 `avro:"ssnrms" json:"ssnrms"`
	// Difference of statistics: dsnrms - ssnrms.
	Dsdiff *float32 `avro:"dsdiff" json:"dsdiff"`
	// Magnitude zero point for photometry estimates [mag].
	Magzpsci *float32 `avro:"magzpsci" json:"magzpsci"`
	// Magnitude zero point uncertainty (in magzpsci) [mag].
	Magzpsciunc *float32 `avro:"magzpsciunc" json:"magzpsciunc"`
	// RMS (deviation from average) in all differences between instrumental photometry and matched photometric calibrators from science image processing [mag].
	Magzpscirms *float32 `avro:"magzpscirms" json:"magzpscirms"`
	// Number of PS1 photometric calibrators used to calibrate science image from science image processing.
	Nmatches int `avro:"nmatches" json:"nmatches"`
	// Color coefficient from linear fit from photometric calibration of science image.
	Clrcoeff *float32 `avro:"clrcoeff" json:"clrcoeff"`
	// Color coefficient uncertainty from linear fit (corresponding to clrcoeff).
	Clrcounc *float32 `avro:"clrcounc" json:"clrcounc"`
	// Covariance in magzpsci and clrcoeff from science image processing [mag^2].
	Zpclrcov *float32 `avro:"zpclrcov" json:"zpclrcov"`
	// Magnitude zero point from median of all differences between instrumental photometry and matched photometric calibrators from science image processing [mag].
	Zpmed *float32 `avro:"zpmed" json:"zpmed"`
	// Median color of all PS1 photometric calibrators used from science image processing [mag]: for filter (fid) = 1, 2, 3, PS1 color used = g-r, g-r, r-i respectively.
	Clrmed *float32 `avro:"clrmed" json:"clrmed"`
	// RMS color (deviation from average) of all PS1 photometric calibrators used from science image processing [mag].
	Clrrms *float32 `avro:"clrrms" json:"clrrms"`
	// Distance to closest source from Gaia DR1 catalog irrespective of magnitude; if exists within 90 arcsec [arcsec].
	Neargaia *float32 `avro:"neargaia" json:"neargaia"`
	// Distance to closest source from Gaia DR1 catalog brighter than magnitude 14; if exists within 90 arcsec [arcsec].
	Neargaiabright *float32 `avro:"neargaiabright" json:"neargaiabright"`
	// Gaia (G-band) magnitude of closest source from Gaia DR1 catalog irrespective of magnitude; if exists within 90 arcsec [mag].
	Maggaia *float32 `avro:"maggaia" json:"maggaia"`
	// Gaia (G-band) magnitude of closest source from Gaia DR1 catalog brighter than magnitude 14; if exists within 90 arcsec [mag].
	Maggaiabright *float32 `avro:"maggaiabright" json:"maggaiabright"`
	// Integration time of camera exposure [sec].
	Exptime *float32 `avro:"exptime" json:"exptime"`
	// RealBogus quality score from Deep-Learning-based classifier; range is 0 to 1 where closer to 1 is more reliable.
	Drb *float32 `avro:"drb" json:"drb"`
	// version of Deep-Learning-based classifier model used to assign RealBogus (drb) quality score.
	Drbversion string `avro:"drbversion" json:"drbversion"`
}

// avro alert schema.
type ZtfAlertPrvCandidate struct {
	// Observation Julian date at start of exposure [days].
	Jd float64 `avro:"jd" json:"jd"`
	// Filter ID (1=g; 2=R; 3=i).
	Fid int `avro:"fid" json:"fid"`
	// Processing ID for image.
	Pid int64 `avro:"pid" json:"pid"`
	// Expected 5-sigma mag limit in difference image based on global noise estimate [mag].
	Diffmaglim *float32 `avro:"diffmaglim" json:"diffmaglim"`
	// filename of positive (sci minus ref) difference image.
	Pdiffimfilename *string `avro:"pdiffimfilename" json:"pdiffimfilename"`
	// Principal investigator attached to program ID.
	Programpi *string `avro:"programpi" json:"programpi"`
	// Program ID: encodes either public, collab, or caltech mode.
	Programid int `avro:"programid" json:"programid"`
	// Candidate ID from operations DB.
	Candid *int64 `avro:"candid" json:"candid"`
	// t or 1 => candidate is from positive (sci minus ref) subtraction; f or 0 => candidate is from negative (ref minus sci) subtraction.
	Isdiffpos *string `avro:"isdiffpos" json:"isdiffpos"`
	// Internal pipeline table extraction ID.
	Tblid *int64 `avro:"tblid" json:"tblid"`
	// Night ID.
	Nid *int `avro:"nid" json:"nid"`
	// Readout channel ID [00 .. 63].
	Rcid *int `avro:"rcid" json:"rcid"`
	// ZTF field ID.
	Field *int `avro:"field" json:"field"`
	// x-image position of candidate [pixels].
	Xpos *float32 `avro:"xpos" json:"xpos"`
	// y-image position of candidate [pixels].
	Ypos *float32 `avro:"ypos" json:"ypos"`
	// Right Ascension of candidate; J2000 [deg].
	Ra *float64 `avro:"ra" json:"ra"`
	// Declination of candidate; J2000 [deg].
	Dec *float64 `avro:"dec" json:"dec"`
	// Magnitude from PSF-fit photometry [mag].
	Magpsf *float32 `avro:"magpsf" json:"magpsf"`
	// 1-sigma uncertainty in magpsf [mag].
	Sigmapsf *float32 `avro:"sigmapsf" json:"sigmapsf"`
	// Reduced chi-square for PSF-fit.
	Chipsf *float32 `avro:"chipsf" json:"chipsf"`
	// Aperture mag using 14 pixel diameter aperture [mag].
	Magap *float32 `avro:"magap" json:"magap"`
	// 1-sigma uncertainty in magap [mag].
	Sigmagap *float32 `avro:"sigmagap" json:"sigmagap"`
	// distance to nearest source in reference image PSF-catalog [pixels].
	Distnr *float32 `avro:"distnr" json:"distnr"`
	// magnitude of nearest source in reference image PSF-catalog [mag].
	Magnr *float32 `avro:"magnr" json:"magnr"`
	// 1-sigma uncertainty in magnr [mag].
	Sigmagnr *float32 `avro:"sigmagnr" json:"sigmagnr"`
	// DAOPhot chi parameter of nearest source in reference image PSF-catalog.
	Chinr *float32 `avro:"chinr" json:"chinr"`
	// DAOPhot sharp parameter of nearest source in reference image PSF-catalog.
	Sharpnr *float32 `avro:"sharpnr" json:"sharpnr"`
	// Local sky background estimate [DN].
	Sky *float32 `avro:"sky" json:"sky"`
	// Difference: magap - magpsf [mag].
	Magdiff *float32 `avro:"magdiff" json:"magdiff"`
	// Full Width Half Max assuming a Gaussian core, from SExtractor [pixels].
	Fwhm *float32 `avro:"fwhm" json:"fwhm"`
	// Star/Galaxy classification score from SExtractor.
	Classtar *float32 `avro:"classtar" json:"classtar"`
	// Distance to nearest edge in image [pixels].
	Mindtoedge *float32 `avro:"mindtoedge" json:"mindtoedge"`
	// Difference: diffmaglim - magap [mag].
	Magfromlim *float32 `avro:"magfromlim" json:"magfromlim"`
	// Ratio: difffwhm / fwhm.
	Seeratio *float32 `avro:"seeratio" json:"seeratio"`
	// Windowed profile RMS afloat major axis from SExtractor [pixels].
	Aimage *float32 `avro:"aimage" json:"aimage"`
	// Windowed profile RMS afloat minor axis from SExtractor [pixels].
	Bimage *float32 `avro:"bimage" json:"bimage"`
	// Ratio: aimage / fwhm.
	Aimagerat *float32 `avro:"aimagerat" json:"aimagerat"`
	// Ratio: bimage / fwhm.
	Bimagerat *float32 `avro:"bimagerat" json:"bimagerat"`
	// Ratio: aimage / bimage.
	Elong *float32 `avro:"elong" json:"elong"`
	// number of negative pixels in a 5 x 5 pixel stamp.
	Nneg *int `avro:"nneg" json:"nneg"`
	// number of prior-tagged bad pixels in a 5 x 5 pixel stamp.
	Nbad *int `avro:"nbad" json:"nbad"`
	// RealBogus quality score; range is 0 to 1 where closer to 1 is more reliable.
	Rb *float32 `avro:"rb" json:"rb"`
	// distance to nearest known solar system object if exists within 30 arcsec [arcsec].
	Ssdistnr *float32 `avro:"ssdistnr" json:"ssdistnr"`
	// magnitude of nearest known solar system object if exists within 30 arcsec (usually V-band from MPC archive) [mag].
	Ssmagnr *float32 `avro:"ssmagnr" json:"ssmagnr"`
	// name of nearest known solar system object if exists within 30 arcsec (from MPC archive).
	Ssnamenr *string `avro:"ssnamenr" json:"ssnamenr"`
	// Ratio: sum(pixels) / sum(|pixels|) in a 5 x 5 pixel stamp where stamp is first median-filtered to mitigate outliers.
	Sumrat *float32 `avro:"sumrat" json:"sumrat"`
	// Aperture mag using 18 pixel diameter aperture [mag].
	Magapbig *float32 `avro:"magapbig" json:"magapbig"`
	// 1-sigma uncertainty in magapbig [mag].
	Sigmagapbig *float32 `avro:"sigmagapbig" json:"sigmagapbig"`
	// Right Ascension of nearest source in reference image PSF-catalog; J2000 [deg].
	Ranr *float64 `avro:"ranr" json:"ranr"`
	// Declination of nearest source in reference image PSF-catalog; J2000 [deg].
	Decnr *float64 `avro:"decnr" json:"decnr"`
	// Peak-pixel signal-to-noise ratio in point source matched-filtered detection image.
	Scorr *float64 `avro:"scorr" json:"scorr"`
	// Magnitude zero point for photometry estimates [mag].
	Magzpsci *float32 `avro:"magzpsci" json:"magzpsci"`
	// Magnitude zero point uncertainty (in magzpsci) [mag].
	Magzpsciunc *float32 `avro:"magzpsciunc" json:"magzpsciunc"`
	// RMS (deviation from average) in all differences between instrumental photometry and matched photometric calibrators from science image processing [mag].
	Magzpscirms *float32 `avro:"magzpscirms" json:"magzpscirms"`
	// Color coefficient from linear fit from photometric calibration of science image.
	Clrcoeff *float32 `avro:"clrcoeff" json:"clrcoeff"`
	// Color coefficient uncertainty from linear fit (corresponding to clrcoeff).
	Clrcounc *float32 `avro:"clrcounc" json:"clrcounc"`
	// version of RealBogus model/classifier used to assign rb quality score.
	Rbversion string `avro:"rbversion" json:"rbversion"`
}

// avro alert schema.
type ZtfAlertFpHist struct {
	// ZTF field ID.
	Field *int `avro:"field" json:"field"`
	// Readout channel ID [00 .. 63].
	Rcid *int `avro:"rcid" json:"rcid"`
	// Filter ID (1=g; 2=R; 3=i).
	Fid int `avro:"fid" json:"fid"`
	// Processing ID for image.
	Pid int64 `avro:"pid" json:"pid"`
	// Processing ID for reference image to facilitate archive retrieval.
	Rfid int64 `avro:"rfid" json:"rfid"`
	// Effective FWHM of sci image [pixels].
	Sciinpseeing *float32 `avro:"sciinpseeing" json:"sciinpseeing"`
	// Background level in sci image [DN].
	Scibckgnd *float32 `avro:"scibckgnd" json:"scibckgnd"`
	// Robust sigma per pixel in sci image [DN].
	Scisigpix *float32 `avro:"scisigpix" json:"scisigpix"`
	// Magnitude zero point for photometry estimates [mag].
	Magzpsci *float32 `avro:"magzpsci" json:"magzpsci"`
	// Magnitude zero point uncertainty (in magzpsci) [mag].
	Magzpsciunc *float32 `avro:"magzpsciunc" json:"magzpsciunc"`
	// RMS (deviation from average) in all differences between instrumental photometry and matched photometric calibrators from science image processing [mag].
	Magzpscirms *float32 `avro:"magzpscirms" json:"magzpscirms"`
	// Color coefficient from linear fit from photometric calibration of science image.
	Clrcoeff *float32 `avro:"clrcoeff" json:"clrcoeff"`
	// Color coefficient uncertainty from linear fit (corresponding to clrcoeff).
	Clrcounc *float32 `avro:"clrcounc" json:"clrcounc"`
	// Integration time of camera exposure [sec].
	Exptime *float32 `avro:"exptime" json:"exptime"`
	// Full sci image astrometric RMS along R.A. with respect to Gaia1 [arcsec].
	Adpctdif1 *float32 `avro:"adpctdif1" json:"adpctdif1"`
	// Full sci image astrometric RMS along Dec. with respect to Gaia1 [arcsec].
	Adpctdif2 *float32 `avro:"adpctdif2" json:"adpctdif2"`
	// Expected 5-sigma mag limit in difference image based on global noise estimate [mag].
	Diffmaglim *float32 `avro:"diffmaglim" json:"diffmaglim"`
	// Program ID: encodes either public, collab, or caltech mode.
	Programid int `avro:"programid" json:"programid"`
	// Observation Julian date at start of exposure [days].
	Jd float64 `avro:"jd" json:"jd"`
	// Forced difference image PSF-fit flux [DN].
	Forcediffimflux *float32 `avro:"forcediffimflux" json:"forcediffimflux"`
	// 1-sigma uncertainty in forcediffimflux [DN].
	Forcediffimfluxunc *float32 `avro:"forcediffimfluxunc" json:"forcediffimfluxunc"`
	// Forced photometry processing status codes (0 => no warnings); see documentation.
	Procstatus *string `avro:"procstatus" json:"procstatus"`
	// distance to nearest source in reference image PSF-catalog [arcsec].
	Distnr *float32 `avro:"distnr" json:"distnr"`
	// Right Ascension of nearest source in reference image PSF-catalog; J2000 [deg].
	Ranr float64 `avro:"ranr" json:"ranr"`
	// Declination of nearest source in reference image PSF-catalog; J2000 [deg].
	Decnr float64 `avro:"decnr" json:"decnr"`
	// magnitude of nearest source in reference image PSF-catalog [mag].
	Magnr *float32 `avro:"magnr" json:"magnr"`
	// 1-sigma uncertainty in magnr [mag].
	Sigmagnr *float32 `avro:"sigmagnr" json:"sigmagnr"`
	// DAOPhot chi parameter of nearest source in reference image PSF-catalog.
	Chinr *float32 `avro:"chinr" json:"chinr"`
	// DAOPhot sharp parameter of nearest source in reference image PSF-catalog.
	Sharpnr *float32 `avro:"sharpnr" json:"sharpnr"`
}

// avro alert schema.
type ZtfAlertCutout struct {
	FileName string `avro:"fileName" json:"file_name"`
	// fits.gz.
	StampData []byte `avro:"stampData" json:"stamp_data"`
}

// avro alert schema for ZTF (www.ztf.caltech.edu).
type ZtfAlert struct {
	// schema version used.
	Schemavsn string `avro:"schemavsn" json:"schemavsn"`
	// origin of alert packet.
	Publisher string `avro:"publisher" json:"publisher"`
	// object identifier or name.
	ObjectID         string                  `avro:"objectId" json:"object_id"`
	Candid           int64                   `avro:"candid" json:"candid"`
	Candidate        ZtfAlertCandidate       `avro:"candidate" json:"candidate"`
	PrvCandidates    *[]ZtfAlertPrvCandidate `avro:"prv_candidates" json:"prv_candidates"`
	FpHists          *[]ZtfAlertFpHist       `avro:"fp_hists" json:"fp_hists"`
	CutoutScience    *ZtfAlertCutout         `avro:"cutoutScience" json:"cutout_science"`
	CutoutTemplate   *ZtfAlertCutout         `avro:"cutoutTemplate" json:"cutout_template"`
	CutoutDifference *ZtfAlertCutout         `avro:"cutoutDifference" json:"cutout_difference"`
}

// Probabilities is a generated struct.
type Probabilities struct {
	Sn       float32 `avro:"SN" json:"sn"`
	Agn      float32 `avro:"AGN" json:"agn"`
	Vs       float32 `avro:"VS" json:"vs"`
	Asteroid float32 `avro:"asteroid" json:"asteroid"`
	Bogus    float32 `avro:"bogus" json:"bogus"`
}

// Early Classification.
type StampProbabilities struct {
	ObjectID      string        `avro:"objectId" json:"object_id"`
	Candid        int64         `avro:"candid" json:"candid"`
	Probabilities Probabilities `avro:"probabilities" json:"probabilities"`
}

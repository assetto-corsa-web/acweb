// Extracts hours and minutes in 24h format from string in format hh:mi
// and returns them as object. Throws an error if time is invalid.
function parseTime(time){
	var split = time.split(':');
	var hours = parseInt(split[0]);
	var minutes = parseInt(split[1]);

	if(hours < 0 || hours > 23 || minutes < 0 || minutes > 59){
		throw 'Hours and minutes must be within range 00-23 and 00-59';
	}

	return {hours: hours, minutes: minutes};
}

// Returns hours, minutes for given date in format hh:mi.
function getTimeHHMI(date){
	var minutes = date/1000/60;
	var hours = Math.floor(minutes/60);
	minutes -= hours*60;
	hours = _addLeadingZero(hours);
	minutes = _addLeadingZero(minutes);

	return [hours, minutes].join(':');
}

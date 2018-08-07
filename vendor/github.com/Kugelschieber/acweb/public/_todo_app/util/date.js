var _MS_PER_DAY = 1000 * 60 * 60 * 24;

// Returns date in format dd.mm.yyyy hh:mi:ss.
Date.prototype.formatDE = function(withTime){
	var day = _addLeadingZero(this.getDate());
	var month = _addLeadingZero(this.getMonth()+1);
	var year = this.getFullYear();
	var hours = _addLeadingZero(this.getHours());
	var minutes = _addLeadingZero(this.getMinutes());
	var seconds = _addLeadingZero(this.getSeconds());

	return [day, month, year].join('.')+' '+[hours, minutes, seconds].join(':');
};

// Returns date in format dd.mm.yyyy.
Date.prototype.formatDDMMYYYY = function(){
	var day = _addLeadingZero(this.getDate());
	var month = _addLeadingZero(this.getMonth()+1);
	var year = this.getFullYear();

	return [day, month, year].join('.');
};

// Returns date in format yyyy-mm-dd.
Date.prototype.formatYYYYMMDD = function(){
	var day = _addLeadingZero(this.getDate());
	var month = _addLeadingZero(this.getMonth()+1);
	var year = this.getFullYear();

	return [year, month, day].join('-');
};

// Returns date in format hh:mi.
Date.prototype.formatHHMI = function(){
	var hours = _addLeadingZero(this.getHours());
	var minutes = _addLeadingZero(this.getMinutes());

	return [hours, minutes].join(':');
};

// Parses time only from given ISO string.
Date.parseTime = function(time){
	var date = new Date(0);
	var hours = parseInt(time.substring(11, 13));
	var minutes = parseInt(time.substring(14, 16));

	date.setHours(hours);
	date.setMinutes(minutes);

	return date;
};

function _addLeadingZero(time){
	if(time < 10){
		return '0'+time
	}

	return time
}

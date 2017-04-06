// Function to convert € to cent.
// Substitutes ',' with '.'.
// Save to use for user input.
function toCent(amount){
	if(!amount){
		amount = 0;
	}

	amount = amount.toString();
	amount = amount.replace(',', '.');

	var sep = amount.indexOf('.');
	var euro = 0;
	var cent = 0;

	if(sep != -1){
		euro = parseInt(amount.substr(0, sep));

		var centStr = amount.substr(sep+1, sep+3);

		if(centStr.length < 2){
			centStr += '0';
		}

		cent = parseInt(centStr);
	}
	else{
		euro = parseInt(amount);
	}

	amount = euro*100+cent;

	return amount;
}

// Function to convert cent to €.
// Adds '.' to number.
function toEuro(amount){
	if(!amount){
		amount = 0;
	}
	
	amount = parseInt(amount);
	amount /= 100;

	return amount;
}

// Adds taxes (german Ust.) to €.
function addUst(amount, rate){
	return amount*(1+rate);
}

// Returns taxes (german Ust.) in € from brutto.
function getUst(amount, rate){
	return amount-amount/(1+rate);
}

// Returns the amount minus discount.
// If absolute is true, it will be subtracted, else it will be 100%-discount.
function getMinusDiscount(amount, discount, absolute){
	if(absolute){
		return amount-discount;
	}

	return amount*(1.0-discount);
}

// Calculates sum for given attribute in given list.
// Tries to parse floating point values of given attribute.
function sumByAttr(list, attr){
	var sum = 0;

	for(var i = 0; i < list.length; i++){
		sum += parseFloat(list[i][attr]);
	}

	return sum;
}

// Calculates sum for given list of items.
// This will multiply the first attribute with the second and sums them up.
// Tries to parse floating point values of given attributes.
function sumItemsByAttr(list, attr0, attr1){
	var sum = 0;

	for(var i = 0; i < list.length; i++){
		sum += parseFloat(list[i][attr0])*parseFloat(list[i][attr1]);
	}

	return sum;
}

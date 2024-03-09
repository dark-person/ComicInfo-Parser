/**
 * This file contains filename utilities.
 */

/** Function for get basename in absolute path. */
export const basename = function (absPath: string): string {
	const temp = (" " + absPath).slice(1);

	if (temp.split("\\") == undefined) {
		return "";
	}

	const strParts = temp.split("\\").pop();
	if (strParts === undefined) {
		return "";
	}

	if (strParts.split("/") === undefined) {
		return "";
	}

	const strParts2 = strParts.split("/").pop();
	if (strParts2 === undefined) {
		return "";
	}

	return strParts2;
};

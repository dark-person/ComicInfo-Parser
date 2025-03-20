/** This file contains filename utilities. */

/**
 * Function for get basename in absolute path, or empty string if error occurs.
 *
 * Basename is retrieved by delete any prefix up to the last slash ('/') character and return the result.
 *
 * @param absPath absolute path of folder/file
 * @returns basename of absolute path, or empty string if any error occurred
 */
export function basename(absPath: string): string {
    const temp = (" " + absPath).slice(1);

    if (temp.split("\\") === undefined) {
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
}

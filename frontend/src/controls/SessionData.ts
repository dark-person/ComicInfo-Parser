/** Method for export comicinfo cbz. */
export enum ExportMethod {
	/** Only export single cbz file. */
	CBZ_ONLY,
	/** Export cbz file with a folder wrapped. */
	DEFAULT_WRAP_CBZ,
}

export type SessionData = {
	/** Method to export. This will keep save last option until program close. */
	exportMethod: ExportMethod;
};

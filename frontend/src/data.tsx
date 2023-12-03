import { comicinfo } from "../wailsjs/go/models";

/**
 * The data container to be passed among different components.
 */
export class DataPass {
	comicInfo: comicinfo.ComicInfo | undefined = undefined;
	folder: string | undefined = undefined;
}

// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {comicinfo} from '../models';
import {application} from '../models';

export function ExportCbzOnly(arg1:string,arg2:string,arg3:comicinfo.ComicInfo):Promise<string>;

export function ExportCbzWithDefaultWrap(arg1:string,arg2:string,arg3:comicinfo.ComicInfo):Promise<string>;

export function ExportCbzWithWrap(arg1:string,arg2:string,arg3:string,arg4:comicinfo.ComicInfo):Promise<string>;

export function ExportXml(arg1:string,arg2:comicinfo.ComicInfo):Promise<string>;

export function GetAllGenreInput():Promise<application.HistoryResp>;

export function GetAllPublisherInput():Promise<application.HistoryResp>;

export function GetAllTagInput():Promise<application.HistoryResp>;

export function GetComicFolder():Promise<string>;

export function GetComicInfo(arg1:string):Promise<application.ComicInfoResponse>;

export function GetDefaultOutputDirectory(arg1:string):Promise<string>;

export function GetDirectory(arg1:string):Promise<application.DirectoryResp>;

export function OpenFolder(arg1:string):Promise<void>;

export function QuickExportKomga(arg1:string):Promise<string>;

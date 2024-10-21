// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {comicinfo} from '../models';
import {application} from '../models';

export function ExportCbz(arg1:string,arg2:string,arg3:comicinfo.ComicInfo,arg4:boolean):Promise<string>;

export function ExportXml(arg1:string,arg2:comicinfo.ComicInfo):Promise<string>;

export function GetAllGenreInput():Promise<application.HistoryResp>;

export function GetAllPublisherInput():Promise<application.HistoryResp>;

export function GetAllTagInput():Promise<application.HistoryResp>;

export function GetComicInfo(arg1:string):Promise<application.ComicInfoResponse>;

export function GetDirectory():Promise<string>;

export function GetDirectoryWithDefault(arg1:string):Promise<string>;

export function QuickExportKomga(arg1:string):Promise<string>;

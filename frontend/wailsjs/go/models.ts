export namespace application {
	
	export class ComicInfoResponse {
	    ComicInfo?: comicinfo.ComicInfo;
	    ErrorMessage: string;
	
	    static createFrom(source: any = {}) {
	        return new ComicInfoResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ComicInfo = this.convertValues(source["ComicInfo"], comicinfo.ComicInfo);
	        this.ErrorMessage = source["ErrorMessage"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DirectoryResp {
	    SelectedDir: string;
	    ErrMsg: string;
	
	    static createFrom(source: any = {}) {
	        return new DirectoryResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SelectedDir = source["SelectedDir"];
	        this.ErrMsg = source["ErrMsg"];
	    }
	}
	export class HistoryResp {
	    Inputs: string[];
	    ErrorMsg: string;
	
	    static createFrom(source: any = {}) {
	        return new HistoryResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Inputs = source["Inputs"];
	        this.ErrorMsg = source["ErrorMsg"];
	    }
	}

}

export namespace comicinfo {
	
	export enum Manga {
	    Unknown = "Unknown",
	    No = "No",
	    Yes = "Yes",
	    YesAndRightToLeft = "YesAndRightToLeft",
	}
	export enum AgeRating {
	    Unknown = "Unknown",
	    AdultsOnly18 = "Adults Only 18+",
	    EarlyChildhood = "Early Childhood",
	    Everyone = "Everyone",
	    Everyone10Plus = "Everyone 10+",
	    G = "G",
	    KidsToAdults = "Kids to Adults",
	    M = "M",
	    MA15Plus = "MA15+",
	    Mature17Plus = "Mature 17+",
	    PG = "PG",
	    R18Plus = "R18+",
	    RatingPending = "Rating Pending",
	    Teen = "Teen",
	    X18Plus = "X18+",
	}
	export class ComicPageInfo {
	    XMLName: xml.Name;
	    Image: number;
	    Type: string;
	    DoublePage: boolean;
	    ImageSize: number;
	    Key: string;
	    Bookmark: string;
	    ImageWidth: string;
	    ImageHeight: string;
	
	    static createFrom(source: any = {}) {
	        return new ComicPageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.XMLName = this.convertValues(source["XMLName"], xml.Name);
	        this.Image = source["Image"];
	        this.Type = source["Type"];
	        this.DoublePage = source["DoublePage"];
	        this.ImageSize = source["ImageSize"];
	        this.Key = source["Key"];
	        this.Bookmark = source["Bookmark"];
	        this.ImageWidth = source["ImageWidth"];
	        this.ImageHeight = source["ImageHeight"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class EscapedString {
	    InnerXML: string;
	
	    static createFrom(source: any = {}) {
	        return new EscapedString(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.InnerXML = source["InnerXML"];
	    }
	}
	export class ComicInfo {
	    XMLName: xml.Name;
	    Xsi: string;
	    Xsd: string;
	    Title: string;
	    Series: string;
	    Number: string;
	    Count: number;
	    Volume: number;
	    AlternateSeries: string;
	    AlternateNumber: string;
	    AlternateCount: number;
	    StoryArc: string;
	    StoryArcNumber: string;
	    SeriesGroup: string;
	    Summary: EscapedString;
	    Notes: string;
	    Year: number;
	    Month: number;
	    Day: number;
	    Writer: string;
	    Penciller: string;
	    Inker: string;
	    Colorist: string;
	    Letterer: string;
	    CoverArtist: string;
	    Editor: string;
	    Translator: string;
	    Publisher: string;
	    Imprint: string;
	    Genre: string;
	    Tags: string;
	    Web: string;
	    PageCount: number;
	    LanguageISO: string;
	    Format: string;
	    AgeRating: AgeRating;
	    BlackAndWhite: string;
	    Manga: Manga;
	    Characters: string;
	    Teams: string;
	    Locations: string;
	    ScanInformation: string;
	    Pages: ComicPageInfo[];
	    CommunityRating: number;
	    MainCharacterOrTeam: string;
	    Review: string;
	    GTIN: string;
	
	    static createFrom(source: any = {}) {
	        return new ComicInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.XMLName = this.convertValues(source["XMLName"], xml.Name);
	        this.Xsi = source["Xsi"];
	        this.Xsd = source["Xsd"];
	        this.Title = source["Title"];
	        this.Series = source["Series"];
	        this.Number = source["Number"];
	        this.Count = source["Count"];
	        this.Volume = source["Volume"];
	        this.AlternateSeries = source["AlternateSeries"];
	        this.AlternateNumber = source["AlternateNumber"];
	        this.AlternateCount = source["AlternateCount"];
	        this.StoryArc = source["StoryArc"];
	        this.StoryArcNumber = source["StoryArcNumber"];
	        this.SeriesGroup = source["SeriesGroup"];
	        this.Summary = this.convertValues(source["Summary"], EscapedString);
	        this.Notes = source["Notes"];
	        this.Year = source["Year"];
	        this.Month = source["Month"];
	        this.Day = source["Day"];
	        this.Writer = source["Writer"];
	        this.Penciller = source["Penciller"];
	        this.Inker = source["Inker"];
	        this.Colorist = source["Colorist"];
	        this.Letterer = source["Letterer"];
	        this.CoverArtist = source["CoverArtist"];
	        this.Editor = source["Editor"];
	        this.Translator = source["Translator"];
	        this.Publisher = source["Publisher"];
	        this.Imprint = source["Imprint"];
	        this.Genre = source["Genre"];
	        this.Tags = source["Tags"];
	        this.Web = source["Web"];
	        this.PageCount = source["PageCount"];
	        this.LanguageISO = source["LanguageISO"];
	        this.Format = source["Format"];
	        this.AgeRating = source["AgeRating"];
	        this.BlackAndWhite = source["BlackAndWhite"];
	        this.Manga = source["Manga"];
	        this.Characters = source["Characters"];
	        this.Teams = source["Teams"];
	        this.Locations = source["Locations"];
	        this.ScanInformation = source["ScanInformation"];
	        this.Pages = this.convertValues(source["Pages"], ComicPageInfo);
	        this.CommunityRating = source["CommunityRating"];
	        this.MainCharacterOrTeam = source["MainCharacterOrTeam"];
	        this.Review = source["Review"];
	        this.GTIN = source["GTIN"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

export namespace store {
	
	export class AutofillWord {
	    id: number;
	    word: string;
	    category: string;
	
	    static createFrom(source: any = {}) {
	        return new AutofillWord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.word = source["word"];
	        this.category = source["category"];
	    }
	}

}

export namespace xml {
	
	export class Name {
	    Space: string;
	    Local: string;
	
	    static createFrom(source: any = {}) {
	        return new Name(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Space = source["Space"];
	        this.Local = source["Local"];
	    }
	}

}


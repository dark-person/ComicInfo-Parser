export namespace comicinfo {
	
	export class ComicPageInfo {
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
	        this.Image = source["Image"];
	        this.Type = source["Type"];
	        this.DoublePage = source["DoublePage"];
	        this.ImageSize = source["ImageSize"];
	        this.Key = source["Key"];
	        this.Bookmark = source["Bookmark"];
	        this.ImageWidth = source["ImageWidth"];
	        this.ImageHeight = source["ImageHeight"];
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
	    AgeRating: string;
	    BlackAndWhite: string;
	    Manga: string;
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
		    if (a.slice) {
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


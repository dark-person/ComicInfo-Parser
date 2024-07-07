// React
import { ChangeEvent, useState } from "react";

// React Component
import { Button, Col, Form, Row } from "react-bootstrap";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";

// Project Specified Component
import { AgeRatingSelect, MangaSelect } from "../components/EnumSelect";
import { TagsArea } from "../components/Tags";
import { basename } from "../filename";
import FormDateRow from "../forms/FormDateRow";
import FormRow from "../forms/FormRow";
import FormSelectRow from "../forms/FormSelectRow";
import OptionFormRow from "../forms/OptionFormRow";

import { GetAllGenreInput } from "../../wailsjs/go/main/App";
import { comicinfo } from "../../wailsjs/go/models";

/** Props Interface for InputPanel */
type InputProps = {
	/** The comic info object. Accept undefined value. */
	comicInfo: comicinfo.ComicInfo | undefined;

	/** The folder name for reference. */
	folderName?: string;

	/** The function to change display panel to export panel. */
	exportFunc: () => void;

	/** The info Setter. This function should be doing setting value, but no verification. */
	infoSetter: (field: string, value: string | number) => void;
};

/** Props Interface for Metadata. */
type MetadataProps = {
	/** The comic info object. Accept undefined value. */
	comicInfo: comicinfo.ComicInfo | undefined;

	/** The method called when input field value is changed. */
	dataHandler: (e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLSelectElement>) => void;
};

/**
 * The interface for show/edit book metadata.
 * @returns JSX Element
 */
function BookMetadata({ comicInfo: info, dataHandler }: Readonly<MetadataProps>) {
	return (
		<div>
			<Form>
				<FormRow title={"Title"} value={info?.Title} onChange={dataHandler} />
				<FormRow title={"Summary"} value={info?.Summary.InnerXML} textareaRow={3} onChange={dataHandler} />
				<FormRow title={"Number"} value={info?.Number} onChange={dataHandler} />
				<FormDateRow
					title={"Year/Month/Day"}
					year={info?.Year}
					month={info?.Month}
					day={info?.Day}
					onYearChange={dataHandler}
					onSelectChange={dataHandler}
				/>
				<FormRow title={"Web"} value={info?.Web} onChange={dataHandler} />
				<FormRow title={"GTIN"} value={info?.GTIN} onChange={dataHandler} />
			</Form>
		</div>
	);
}

/**
 * The interface for show/edit creator metadata.
 * @returns JSX Element
 */
function CreatorMetadata({ comicInfo: info, dataHandler }: Readonly<MetadataProps>) {
	return (
		<div>
			<Form>
				<FormRow title={"Writer"} value={info?.Writer} onChange={dataHandler} />
				<FormRow title={"Penciller"} value={info?.Penciller} onChange={dataHandler} />
				<FormRow title={"Inker"} value={info?.Inker} onChange={dataHandler} />
				<FormRow title={"Colorist"} value={info?.Colorist} onChange={dataHandler} />
				<FormRow title={"Letterer"} value={info?.Letterer} onChange={dataHandler} />
				<FormRow title={"CoverArtist"} value={info?.CoverArtist} onChange={dataHandler} />
				<FormRow title={"Editor"} value={info?.Editor} onChange={dataHandler} />
				<FormRow title={"Translator"} value={info?.Translator} onChange={dataHandler} />
				<FormRow title={"Publisher"} value={info?.Publisher} onChange={dataHandler} />
			</Form>
		</div>
	);
}

/** The user interface for show/edit Series MetaData. */
function SeriesMetadata({ comicInfo, infoSetter }: Readonly<TagMetadataProps>) {
	/**
	 * Handler for all input field in this panel.
	 * This method will use <code>infoSetter</code> as core,
	 * and apply change to comicInfo content.
	 * <p>
	 * This method will try to find the field is number input first,
	 * if field name is number/related, then it will call Number() method before set to ComicInfo
	 *
	 * @param e the event object, for identify target element
	 * @returns void
	 */
	function handleChanges(e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLSelectElement>) {
		console.log(e.target.title, e.target.value, e.target.type);

		// Identify Number
		if (e.target.type === "number") {
			infoSetter(e.target.title, Number(e.target.value));
			return;
		}

		// Normal Cases
		infoSetter(e.target.title, e.target.value);
	}

	return (
		<div>
			<Form>
				<FormRow title={"Series"} value={comicInfo?.Series} onChange={handleChanges} />
				<FormRow title={"Volume"} value={comicInfo?.Volume} onChange={handleChanges} />
				<FormRow title={"Count"} value={comicInfo?.Count} onChange={handleChanges} />
				<FormSelectRow
					title={"AgeRating"}
					selectElement={<AgeRatingSelect value={comicInfo?.AgeRating} onChange={handleChanges} />}
				/>
				<FormSelectRow
					title={"Manga"}
					selectElement={<MangaSelect value={comicInfo?.Manga} onChange={handleChanges} />}
				/>
				<OptionFormRow
					title={"Genre"}
					value={comicInfo?.Genre}
					setValue={(val) => infoSetter("Genre", val)}
					getDefaultOpt={GetAllGenreInput}
				/>
				<FormRow title={"LanguageISO"} value={comicInfo?.LanguageISO} onChange={handleChanges} />
			</Form>
		</div>
	);
}

/** The user interface for show/edit Collection & ReadList MetaData. */
function MiscMetadata({ comicInfo, dataHandler }: Readonly<MetadataProps>) {
	return (
		<div>
			<Form>
				<FormRow title={"SeriesGroup"} value={comicInfo?.SeriesGroup} onChange={dataHandler} />
				<FormRow title={"AlternateSeries"} value={comicInfo?.AlternateSeries} onChange={dataHandler} />
				<FormRow title={"AlternateNumber "} value={comicInfo?.AlternateNumber} onChange={dataHandler} />
				<FormRow title={"AlternateCount"} value={comicInfo?.AlternateCount} onChange={dataHandler} />
				<FormRow title={"StoryArc"} value={comicInfo?.StoryArc} onChange={dataHandler} />
				<FormRow title={"StoryArcNumber"} value={comicInfo?.StoryArcNumber} onChange={dataHandler} />
			</Form>
		</div>
	);
}

/** The Props for Tag Metadata Component. */
type TagMetadataProps = {
	/** The comic info object. Accept undefined value. */
	comicInfo: comicinfo.ComicInfo | undefined;

	/** The info Setter. This function should be doing setting value, but no verification. */
	infoSetter: (field: string, value: string | number) => void;
};

/**
 * The interface for show/edit tags metadata.
 * @returns JSX Element
 */
function TagMetadata({ comicInfo: info, infoSetter }: Readonly<TagMetadataProps>) {
	/** Hooks of tag that to be added. Only allow single tag to be added at a time. */
	const [singleTag, setSingleTag] = useState<string>("");

	/** Function for handing enter key pressed when entering custom tags. */
	const handleKeyDown = (event: React.KeyboardEvent) => {
		if (event.key === "Enter") {
			event.preventDefault();
			handleAdd();
		}
	};

	/**
	 * Function to handle the textfield of tag to be added.
	 * @param e the react event
	 */
	function handleChanges(e: ChangeEvent<HTMLInputElement>) {
		setSingleTag(e.target.value);
	}

	/** Function for handling add button click */
	function handleAdd() {
		// Prevent Empty tags
		if (singleTag === "") {
			return;
		}

		// Retrieve Tags
		let temp = info?.Tags;

		// Append Tags to comic info
		if (temp === "" || temp === undefined) {
			temp = singleTag;
		} else {
			temp += ", " + singleTag;
		}

		// Apply Change to ComicInfo
		infoSetter("Tags", temp);

		// Empty Tags in textfield
		setSingleTag("");
	}

	/** Function for handling delete button click */
	function handleDelete(id: number) {
		if (info?.Tags === undefined) {
			return;
		}

		// Parse tags to array of strings
		let array = info.Tags.split(",");

		// Remove tag by index
		array.splice(id, 1);

		// Concat array of string to string
		let str = array.join(", ");

		// Set by info setter
		infoSetter("Tags", str);
	}

	return (
		<div>
			<Form>
				{/* A Text Area for holding lines of tags. */}
				<Row>
					<Col sm={2} className="mt-1">
						{"Tags"}
					</Col>
					<Col sm={9}>
						<TagsArea rawTags={info?.Tags} handleDelete={handleDelete} />
					</Col>
				</Row>

				{/* Empty Rows for margin */}
				<Row className="mb-3" />

				<Row>
					{/* Empty Column */}
					<Col sm={2} className="mt-1">
						Add Custom Tag
					</Col>

					{/* Column of adding tags */}
					<Col sm={8}>
						<Form.Control value={singleTag} onChange={handleChanges} onKeyDown={handleKeyDown} />
					</Col>

					{/* Column of add button */}
					<Col sm={1}>
						<Button variant="outline-info" onClick={handleAdd}>
							Add
						</Button>
					</Col>
				</Row>
			</Form>
		</div>
	);
}

/**
 * The panel for input/edit content of ComicInfo.xml
 * @returns JSX Element
 */
export default function InputPanel({ comicInfo, folderName, exportFunc, infoSetter }: Readonly<InputProps>) {
	/**
	 * Handler for all input field in this panel.
	 * This method will use <code>infoSetter</code> as core,
	 * and apply change to comicInfo content.
	 * <p>
	 * This method will try to find the field is number input first,
	 * if field name is number/related, then it will call Number() method before set to ComicInfo
	 *
	 * @param e the event object, for identify target element
	 * @returns void
	 */
	function handleChanges(e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLSelectElement>) {
		console.log(e.target.title, e.target.value, e.target.type);

		// Identify Number & Year
		if (e.target.type === "number" || e.target.type === "Year") {
			infoSetter(e.target.title, Number(e.target.value));
			return;
		}

		// Identify Month & Day
		if (e.target.title === "Month" || e.target.title === "Day") {
			infoSetter(e.target.title, Number(e.target.value));
			return;
		}

		// Normal Cases
		infoSetter(e.target.title, e.target.value);
	}

	return (
		<div id="Input-Panel" className="mt-5">
			<h5 className="mb-4">Modify ComicInfo.xml</h5>

			{/* Component for showing folder name (with basename only) */}
			<FormRow
				title={"Folder Name"}
				titleClass="fst-italic"
				value={folderName != undefined ? basename(folderName) : "(N/A)"}
				disabled
			/>

			{/* The Tabs Group to display metadata. */}
			<Tabs defaultActiveKey="Main" id="uncontrolled-tab-example" className="mb-3">
				<Tab eventKey="Main" title="Book Metadata">
					<BookMetadata comicInfo={comicInfo} dataHandler={handleChanges} />
				</Tab>

				<Tab eventKey="Creator" title="Creator">
					<CreatorMetadata comicInfo={comicInfo} dataHandler={handleChanges} />
				</Tab>

				<Tab eventKey="Tags" title="Tags">
					<TagMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
				</Tab>

				<Tab eventKey="Series" title="Series">
					<SeriesMetadata comicInfo={comicInfo} infoSetter={infoSetter} />
				</Tab>

				<Tab eventKey="Misc" title="Collection & ReadList">
					<MiscMetadata comicInfo={comicInfo} dataHandler={handleChanges} />
				</Tab>
			</Tabs>

			{/* The button that will always at the bottom of screen. Should ensure there has enough space */}
			<div className="fixed-bottom mb-3">
				<Button variant="success" className="mx-2 " id="btn-export-cbz" onClick={exportFunc}>
					Export to .cbz
				</Button>
			</div>
		</div>
	);
}

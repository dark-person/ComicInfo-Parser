// React
import { ChangeEvent, useState } from "react";

// React Component
import { Button, Col, Form, Row } from "react-bootstrap";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";

// Project Specified Component
import { comicinfo } from "../../wailsjs/go/models";
import { TagsArea } from "../components/Tags";
import { FormDateRow, FormRow } from "../formRow";

/** Props Interface for InputPanel */
type InputProps = {
	/** The comic info object. Accept undefined value. */
	comicInfo: comicinfo.ComicInfo | undefined;

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
function BookMetadata({ comicInfo: info, dataHandler }: MetadataProps) {
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
function CreatorMetadata({ comicInfo: info, dataHandler }: MetadataProps) {
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
function SeriesMetadata({ comicInfo, dataHandler }: MetadataProps) {
	return (
		<div>
			<Form>
				<FormRow title={"Series"} value={comicInfo?.Series} onChange={dataHandler} />
				<FormRow title={"Volume"} value={comicInfo?.Volume} onChange={dataHandler} />
				<FormRow title={"Count"} value={comicInfo?.Count} onChange={dataHandler} />
				<FormRow title={"AgeRating"} value={comicInfo?.AgeRating} onChange={dataHandler} />
				<FormRow title={"Manga"} value={comicInfo?.Manga} onChange={dataHandler} />
				<FormRow title={"Genre"} value={comicInfo?.Genre} onChange={dataHandler} />
				<FormRow title={"LanguageISO"} value={comicInfo?.LanguageISO} onChange={dataHandler} />
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
function TagMetadata({ comicInfo: info, infoSetter }: TagMetadataProps) {
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
		if (info === undefined || info.Tags === undefined) {
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
				{/* <FormRow title={"Tags"} textareaRow={10} value={info?.Tags} onChange={dataHandler} /> */}
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
export default function InputPanel({ comicInfo, exportFunc, infoSetter }: InputProps) {
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
					<SeriesMetadata comicInfo={comicInfo} dataHandler={handleChanges} />
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

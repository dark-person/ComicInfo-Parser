// React
import { ChangeEvent } from "react";

// React Component
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import { Button, Form } from "react-bootstrap";

// Project Specified Component
import { FormDateRow, FormRow } from "../formRow";
import { comicinfo } from "../../wailsjs/go/models";

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

/**
 * The interface for show/edit tags metadata.
 * @returns JSX Element
 */
function TagMetadata({ comicInfo: info, dataHandler }: MetadataProps) {
	return (
		<div>
			<Form>
				{/* A Text Area for holding lines of tags. */}
				<FormRow title={"Tags"} textareaRow={10} value={info?.Tags} onChange={dataHandler} />
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
					<TagMetadata comicInfo={comicInfo} dataHandler={handleChanges} />
				</Tab>
			</Tabs>

			{/* The button that will always at the bottom of screen. Should ensure there has enough space */}
			<div className="fixed-bottom mb-3">
				<Button variant="outline-success" className="mx-2 " id="btn-export-cbz" onClick={exportFunc}>
					Export to .cbz
				</Button>
			</div>
		</div>
	);
}

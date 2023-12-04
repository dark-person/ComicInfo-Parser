// React Component
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import { Button, Form } from "react-bootstrap";

// Project Specified Component
import { FormDateRow, FormRow } from "../formRow";
import { comicinfo } from "../../wailsjs/go/models";

/** Props Interface for InputPanel */
type InputProps = {
	comicInfo: comicinfo.ComicInfo | undefined;
	exportFunc: () => void;
	// returnFunc: (event: React.MouseEvent) => void;
};

/** Props Interface for Metadata */
type MetadataProps = {
	comicInfo: comicinfo.ComicInfo | undefined;
};

/**
 * The interface for show/edit book metadata.
 * @returns JSX Element
 */
function BookMetadata({ comicInfo: info }: MetadataProps) {
	return (
		<div>
			<Form>
				<FormRow title={"Title"} value={info?.Title} disabled />
				<FormRow
					title={"Summary"}
					value={info?.Summary.InnerXML}
					textareaRow={3}
					disabled
				/>
				<FormRow title={"Number"} inputType="number" value={info?.Number} disabled />
				<FormDateRow
					title={"Year/Month/Day"}
					year={info?.Year}
					month={info?.Month}
					day={info?.Day}
					disabled
				/>
				<FormRow title={"Web"} value={info?.Web} disabled />
				<FormRow title={"GTIN"} value={info?.GTIN} disabled />
			</Form>
		</div>
	);
}

/**
 * The interface for show/edit creator metadata.
 * @returns JSX Element
 */
function CreatorMetadata({ comicInfo: info }: MetadataProps) {
	return (
		<div>
			<Form>
				<FormRow title={"Writer"} value={info?.Writer} disabled />
				<FormRow title={"Penciller"} value={info?.Penciller} disabled />
				<FormRow title={"Inker"} value={info?.Inker} disabled />
				<FormRow title={"Colorist"} value={info?.Colorist} disabled />
				<FormRow title={"Letterer"} value={info?.Letterer} disabled />
				<FormRow title={"CoverArtist"} value={info?.CoverArtist} disabled />
				<FormRow title={"Editor"} value={info?.Editor} disabled />
				<FormRow title={"Translator"} value={info?.Translator} disabled />
				<FormRow title={"Publisher"} value={info?.Publisher} disabled />
			</Form>
		</div>
	);
}

/**
 * The interface for show/edit tags metadata.
 * @returns JSX Element
 */
function TagMetadata({ comicInfo: info }: MetadataProps) {
	return (
		<div>
			<Form>
				<FormRow title={"Tag"} textareaRow={10} value={info?.Tags} disabled />
			</Form>
		</div>
	);
}

/**
 * The panel for input/edit content of ComicInfo.xml
 * @returns JSX Element
 */
export default function InputPanel({ comicInfo, exportFunc }: InputProps) {
	return (
		<div id="Input-Panel" className="mt-5">
			<h5 className="mb-4">Modify ComicInfo.xml</h5>

			<Tabs defaultActiveKey="Main" id="uncontrolled-tab-example" className="mb-3">
				<Tab eventKey="Main" title="Book Metadata">
					<BookMetadata comicInfo={comicInfo} />
				</Tab>

				<Tab eventKey="Creator" title="Creator">
					<CreatorMetadata comicInfo={comicInfo} />
				</Tab>
				<Tab eventKey="Tags" title="Tags">
					<TagMetadata comicInfo={comicInfo} />
				</Tab>
			</Tabs>

			<div className="fixed-bottom mb-3">
				<Button
					variant="outline-success"
					className="mx-2 "
					id="btn-export-cbz"
					onClick={exportFunc}>
					Export to .cbz
				</Button>
			</div>
		</div>
	);
}

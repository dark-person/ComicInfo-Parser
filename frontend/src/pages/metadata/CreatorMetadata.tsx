// React
import { ChangeEvent } from "react";

// React Component
import { Form } from "react-bootstrap";

// Project Specified Component
import FormRow from "../../forms/FormRow";
import OptionFormRow from "../../forms/OptionFormRow";
import { MetadataProps } from "./MetadataProps";

// Wails binding
import { GetAllPublisherInput } from "../../../wailsjs/go/main/App";

/**
 * The interface for show/edit creator metadata.
 * @returns JSX Element
 */
export default function CreatorMetadata({ comicInfo: info, infoSetter }: Readonly<MetadataProps>) {
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
		<div>
			<Form>
				<FormRow title={"Writer"} value={info?.Writer} onChange={handleChanges} />
				<FormRow title={"Penciller"} value={info?.Penciller} onChange={handleChanges} />
				<FormRow title={"Inker"} value={info?.Inker} onChange={handleChanges} />
				<FormRow title={"Colorist"} value={info?.Colorist} onChange={handleChanges} />
				<FormRow title={"Letterer"} value={info?.Letterer} onChange={handleChanges} />
				<FormRow title={"CoverArtist"} value={info?.CoverArtist} onChange={handleChanges} />
				<FormRow title={"Editor"} value={info?.Editor} onChange={handleChanges} />
				<FormRow title={"Translator"} value={info?.Translator} onChange={handleChanges} />
				<OptionFormRow
					title={"Publisher"}
					value={info?.Publisher}
					setValue={(val) => infoSetter("Publisher", val)}
					getDefaultOpt={GetAllPublisherInput}
				/>
			</Form>
		</div>
	);
}

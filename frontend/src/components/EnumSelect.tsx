import { Form } from "react-bootstrap";
import { comicinfo } from "../../wailsjs/go/models";

/** Props for `EnumOptions` */
type OptionProps = {
	/** Both displayed value & actual value for `<option>`. */
	value: string | undefined;
};

/** JSX element for `<option>`, designed to hold enum value. */
function EnumOptions({ value }: Readonly<OptionProps>) {
	return <option value={value}>{value}</option>;
}

/** Props for select box. */
type EnumSelectProps = {
	/** Current value of input field. */
	value: string | undefined;
	/** Handle value change of input field. */
	onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
};

/** Form Select Element specified for `Manga` enum. */
export function MangaSelect({ value, onChange }: Readonly<EnumSelectProps>) {
	return (
		<Form.Select value={value} title="Manga" onChange={onChange}>
			<option value={""}></option>
			{Object.values(comicinfo.Manga).map((item, idx) => (
				<EnumOptions value={item} key={"manga-opt-" + idx} />
			))}
		</Form.Select>
	);
}

/** Form Select Element specified for `AgeRating` enum. */
export function AgeRatingSelect({ value, onChange }: Readonly<EnumSelectProps>) {
	return (
		<Form.Select value={value} title="AgeRating" onChange={onChange}>
			<option value={""}></option>
			{Object.values(comicinfo.AgeRating).map((item, idx) => (
				<EnumOptions value={item} key={"age-opt-" + idx} />
			))}
		</Form.Select>
	);
}

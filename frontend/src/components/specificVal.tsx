import { Form } from "react-bootstrap";
import { comicinfo } from "../../wailsjs/go/models";

type props = {
	value: string | undefined;
};

function MangaOptions({ value }: props) {
	return <option value={value}>{value}</option>;
}

type MangaSelectProps = {
	value: string | undefined;
	/** Handle value change of input field. */
	onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
};

export function MangaSelect({ value, onChange }: MangaSelectProps) {
	return (
		<Form.Select value={value} title="Manga" onChange={onChange}>
			<option value={""}></option>
			{Object.values(comicinfo.Manga).map((item) => (
				<MangaOptions value={item} />
			))}
		</Form.Select>
	);
}

export function AgeRatingSelect({ value, onChange }: MangaSelectProps) {
	return (
		<Form.Select value={value} title="AgeRating" onChange={onChange}>
			<option value={""}></option>
			{Object.values(comicinfo.AgeRating).map((item) => (
				<MangaOptions value={item} />
			))}
		</Form.Select>
	);
}

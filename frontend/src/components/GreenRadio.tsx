import { ChangeEvent } from "react";
import { Form } from "react-bootstrap";
import "./GreenRadio.css";

type GreenRadioProps = {
	/** HTML element ID. */
	id?: string;
	/** Radio button group usage. */
	name?: string;
	/** String to display as label. */
	label?: string;
	/** Additional HTML class to append. */
	className?: string;
	/** Determine radio button is clicked or not. */
	checked?: boolean;
	/** Onchange handler for radio button. */
	onChange?: (evt: ChangeEvent<HTMLInputElement>) => void;
};

/** Component of a green radio button. */
export default function GreenRadio({ id, name, label, className, checked, onChange }: Readonly<GreenRadioProps>) {
	return (
		<Form.Check
			type={"radio"}
			label={label}
			name={name}
			id={id}
			className={"text-start green mt-1 mb-2 cursor-pointer " + className}
			checked={checked}
			onChange={onChange}
		/>
	);
}

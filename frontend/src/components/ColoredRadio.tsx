import { ChangeEvent } from "react";
import { Form } from "react-bootstrap";
import "./ColoredRadio.css";

type supportedColors = "dark-red" | "dark-orange" | "dark-yellow" | "dark-green" | "dark-blue";

type ColoredRadioProps = {
	/** HTML element ID. */
	id?: string;
	/** Radio button group usage. */
	name?: string;
	/** Color to be used for radio.*/
	color?: supportedColors;
	/** String to display as label. */
	label?: string;
	/** Additional HTML class to append. */
	className?: string;
	/** Determine radio button is clicked or not. */
	checked?: boolean;
	/** Onchange handler for radio button. */
	onChange?: (evt: ChangeEvent<HTMLInputElement>) => void;
};

/** Component of a colored radio button. */
export default function ColoredRadio(props: Readonly<ColoredRadioProps>) {
	return (
		<Form.Check
			type={"radio"}
			label={props.label}
			name={props.name}
			id={props.id}
			className={`text-start mt-1 mb-2 cursor-pointer hide-focus ${props.color} ${props.className}`}
			checked={props.checked}
			onChange={props.onChange}
		/>
	);
}

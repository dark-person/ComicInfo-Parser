// React Component
import Form from "react-bootstrap/Form";
import { Row, Col } from "react-bootstrap";

/**
 * Get an array that start with min value and end with max value.
 * @param min the minimum value of the array
 * @param max the maximum value of the array
 * @returns array of numbers, value between min and max
 */
function getArray(min: number, max: number): Array<number> {
	return Array.from({ length: max - min + 1 }, (_, i) => i + min);
}

/**
 * Get a select input, provide option from min ~ max. Also provide an option of empty value.
 * @param min the minimum value to be shown in select menu.
 * @param max the maximum value to be shown in select menu.
 * @param value the current value that selected
 * @param disabled disable this element
 * @returns JSX.Element of input select
 */
export function RangeSelect(props: {
	min: number;
	max: number;
	value: number | undefined;
	disabled: boolean | undefined;
}) {
	return (
		<Form.Select
			aria-label="Default select example"
			value={props.value}
			disabled={props.disabled}>
			<option value={undefined} key={"val-undefined"}></option>
			{getArray(props.min, props.max).map((num) => (
				<option value={num} key={"val-" + num.toString()}>
					{num}
				</option>
			))}
		</Form.Select>
	);
}

/**
 * Create a uniform Form.Group Element as Row.
 * @param title the title/label of this input group.
 * @param inputType the type of input, same with HTML input type
 * @param value the current inputted value
 * @param textareaRow the number of row of textarea
 * @param disabled determines whether the input is disabled
 * @returns A Row Element, Contains one input group with label.
 */
export function FormRow(props: {
	title: string;
	inputType?: string;
	value?: string;
	textareaRow?: number | undefined;
	disabled?: boolean;
}) {
	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2">
				{props.title}
			</Form.Label>
			<Col sm="9">
				<Form.Control
					as={props.textareaRow != undefined ? "textarea" : undefined}
					type={props.inputType}
					value={props.value}
					rows={props.textareaRow != undefined ? props.textareaRow : 1}
					disabled={props.disabled}
				/>
			</Col>
		</Form.Group>
	);
}

/**
 * Create a uniform Form.Group Element as Row. This element specify designed for input date.
 * @param title the title/label of this input group.
 * @param year the value of year
 * @param month the value of month
 * @param day the value of day
 * @param disabled determines whether the input is disabled
 * @returns A Row Element, Contains one input group with label.
 */
export function FormDateRow(props: {
	title: string;
	year?: number;
	month?: number;
	day?: number;
	disabled?: boolean;
}) {
	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2">
				{props.title}
			</Form.Label>

			<Col sm="3">
				<Form.Control
					type="number"
					max="9999"
					value={props.year === 0 ? "" : props.year}
					disabled={props.disabled}
				/>
			</Col>

			<Col sm="3">
				<RangeSelect min={1} max={12} value={props.month} disabled={props.disabled} />
			</Col>

			<Col sm="3">
				<RangeSelect min={1} max={31} value={props.day} disabled={props.disabled} />
			</Col>
		</Form.Group>
	);
}

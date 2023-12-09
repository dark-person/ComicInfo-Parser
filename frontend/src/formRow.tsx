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

/** Props for Range Select */
type RangeSelectProps = {
	/** the minimum value to be shown in select menu. */
	min: number;
	/** the maximum value to be shown in select menu. */
	max: number;
	/**  the current value that selected */
	value: number | undefined;
	/** determines whether the input is disabled */
	disabled: boolean | undefined;
};

/**
 * Get a select input, provide option from min ~ max. Also provide an option of empty value.
 * @returns JSX.Element of input select
 */
export function RangeSelect({ min, max, value, disabled }: RangeSelectProps) {
	return (
		<Form.Select aria-label="Default select example" value={value} disabled={disabled}>
			<option value={undefined} key={"val-undefined"}></option>
			{getArray(min, max).map((num) => (
				<option value={num} key={"val-" + num.toString()}>
					{num}
				</option>
			))}
		</Form.Select>
	);
}

/** Props for FormRow */
type FormRowProps = {
	/** the title/label of this input group */
	title: string;
	/** the type of input, same with HTML input type */
	inputType?: string;
	/** current inputted value */
	value?: string;
	/** number of row of textarea */
	textareaRow?: number | undefined;
	/** determines whether the input is disabled */
	disabled?: boolean;

	onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

/**
 * Create a uniform Form.Group Element as Row.
 * @returns A Row Element, Contains one input group with label.
 */
export function FormRow({
	title,
	inputType,
	value,
	textareaRow,
	disabled,
	onChange,
}: FormRowProps) {
	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2">
				{title}
			</Form.Label>
			<Col sm="9">
				<Form.Control
					as={textareaRow != undefined ? "textarea" : undefined}
					type={inputType}
					value={value}
					title={title}
					onChange={onChange}
					rows={textareaRow != undefined ? textareaRow : 1}
					disabled={disabled}
				/>
			</Col>
		</Form.Group>
	);
}

/** Props for FormDateRow */
type FormDateRowProps = {
	/** the title/label of this input group. */
	title: string;
	/** the value of year */
	year?: number;
	/** the value of month */
	month?: number;
	/** the value of day */
	day?: number;
	/** determines whether the input is disabled */
	disabled?: boolean;
};

/**
 * Create a uniform Form.Group Element as Row. This element specify designed for input date.
 * @returns A Row Element, Contains one input group with label. Input group contains three input field for year, month, day.
 */
export function FormDateRow({ title, year, month, day, disabled }: FormDateRowProps) {
	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2">
				{title}
			</Form.Label>

			<Col sm="3">
				<Form.Control
					type="number"
					max="9999"
					value={year === 0 ? "" : year}
					disabled={disabled}
				/>
			</Col>

			<Col sm="3">
				<RangeSelect min={1} max={12} value={month} disabled={disabled} />
			</Col>

			<Col sm="3">
				<RangeSelect min={1} max={31} value={day} disabled={disabled} />
			</Col>
		</Form.Group>
	);
}

// React Component
import { Col, Row } from "react-bootstrap";
import Form from "react-bootstrap/Form";

import RangeSelect from "./RangeSelect";

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
	/** Handler when Year field has changed */
	onYearChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
	/** Handler when select HTML element changed */
	onSelectChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
};

/**
 * Create a uniform Form.Group Element as Row. This element specify designed for input date.
 * @returns A Row Element, Contains one input group with label. Input group contains three input field for year, month, day.
 */
export default function FormDateRow({
	title,
	year,
	month,
	day,
	disabled,
	onYearChange,
	onSelectChange,
}: Readonly<FormDateRowProps>) {
	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2">
				{title}
			</Form.Label>

			{/* Year Field */}
			<Col sm="3">
				<Form.Control
					type="number"
					max="9999"
					title="Year"
					value={year === 0 ? "" : year}
					disabled={disabled}
					onChange={onYearChange}
				/>
			</Col>

			{/* Month Field */}
			<Col sm="3">
				<RangeSelect
					min={1}
					max={12}
					title="Month"
					value={month}
					disabled={disabled}
					onChange={onSelectChange}
				/>
			</Col>

			{/* Day Field */}
			<Col sm="3">
				<RangeSelect min={1} max={31} title="Day" value={day} disabled={disabled} onChange={onSelectChange} />
			</Col>
		</Form.Group>
	);
}

// React Component
import { Col, Row } from "react-bootstrap";
import Form from "react-bootstrap/Form";

/** Props for FormRow */
type FormRowProps = {
	/** the title/label of this input group */
	title: string;
	/** Class name for label. Can be used in styling. */
	titleClass?: string;
	/** the type of input, same with HTML input type */
	inputType?: string;
	/** current inputted value */
	value?: string | number;
	/** number of row of textarea */
	textareaRow?: number;
	/** determines whether the input is disabled */
	disabled?: boolean;
	/** Handle value change of input field. */
	onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

/**
 * Create a uniform Form.Group Element as Row.
 *
 * There has some special handling for `number` values:
 * - When `value == 0`, display empty string instead of `0`
 * - input type of this element will force to `number`
 *
 * @returns A Row Element, Contains one input group with label.
 */
export default function FormRow({
	title,
	titleClass,
	inputType,
	value,
	textareaRow,
	disabled,
	onChange,
}: Readonly<FormRowProps>) {
	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2" className={titleClass ?? ""}>
				{title}
			</Form.Label>
			<Col sm="9">
				<Form.Control
					as={textareaRow !== undefined ? "textarea" : undefined}
					type={typeof value == "number" ? "number" : inputType}
					value={typeof value == "number" && value === 0 ? "" : value}
					title={title}
					onChange={onChange}
					rows={textareaRow ?? 1}
					disabled={disabled}
				/>
			</Col>
		</Form.Group>
	);
}

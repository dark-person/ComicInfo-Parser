// React
import { ChangeEventHandler, useEffect, useState } from "react";

// React Component
import { Col, Row } from "react-bootstrap";
import Form from "react-bootstrap/Form";

import { ActionMeta, GroupBase, MultiValue, StylesConfig } from "react-select";
import CreatableSelect from "react-select/creatable";

import { main } from "../wailsjs/go/models";

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
	/** The title of RangeSelect. It will be used to identify changed field. */
	title: string;
	/** The OnChange Handler, must be used with value by react hook */
	onChange: ChangeEventHandler<HTMLSelectElement>;
};

/**
 * Get a select input, provide option from min ~ max. Also provide an option of empty value.
 * @returns JSX.Element of input select
 */
export function RangeSelect({ min, max, value, title, disabled, onChange }: RangeSelectProps) {
	return (
		<Form.Select
			aria-label="Default select example"
			value={value}
			title={title}
			disabled={disabled}
			onChange={onChange}>
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
	/** Class name for label. Can be used in styling. */
	titleClass?: string;
	/** the type of input, same with HTML input type */
	inputType?: string;
	/** current inputted value */
	value?: string | number;
	/** number of row of textarea */
	textareaRow?: number | undefined;
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
export function FormRow({ title, titleClass, inputType, value, textareaRow, disabled, onChange }: FormRowProps) {
	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2" className={titleClass != undefined ? titleClass : ""}>
				{title}
			</Form.Label>
			<Col sm="9">
				<Form.Control
					as={textareaRow != undefined ? "textarea" : undefined}
					type={typeof value == "number" ? "number" : inputType}
					value={typeof value == "number" && value == 0 ? "" : value}
					title={title}
					onChange={onChange}
					rows={textareaRow != undefined ? textareaRow : 1}
					disabled={disabled}
				/>
			</Col>
		</Form.Group>
	);
}

/** Props for `OptionFormRow`. */
type OptionFormRowProps = {
	/** the title/label of this input group */
	title: string;
	/** Class name for label. Can be used in styling. */
	titleClass?: string;
	/** current inputted value */
	value?: string;
	/** determines whether the input is disabled */
	disabled?: boolean;
	/** Function to set value back to primary comicinfo object. */
	setValue?: (value: string) => void;
	/** Function to get default values from wails backend. */
	getDefaultOpt: () => Promise<main.HistoryResp>;
};

/** Create a uniform Form.Group Element as Row, which contains select component that allow create new options. */
export function OptionFormRow({ title, titleClass, value, disabled, setValue, getDefaultOpt }: OptionFormRowProps) {
	/** Interface for react-select option. */
	interface SelectOption {
		label: string; // Necessary field
		value: string; // Necessary field
	}

	/** Options for react-select, which contains no options (i.e. empty). */
	const emptyOption: MultiValue<SelectOption> = [];

	/** Default option for react-select component. */
	const defaultOptions: SelectOption[] = [
		{ label: "hello", value: "hello" },
		{ label: "world", value: "world" },
	];

	/** Options that used as default options. */
	const [options, setOptions] = useState<SelectOption[]>(defaultOptions);

	/**
	 * Convert values from react-select to single string, joined by comma character.
	 * @param opts options that retrieved from CreatableSelect
	 * @returns string of values, joined by ',' character
	 */
	function concatOptions(opts: MultiValue<SelectOption>): string {
		let simpleOptions = opts.map((item) => item.value);
		return simpleOptions.join(",");
	}

	/**
	 * Convert string (contain ',' or not), to MultiValue that is accepted by react-select components.
	 * @param opt string to be converted to MultiValue
	 * @returns
	 * MultiValue that converted by string, separated by comma character.
	 * Both `label` & `value` in `SelectOption` has same string value.
	 */
	function convert(opt?: string): MultiValue<SelectOption> {
		// Prevent undefined
		if (opt === undefined || opt === "") {
			return [];
		}

		// Split options into string array
		let splitOpts = opt.split(",");

		// Convert to Multiple Values
		return splitOpts.map((item) => ({ label: item, value: item }));
	}

	// Init components
	useEffect(() => {
		getOptions();
	}, []);

	/** Get option for wails binding, which value came from database. */
	function getOptions() {
		// Get data from wails binding
		getDefaultOpt().then((response) => {
			console.log("[getOptions] " + JSON.stringify(response, null, 4));

			if (response.ErrorMsg != "") {
				return;
			}

			// Set options
			let tmpOptions: SelectOption[] = [];

			response.Inputs.forEach((item) => {
				tmpOptions.push({ label: item, value: item });
			});

			setOptions(tmpOptions);
		});
	}

	/** Method to handle onChange of CreatableSelect. */
	const handleChange = (newValue: MultiValue<SelectOption>, actionMeta: ActionMeta<SelectOption>): void => {
		console.log("New value: " + JSON.stringify(newValue) + ", actionMeta: " + JSON.stringify(actionMeta));

		// Skip if setValue is null
		if (setValue === undefined) {
			return;
		}

		// Handle create
		if (actionMeta.action === "create-option" && newValue != undefined) {
			setValue(concatOptions(newValue));
			return;
		}

		// Handle clear
		if (actionMeta.action === "clear") {
			setValue("");
			return;
		}

		// Handle Select
		if (actionMeta.action === "select-option" && newValue != undefined) {
			setValue(concatOptions(newValue));
			return;
		}

		// Handle Remove
		if (actionMeta.action === "remove-value" && actionMeta.removedValue != undefined) {
			setValue(concatOptions(newValue));
			return;
		}
	};

	// CSS options for components
	const selectStyles: StylesConfig<SelectOption, true, GroupBase<SelectOption>> = {
		container: (baseStyles) => ({
			...baseStyles,
			border: " var(--bs-border-width) solid var(--bs-border-color)",
			borderRadius: "var(--bs-border-radius)",
			transition: "border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out",
			textAlign: "left",
		}),
		control: (baseStyles) => ({
			...baseStyles,
			backgroundColor: "transparent",
			borderStyle: "none",
		}),
		indicatorSeparator: (baseStyles) => ({
			...baseStyles,
			display: "none",
		}),
		input: (baseStyles) => ({
			...baseStyles,
			color: "lightgrey",
		}),
		multiValue: (baseStyles) => ({
			...baseStyles,
			backgroundColor: "#6c757d !important",
			border: "1px solid #495057",
			borderRadius: "0.375rem",
		}),
		multiValueLabel: (baseStyles) => ({
			...baseStyles,
			color: "lightgrey",
			fontSize: "100%",
			padding: "1px",
		}),
		menu: (baseStyles) => ({
			...baseStyles,
			background: "var(--bs-gray-dark)",
			marginTop: "0.125rem",
			border: "solid 1px",
			borderColor: "var(--bs-border-color)",
		}),
		option: (baseStyles) => ({
			...baseStyles,
			backgroundColor: "inherit",
		}),
	};

	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2" className={titleClass != undefined ? titleClass : ""}>
				{title}
			</Form.Label>
			<Col sm="9">
				<CreatableSelect<SelectOption, true, GroupBase<SelectOption>>
					isMulti
					className="dark-creatable-select"
					isClearable
					onChange={handleChange}
					options={options}
					value={convert(value)}
					isDisabled={disabled}
					styles={selectStyles}
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
	/** Handler when Year field has changed */
	onYearChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
	/** Handler when select HTML element changed */
	onSelectChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
};

/**
 * Create a uniform Form.Group Element as Row. This element specify designed for input date.
 * @returns A Row Element, Contains one input group with label. Input group contains three input field for year, month, day.
 */
export function FormDateRow({ title, year, month, day, disabled, onYearChange, onSelectChange }: FormDateRowProps) {
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

/** Props for `FormSelectRow`. */
type FormSelectRowProps = {
	/** the title/label of this input group */
	title: string;
	/** The JSX.Element of `<select>`. */
	selectElement: JSX.Element;
};

/**
 * Create a uniform Form.Group Element as Row.
 * This element specify designed for select element.
 */
export function FormSelectRow({ title, selectElement }: FormSelectRowProps) {
	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2">
				{title}
			</Form.Label>
			<Col sm="9">{selectElement}</Col>
		</Form.Group>
	);
}

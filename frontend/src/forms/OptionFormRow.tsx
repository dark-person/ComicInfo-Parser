// React
import { useEffect, useState } from "react";

// React Component
import { Col, Row } from "react-bootstrap";
import Form from "react-bootstrap/Form";

import { ActionMeta, GroupBase, MultiValue, StylesConfig } from "react-select";
import CreatableSelect from "react-select/creatable";

import { application } from "../../wailsjs/go/models";

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
	getDefaultOpt: () => Promise<application.HistoryResp>;
	/** Max height for menu list, same with css `max-height`. Default is `9em` */
	menuMaxHeight?: string;
	/** Current height for component, same with css `height`.*/
	componentHeight?: string;
};

/** Create a uniform Form.Group Element as Row, which contains select component that allow create new options. */
export default function OptionFormRow({
	title,
	titleClass,
	value,
	disabled,
	setValue,
	getDefaultOpt,
	menuMaxHeight,
	componentHeight,
}: Readonly<OptionFormRowProps>) {
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
		const simpleOptions = opts.map((item) => item.value);
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
		const splitOpts = opt.split(",");

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

			if (response.ErrorMsg !== "") {
				return;
			}

			// Set options
			const tmpOptions: SelectOption[] = [];

			response.Inputs.forEach((item: string) => {
				tmpOptions.push({ label: item, value: item });
			});

			setOptions(tmpOptions);
		});
	}

	/** Method to handle onChange of CreatableSelect. */
	const handleChange = (newValue: MultiValue<SelectOption>, actionMeta: ActionMeta<SelectOption>): void => {
		console.log(`New value: ${JSON.stringify(newValue)}, actionMeta: ${JSON.stringify(actionMeta)}`);

		// Skip if setValue is null
		if (setValue === undefined) {
			return;
		}

		// Handle create
		if (actionMeta.action === "create-option" && newValue !== undefined) {
			setValue(concatOptions(newValue));
			return;
		}

		// Handle clear
		if (actionMeta.action === "clear") {
			setValue("");
			return;
		}

		// Handle Select
		if (actionMeta.action === "select-option" && newValue !== undefined) {
			setValue(concatOptions(newValue));
			return;
		}

		// Handle Remove
		if (actionMeta.action === "remove-value" && actionMeta.removedValue !== undefined) {
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
			height: componentHeight ?? baseStyles.height,
			alignItems: componentHeight === null ? baseStyles.alignItems : "flex-start",
		}),
		indicatorsContainer: (baseStyles) => ({
			...baseStyles,
			alignItems: componentHeight === null ? baseStyles.alignItems : "flex-start",
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
		menuList: (baseStyles) => ({
			...baseStyles,
			maxHeight: menuMaxHeight ?? "9em",
		}),
		option: (baseStyles) => ({
			...baseStyles,
			backgroundColor: "inherit",
		}),
	};

	return (
		<Form.Group as={Row} className="mb-3">
			<Form.Label column sm="2" className={titleClass ?? ""}>
				{title}
			</Form.Label>
			<Col sm="9">
				<CreatableSelect<SelectOption, true, GroupBase<SelectOption>>
					isMulti
					menuPlacement="auto"
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

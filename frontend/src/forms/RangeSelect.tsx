// React
import { ChangeEventHandler } from "react";

// React Component
import Form from "react-bootstrap/Form";

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
export default function RangeSelect({ min, max, value, title, disabled, onChange }: Readonly<RangeSelectProps>) {
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

// React
import { type ChangeEvent } from "react";

// React Component
import { Form } from "react-bootstrap";

// Project Specified Component
import FormDateRow from "@/forms/FormDateRow";
import FormRow from "@/forms/FormRow";
import { type MetadataProps } from "@/pages/metadata/MetadataProps";

/**
 * The interface for show/edit book metadata.
 * @returns JSX Element
 */
export default function BookMetadata({ comicInfo: info, infoSetter }: Readonly<MetadataProps>) {
    /**
     * Handler for all input field in this panel.
     * This method will use <code>infoSetter</code> as core,
     * and apply change to comicInfo content.
     * <p>
     * This method will try to find the field is number input first,
     * if field name is number/related, then it will call Number() method before set to ComicInfo
     *
     * @param e the event object, for identify target element
     * @returns void
     */
    function handleChanges(e: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLSelectElement>) {
        console.log(e.target.title, e.target.value, e.target.type);

        // Identify Number & Year
        if (e.target.type === "number" || e.target.type === "Year") {
            infoSetter(e.target.title, Number(e.target.value));
            return;
        }

        // Identify Month & Day
        if (e.target.title === "Month" || e.target.title === "Day") {
            infoSetter(e.target.title, Number(e.target.value));
            return;
        }

        // Normal Cases
        infoSetter(e.target.title, e.target.value);
    }

    return (
        <div>
            <Form>
                <FormRow title={"Title"} value={info?.Title} onChange={handleChanges} />
                <FormRow title={"Summary"} value={info?.Summary.InnerXML} textareaRow={3} onChange={handleChanges} />
                <FormRow title={"Number"} value={info?.Number} onChange={handleChanges} />
                <FormDateRow
                    title={"Year/Month/Day"}
                    year={info?.Year}
                    month={info?.Month}
                    day={info?.Day}
                    onYearChange={handleChanges}
                    onSelectChange={handleChanges}
                />
                <FormRow title={"Web"} value={info?.Web} onChange={handleChanges} />
                <FormRow title={"GTIN"} value={info?.GTIN} onChange={handleChanges} />
            </Form>
        </div>
    );
}

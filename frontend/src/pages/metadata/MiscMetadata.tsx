// React
import { ChangeEvent } from "react";

// React Component
import { Form } from "react-bootstrap";

// Project Specified Component
import FormRow from "../../forms/FormRow";
import { MetadataProps } from "./MetadataProps";

/** The user interface for show/edit Collection & ReadList MetaData. */
export default function MiscMetadata({ comicInfo, infoSetter }: Readonly<MetadataProps>) {
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
                <FormRow title={"SeriesGroup"} value={comicInfo?.SeriesGroup} onChange={handleChanges} />
                <FormRow title={"AlternateSeries"} value={comicInfo?.AlternateSeries} onChange={handleChanges} />
                <FormRow title={"AlternateNumber "} value={comicInfo?.AlternateNumber} onChange={handleChanges} />
                <FormRow title={"AlternateCount"} value={comicInfo?.AlternateCount} onChange={handleChanges} />
                <FormRow title={"StoryArc"} value={comicInfo?.StoryArc} onChange={handleChanges} />
                <FormRow title={"StoryArcNumber"} value={comicInfo?.StoryArcNumber} onChange={handleChanges} />
            </Form>
        </div>
    );
}
